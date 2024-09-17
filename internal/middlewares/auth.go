package middlewares

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/config"
	"github.com/SpaceSlow/gophkeeper/internal/store"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func AuthMiddleware(db *store.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isAuthenticated(c, db) {
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func isAuthenticated(c *gin.Context, db *store.DB) bool {
	jwt, err := crypto.ExtractToken(c)
	if err != nil {
		return false
	}
	username, err := crypto.Username(jwt, config.GetServerConfig().SecretKey)
	if err != nil {
		return false
	}
	isExisted, err := db.ExistUsername(context.TODO(), username)
	if err != nil || !isExisted {
		return false
	}
	return true
}
