package crypto

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UsernameClaims struct {
	jwt.RegisteredClaims
	Username string
}

func BuildJWT(username string, tokenLifetime time.Duration, secretKey string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, UsernameClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenLifetime)),
		},
		Username: username,
	})

	jwt, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func Username(tokenString, secretKey string) (string, error) {
	claims := &UsernameClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
	)
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", ErrInvalidToken
	}

	return claims.Username, nil
}
