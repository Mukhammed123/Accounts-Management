package middleware

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/casbin/casbin/v2"
	echocasbin "github.com/labstack/echo-contrib/casbin"
)

const (
	jwtContextKey = "jwt"
)

func JWT() echo.MiddlewareFunc {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(context echo.Context) bool {
			return context.Path() == "/api/user/sign-in" || context.Path() == "/api/user/check"
		},
		SigningKey: secretKey,
		ContextKey: jwtContextKey,
		Claims:     &jwt.StandardClaims{},
	})
}

func Auth(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	return echocasbin.MiddlewareWithConfig(echocasbin.Config{
		Enforcer: enforcer,
		UserGetter: func(context echo.Context) (string, error) {
			userID := context.Get(jwtContextKey).(*jwt.Token).Claims.(*jwt.StandardClaims).Subject
			return userID, nil
		},
	})
}

func Assets() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "../assets",
		Index: "",
	})
}
