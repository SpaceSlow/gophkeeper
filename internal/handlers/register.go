package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func RegisterHandler(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, RegisterResponse{
			Errors: err.Error(), // TODO: error handling
		})
		return
	}

	//TODO: add data to db

	c.JSON(http.StatusOK, RegisterResponse{
		Message: "user registered",
	})
}
