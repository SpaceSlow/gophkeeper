package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/config"
	"github.com/SpaceSlow/gophkeeper/internal/store"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Errors string `json:"errors"`
}

func LoginHandler(db *store.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var loginRequest LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, LoginResponse{
				Errors: err.Error(), // TODO: error handling
			})
			return
		}

		fetchedPasswordHash, err := db.FetchPasswordHash(context.TODO(), loginRequest.Username)
		var errNoUser *store.NoUserError
		if errors.As(err, &errNoUser) {
			c.JSON(http.StatusUnauthorized, LoginResponse{
				Errors: errNoUser.Error(),
			})
			return
		} else if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		cfg := config.GetServerConfig()
		if isValid, err := crypto.IsValid(loginRequest.Password, fetchedPasswordHash, cfg.KeyLen); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		} else if !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		jwt, err := crypto.BuildJWT(loginRequest.Username, cfg.TokenLifetime, cfg.SecretKey)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, LoginResponse{
			Token: jwt,
		})
	}
}
