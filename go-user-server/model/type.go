package model

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-sql/civil"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	JSON      = datatypes.JSON
	JSONMap   = datatypes.JSONMap
	NullInt64 = sql.NullInt64
	NullTime  = sql.NullTime
	NullUUID  = uuid.NullUUID
	Time      = time.Time
	UUID      = uuid.UUID
)

type Rune rune

func (r *Rune) Scan(src interface{}) error {
	runes := []rune(src.(string))
	*r = Rune(runes[0])
	return nil
}

func (r Rune) Value() (driver.Value, error) {
	return string(r), nil
}

func (Rune) GormDataType() string {
	return "char(1)"
}

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
)

func (s *Sex) Scan(src interface{}) error {
	*s = Sex(src.(string))
	return nil
}

func (s Sex) Value() (driver.Value, error) {
	return string(s), nil
}

func (s Sex) String() string {
	return string(s)
}

func (Sex) GormDataType() string {
	return "sex"
}

type Weekday time.Weekday

func (w *Weekday) Scan(src interface{}) error {
	switch src.(string) {
	case "sunday":
		*w = Weekday(time.Sunday)
	case "monday":
		*w = Weekday(time.Monday)
	case "tuesday":
		*w = Weekday(time.Tuesday)
	case "wednesday":
		*w = Weekday(time.Wednesday)
	case "thursday":
		*w = Weekday(time.Thursday)
	case "friday":
		*w = Weekday(time.Friday)
	case "saturday":
		*w = Weekday(time.Saturday)
	default:
		return fmt.Errorf("model: Weekday.Scan: src value '%s' is not valid", src)
	}
	return nil
}

func (w Weekday) Value() (driver.Value, error) {
	return w.String(), nil
}

func (w Weekday) String() string {
	return strings.ToLower(time.Weekday(w).String())
}

func (Weekday) GormDataType() string {
	return "weekday"
}

// Reference: https://github.com/golang-sql/civil/blob/master/civil.go

type Date civil.Date

func DateOf(t time.Time) Date {
	return Date(civil.DateOf(t))
}

func ParseDate(s string) (Date, error) {
	if d, err := civil.ParseDate(s); err != nil {
		return Date{}, err
	} else {
		return Date(d), nil
	}
}

func (d Date) String() string {
	return civil.Date(d).String()
}

func (d Date) In(loc *time.Location) time.Time {
	return civil.Date(d).In(loc)
}

func (d *Date) Scan(src interface{}) error {
	if time, ok := src.(time.Time); ok {
		*d = Date{
			Year:  time.Year(),
			Month: time.Month(),
			Day:   time.Day(),
		}
		return nil
	}
	return errors.New("model: Date.Scan: src must be time.Time type")
}

func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func (Date) GormDataType() string {
	return "date"
}

/*

// https://github.com/go-gorm/datatypes/blob/master/json_map.go

type JSONSlice []interface{}

// Value return json value, implement driver.Valuer interface
func (m JSONSlice) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	ba, err := m.MarshalJSON()
	return string(ba), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (m *JSONSlice) Scan(val interface{}) error {
	if val == nil {
		*m = make(JSONSlice, 0)
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	t := []interface{}{}
	err := json.Unmarshal(ba, &t)
	*m = t
	return err
}

// MarshalJSON to output non base64 encoded []byte
func (m JSONSlice) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	t := ([]interface{})(m)
	return json.Marshal(t)
}

// UnmarshalJSON to deserialize []byte
func (m *JSONSlice) UnmarshalJSON(b []byte) error {
	t := []interface{}{}
	err := json.Unmarshal(b, &t)
	*m = JSONSlice(t)
	return err
}

// GormDataType gorm common data type
func (m JSONSlice) GormDataType() string {
	return "jsonslice"
}

// GormDBDataType gorm db data type
func (JSONSlice) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return new(JSON).GormDBDataType(db, field)
}

func (jm JSONSlice) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	data, _ := jm.MarshalJSON()
	return JSON(data).GormValue(ctx, db)
}

*/
