package user

import (
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zeebo/blake3"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"apulse.ai/tzuchi-upmp/server/model"
	"apulse.ai/tzuchi-upmp/server/service"
	"apulse.ai/tzuchi-upmp/server/service/requester"
	"apulse.ai/tzuchi-upmp/server/store"
	"apulse.ai/tzuchi-upmp/server/utils/argon2id"
)

type Handler struct {
	store   *store.Store
	service *service.Service
}

func NewHandler(store *store.Store, service *service.Service) *Handler {
	return &Handler{store, service}
}

func getCurrentUserID(context echo.Context) uuid.UUID {
	return uuid.MustParse(context.Get("jwt").(*jwt.Token).Claims.(*jwt.StandardClaims).Subject)
}

func (h *Handler) GetUsers(context echo.Context) error {
	if data, err := h.store.User.GetData(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return context.JSON(http.StatusOK, newUsersResponse(data))
	}
}

func (h *Handler) AddUsers(context echo.Context) error {
	request := new(addUsersRequest)
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	data := request.toModels()
	if err := h.store.User.AddData(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	for _, datum := range data {
		if _, err := h.store.Enforcer.AddRoleForUser(datum.ID.String(), "admin"); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return context.JSON(http.StatusCreated, newUsersResponse(data))
}

func (h *Handler) GetUserByUserID(context echo.Context) error {
	request := new(getUserByUserIDRequest)
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	userID := uuid.MustParse(request.UserID)
	if datum, err := h.store.User.GetDatumByID(userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return context.JSON(http.StatusOK, newUserResponse(datum))
	}
}

func (h *Handler) UpdateUserByUserID(context echo.Context) error {
	request := new(updateUserByUserIDRequest)
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	userID := uuid.MustParse(request.UserID)
	patch := request.toPatch()
	if err := h.store.User.UpdateDatumByID(userID, patch); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteUserByUserID(context echo.Context) error {
	request := new(deleteUserByUserIDRequest)
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	userID := uuid.MustParse(request.UserID)
	if userID == getCurrentUserID(context) {
		return currentUserCanNotBeDeletedError
	}
	if err := h.store.User.DeleteDatumByID(userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if _, err := h.store.Enforcer.DeleteRolesForUser(request.UserID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.NoContent(http.StatusNoContent)
}

/*
func (h *Handler) GetRolesOfUserByUserID(context echo.Context) error {
	request := new(getRolesOfUserByUserIDRequest)
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	if _, err := h.store.User.GetUserByUserID(uuid.MustParse(request.UserID)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if roles, err := h.store.Enforcer.GetRolesForUser(request.UserID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return context.JSON(http.StatusOK, newRolesOfUserResponse(roles))
	}
}

func (h *Handler) SetRolesOfUserByUserID(context echo.Context) error {
	request := new(setRolesOfUserByUserIDRequest)
	binder := new(echo.DefaultBinder)
	if err := binder.BindPathParams(context, request); err != nil {
		return err
	}
	if err := binder.BindBody(context, &request.Roles); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	if _, err := h.store.User.GetUserByUserID(uuid.MustParse(request.UserID)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if _, err := h.store.Enforcer.DeleteRolesForUser(request.UserID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if ok, err := h.store.Enforcer.AddRolesForUser(request.UserID, request.Roles); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if !ok {
		return setRolesOfUserFailedError
	}
	return context.NoContent(http.StatusNoContent)
}
*/

func (h *Handler) CheckUserExists(context echo.Context) error {
	request := new(checkUserExistsRequest)
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}

	hashPassword := func(username, password string) []byte {
		hashedPassword := blake3.Sum512([]byte(password))
		return argon2id.Hash(username, hex.EncodeToString(hashedPassword[:]))
	}

	requestDataOfSignIn := datatypes.JSONMap{"username": request.Username, "password": request.Password}
	if responseOfSignIn, err := requester.MakeRequestToHDSS(
		http.MethodPost, "users/login/", requestDataOfSignIn, nil, nil, []int{http.StatusOK}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		signInMap := responseOfSignIn.(map[string]interface{})
		accessToken := signInMap["access_token"].(string)
		userID := signInMap["id"].(string)
		if responseOfUser, err := requester.MakeRequestToHDSS(
			http.MethodGet, "users/"+userID+"/", nil, &accessToken, nil, []int{http.StatusOK}); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			password := hashPassword(request.Username, request.Password)
			userMap := responseOfUser.(map[string]interface{})
			datum := model.User{
				Username: userMap["username"].(string),
				Password: password,
				FullName: userMap["full_name"].(string),
				IDNumber: userMap["id_number"].(string),
				Email:    userMap["email"].(string),
				Role:     userMap["role"].(string),
			}
			if !model.IsIDNumberValid(datum.IDNumber) {
				datum.IDNumber = ""
			}
			data := []model.User{datum}
			if err := h.store.User.SyncData(data, true); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			if _, err := h.store.Enforcer.AddRolesForUser(data[0].ID.String(), []string{"admin"}); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
	}
	return context.NoContent(http.StatusOK)
}

func (h *Handler) SignIn(context echo.Context) error {
	request := new(signInRequest)
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	if datum, err := h.store.User.GetDatumByUsername(request.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usernameOrPasswordIsNotCorrectError
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if password, err := h.store.User.GetPasswordOfDatumByID(datum.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if !argon2id.IsMatched(request.Username, request.Password, password) {
		return usernameOrPasswordIsNotCorrectError
	} else if response, err := newAccessTokenResponse(datum); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return context.JSON(http.StatusOK, response)
	}
}

func (h *Handler) GetCurrentUser(context echo.Context) error {
	userID := getCurrentUserID(context)
	if datum, err := h.store.User.GetDatumByID(userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return context.JSON(http.StatusOK, newUserResponse(datum))
	}
}

func (h *Handler) UpdateCurrentUser(context echo.Context) error {
	request := new(updateCurrentUserRequest)
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	userID := getCurrentUserID(context)
	patch := request.toPatch()
	if err := h.store.User.UpdateDatumByID(userID, patch); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.NoContent(http.StatusNoContent)
}
