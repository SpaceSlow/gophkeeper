package crypto

import (
	"context"
	"errors"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	ginmiddleware "github.com/oapi-codegen/gin-middleware"
)

type Repository interface {
	ExistUser(userID int) (bool, error)
}

type JWTMiddleware struct {
	secretKey string
	repo      Repository
}

func NewJWTMiddleware(secretKey string, repo Repository) *JWTMiddleware {
	return &JWTMiddleware{
		secretKey: secretKey,
		repo:      repo,
	}
}

func (m *JWTMiddleware) AuthenticateRequest(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	authHeader := input.RequestValidationInput.Request.Header.Get("Authorization")
	if authHeader == "" {
		return errors.New("missing or invalid token")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return errors.New("invalid token format")
	}

	jwt := parts[1]

	userID, err := UserIDFromToken(jwt, m.secretKey)
	if err != nil {
		return errors.New("invalid token format")
	}
	isExisted, err := m.repo.ExistUser(userID)
	if err != nil || !isExisted {
		return errors.New("invalid token")
	}

	ginCtx := ginmiddleware.GetGinContext(ctx)
	ginCtx.Set("userID", userID)

	return nil
}

func UserID(c *gin.Context) (int, error) {
	value, exist := c.Get("userID")
	if !exist {
		return 0, ErrNoUserID
	}
	userID, ok := value.(int)
	if !ok {
		return 0, ErrNoUserID
	}
	return userID, nil
}
