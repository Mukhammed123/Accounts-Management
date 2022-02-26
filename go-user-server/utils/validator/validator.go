package validator

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"apulse.ai/tzuchi-upmp/server/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

// https://security.stackexchange.com/questions/53594/why-is-client-side-hashing-of-a-password-so-uncommon
var (
	dateRegexp     = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
	timeRegexp     = regexp.MustCompile(`^(\d{2}):(\d{2})$`)
	usernameRegexp = regexp.MustCompile(`^\w[\w-.~]{2,}\w$`)
	passwordRegexp = regexp.MustCompile(`^[0-9a-f]{128}$`) // hex after hashing with blake3 (64 bytes)
)

func mustValidator(fl validator.FieldLevel) bool {
	acceptedValues := strings.Split(fl.Param(), " ")
	field := fl.Field()
	isValid := func(value reflect.Value) bool {
		// Reference: https://bit.ly/3ADOmRx
		var reflected string
		switch value.Kind() {
		case reflect.String:
			reflected = value.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			reflected = strconv.FormatInt(value.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			reflected = strconv.FormatUint(value.Uint(), 10)
		default:
			panic(fmt.Sprintf("Bad field type %T", value.Interface()))
		}
		for i := 0; i < len(acceptedValues); i++ {
			if acceptedValues[i] == reflected {
				return true
			}
		}
		return false
	}
	if field.Kind() == reflect.Slice {
		for i := 0; i < field.Len(); i++ {

			if !isValid(field.Index(i)) {
				return false
			}
		}
		return true
	} else {
		return isValid(field)
	}
}

func dateValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if len(value) == 0 {
		return true
	}
	if matched := dateRegexp.FindStringSubmatch(value); matched == nil || len(matched) != 4 {
		return false
	} else if year, err := strconv.Atoi(matched[1]); err != nil {
		return false
	} else if month, err := strconv.Atoi(matched[2]); err != nil {
		return false
	} else if day, err := strconv.Atoi(matched[3]); err != nil {
		return false
	} else {
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		return year == date.Year() && month == int(date.Month()) && day == date.Day()
	}
}

func timeValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if len(value) == 0 {
		return true
	}
	if matched := timeRegexp.FindStringSubmatch(value); matched == nil || len(matched) != 3 {
		return false
	} else if hour, err := strconv.Atoi(matched[1]); err != nil {
		return false
	} else if minute, err := strconv.Atoi(matched[2]); err != nil {
		return false
	} else {
		return hour < 24 && minute < 60
	}
}

func usernameValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return usernameRegexp.MatchString(value)
}

func passwordValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return passwordRegexp.MatchString(value)
}

func caseNumberValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return model.IsCaseNumberValid(value)
}

func idNumberValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return model.IsIDNumberValid(value)
}

func NewValidator() *Validator {
	validator := validator.New()
	validator.RegisterValidation("must", mustValidator)
	validator.RegisterValidation("date", dateValidator)
	validator.RegisterValidation("time", timeValidator)
	validator.RegisterValidation("username", usernameValidator)
	validator.RegisterValidation("password", passwordValidator)
	validator.RegisterValidation("case_number", caseNumberValidator)
	validator.RegisterValidation("id_number", idNumberValidator)
	return &Validator{
		validator: validator,
	}
}

func (v *Validator) Validate(i interface{}) error {
	value := reflect.Indirect(reflect.ValueOf(i))
	switch value.Kind() {
	case reflect.Struct:
		if err := v.validator.Struct(i); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if err := v.validator.Struct(value.Index(i).Interface()); err != nil {
				return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
			}
		}
	}
	return nil
}
