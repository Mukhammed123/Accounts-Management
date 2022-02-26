package handler

import (
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"

	"apulse.ai/tzuchi-upmp/server/middleware"
)

/*
GET Get
POST Add, actions
PATCH Update
PUT Set
DELETE Delete
*/

func (h *Handler) RegisterAPI(group *echo.Group, enforcer *casbin.Enforcer) {
	jwt := middleware.JWT()
	auth := middleware.Auth(enforcer)

	users := group.Group("/users", jwt, auth)
	enforcer.AddPolicy("admin", "/api/users", "^(GET|POST)$")
	users.GET("", h.user.GetUsers)
	users.POST("", h.user.AddUsers)
	user := users.Group("/:userID", jwt, auth)
	enforcer.AddPolicy("admin", "/api/users/:userID", "^(GET|PATCH|DELETE)$")
	user.GET("", h.user.GetUserByUserID)
	user.PATCH("", h.user.UpdateUserByUserID)
	user.DELETE("", h.user.DeleteUserByUserID)
	// enforcer.AddPolicy("admin", "/api/users/:userID/roles", "^(GET|PUT)$")
	// user.GET("/roles", h.user.GetRolesOfUserByUserID)
	// user.PUT("/roles", h.user.SetRolesOfUserByUserID)
	currentUser := group.Group("/user", jwt)
	currentUser.POST("/check", h.user.CheckUserExists)
	currentUser.POST("/sign-in", h.user.SignIn)
	currentUser.GET("", h.user.GetCurrentUser)
	currentUser.PATCH("", h.user.UpdateCurrentUser)
}

func (h *Handler) RegisterAssets(group *echo.Group) {
	jwt := middleware.JWT()
	assets := middleware.Assets()
	group.Use(assets, jwt)
}
