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

	user, err := h.repo.FetchUser(req.Username)
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

	if isValid, err := user.CheckPasswordHash(req.Password, h.cfg.KeyLen()); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if !isValid {
		c.JSON(http.StatusUnauthorized, openapi.ErrorResponse{
			Errors: users.ErrUserLogin.Error(),
		})
		return
	}

	jwt, err := crypto.BuildJWT(user.Id(), h.cfg.TokenLifetime(), h.cfg.SecretKey())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("X-Expires-After", time.Now().Add(h.cfg.TokenLifetime()).UTC().String())
	c.JSON(http.StatusOK, openapi.LoginUserResponse{
		Token: jwt,
	})
}
