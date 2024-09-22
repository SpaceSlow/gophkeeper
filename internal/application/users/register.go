package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/domain/users"
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

func (h UserHandlers) RegisterUser(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Errors: err.Error(), // TODO: error handling
		})
		return
	}

	if existed, err := h.repo.ExistUsername(registerRequest.Username); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if existed {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	user, err := users.CreateUser(registerRequest.Username, registerRequest.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	passwordHash, err := user.GeneratePasswordHash(h.cfg.KeyLen(), h.cfg.PasswordIterationNum())
	err = h.repo.RegisterUser(registerRequest.Username, passwordHash)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		Message: "user registered",
	})
}
