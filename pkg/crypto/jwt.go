package crypto

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UsernameClaims struct {
	jwt.RegisteredClaims
	UserID int
}

func BuildJWT(userID int, tokenLifetime time.Duration, secretKey string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, UsernameClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenLifetime)),
		},
		UserID: userID,
	})

	jwt, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func UserIDFromToken(tokenString, secretKey string) (int, error) {
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
		return 0, err
	}

	if !token.Valid {
		return 0, ErrInvalidToken
	}

	return claims.UserID, nil
}
