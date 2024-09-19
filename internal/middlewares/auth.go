package middlewares

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal"
	"github.com/SpaceSlow/gophkeeper/internal/application/users"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func AuthMiddleware(r users.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isAuthenticated(c, r) {
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func isAuthenticated(c *gin.Context, r users.Repository) bool {
	jwt, err := crypto.ExtractToken(c)
	if err != nil {
		return false
	}
	username, err := crypto.Username(jwt, internal.GetServerConfig().SecretKey)
	if err != nil {
		return false
	}
	isExisted, err := r.ExistUsername(context.TODO(), username)
	if err != nil || !isExisted {
		return false
	}
	return true
}
