package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/users"
)

func (h UserHandlers) RegisterUser(c *gin.Context) {
	var req openapi.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, openapi.ErrorResponse{
			Errors: err.Error(), // TODO: error handling
		})
		return
	}

	if existed, err := h.repo.ExistUsername(req.Username); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if existed {
		c.JSON(http.StatusConflict, openapi.ErrorResponse{
			Errors: users.NewRegisteredUserError(req.Username).Error(),
		})
		return
	}

	user, err := users.CreateUser(req.Username, req.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	passwordHash, err := user.GeneratePasswordHash(h.cfg.KeyLen(), h.cfg.PasswordIterationNum())
	err = h.repo.RegisterUser(req.Username, passwordHash)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}
