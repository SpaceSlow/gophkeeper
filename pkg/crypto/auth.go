package crypto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config interface {
	SecretKey() string
}

type Repository interface {
	ExistUsername(username string) (bool, error)
}

func AuthMiddleware(r Repository, cfg Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isAuthenticated(c, r, cfg) {
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func isAuthenticated(c *gin.Context, r Repository, cfg Config) bool {
	jwt, err := ExtractToken(c)
	if err != nil {
		return false
	}
	username, err := Username(jwt, cfg.SecretKey())
	if err != nil {
		return false
	}
	isExisted, err := r.ExistUsername(username)
	if err != nil || !isExisted {
		return false
	}
	return true
}
