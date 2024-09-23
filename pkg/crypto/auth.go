package crypto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config interface {
	SecretKey() string
}

type Repository interface {
	ExistUser(userID int) (bool, error)
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
	userID, err := UserID(jwt, cfg.SecretKey())
	if err != nil {
		return false
	}
	isExisted, err := r.ExistUser(userID)
	if err != nil || !isExisted {
		return false
	}
	return true
}
