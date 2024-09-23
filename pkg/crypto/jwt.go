package crypto

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func ExtractToken(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return "", ErrInvalidAuthorizationHeader
	}
	if parts[0] != "Bearer" {
		return "", ErrInvalidAuthorizationHeader
	}
	return parts[1], nil
}

func UserID(tokenString, secretKey string) (int, error) {
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
