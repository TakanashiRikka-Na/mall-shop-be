package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

type JWT struct {
	jwt.RegisteredClaims
	UserName string
}

func AUTH() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				if tokenString == "" {
					return nil, errors.New("token is nil")
				}

				auths := strings.SplitN(tokenString, " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
					return nil, errors.New("jwt token missing")
				}
				var claims *JWT
				claims, err = ParseToken(auths[1])
				if err != nil {
					return nil, err
				}
				ctx = WithContext(ctx, claims.UserName)
			}
			return handler(ctx, req)
		}
	}
}

func WithContext(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, "username", username)
}
func ParseToken(tokenString string) (*JWT, error) {
	Key := os.Getenv("JWT_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &JWT{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWT)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
