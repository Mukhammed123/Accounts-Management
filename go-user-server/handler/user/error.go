package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	usernameOrPasswordIsNotCorrectError = echo.NewHTTPError(http.StatusForbidden, "Username or password is not correct")
	currentUserCanNotBeDeletedError     = echo.NewHTTPError(http.StatusBadRequest, "Current user can not be deleted")
	// setRolesOfUserFailedError = echo.NewHTTPError(http.StatusInternalServerError, "Set roles of user failed")
)
