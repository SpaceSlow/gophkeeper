package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/users"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func (h UserHandlers) LoginUser(c *gin.Context) {
	var req openapi.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, openapi.ErrorResponse{
			Errors: err.Error(), // TODO: error handling
		})
		return
	}

	fetchedPasswordHash, err := h.repo.FetchPasswordHash(req.Username)
	var errNoUser users.NoUserError
	if errors.As(err, &errNoUser) {
		c.JSON(http.StatusUnauthorized, openapi.ErrorResponse{
			Errors: errNoUser.Error(),
		})
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := users.CreateUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, openapi.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	if isValid, err := user.CheckPasswordHash(fetchedPasswordHash, h.cfg.KeyLen()); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if !isValid {
		c.JSON(http.StatusUnauthorized, openapi.ErrorResponse{
			Errors: users.ErrUserLogin.Error(),
		})
		return
	}

	jwt, err := crypto.BuildJWT(req.Username, h.cfg.TokenLifetime(), h.cfg.SecretKey())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("X-Expires-After", time.Now().Add(h.cfg.TokenLifetime()).UTC().String())
	c.JSON(http.StatusOK, openapi.LoginUserResponse{
		Token: jwt,
	})
}
