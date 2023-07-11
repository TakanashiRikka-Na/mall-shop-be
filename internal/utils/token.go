package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type JWT struct {
	jwt.RegisteredClaims
	UserName string
}

func GenerateToken(userID string) (string, error) {
	Key := os.Getenv("JWT_KEY")
	prefix := "Bearer "

	claims := JWT{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour))},
		UserName:         userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Key))
	if err != nil {
		return "", err
	}

	return prefix + signedToken, nil

}
