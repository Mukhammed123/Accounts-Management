package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func Generate(userID uuid.UUID) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	now := time.Now()
	claims := jwt.StandardClaims{
		Subject:   userID.String(),
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Hour * 24 * 7).Unix(),
	}
	if jti, err := uuid.NewRandom(); err != nil {
		return "", err
	} else {
		claims.Id = jti.String()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if signed, err := token.SignedString(secretKey); err != nil {
		return "", err
	} else {
		return signed, nil
	}
}
