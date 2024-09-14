package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/config"
	"github.com/SpaceSlow/gophkeeper/internal/store"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

type RegisterRequest struct {
	Username         string `json:"username" binding:"required,min=8"`
	Password         string `json:"password" binding:"required,min=8"`
	RepeatedPassword string `json:"repeated_password" binding:"required,eqfield=Password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

func RegisterHandler(db *store.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var registerRequest RegisterRequest
		if err := c.ShouldBindJSON(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, RegisterResponse{
				Errors: err.Error(), // TODO: error handling
			})
			return
		}

		if exist, err := db.CheckUsername(context.TODO(), registerRequest.Username); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		} else if exist {
			c.AbortWithStatus(http.StatusConflict)
			return
		}

		cfg := config.GetServerConfig()
		passwordHash, err := crypto.GenerateHash(registerRequest.Password, cfg.KeyLen, cfg.PasswordIterationNum)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		err = db.RegisterUser(context.TODO(), registerRequest.Username, passwordHash)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, RegisterResponse{
			Message: "user registered",
		})
	}
}
