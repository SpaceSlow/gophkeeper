package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=8"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Errors string `json:"errors"`
}

func LoginHandler(c *gin.Context) {
	var registerRequest LoginRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, LoginResponse{
			Errors: err.Error(), // TODO: error handling
		})
		return
	}

	// TODO: hash password and check data in db, if ok return Bearer token

	c.JSON(http.StatusOK, LoginResponse{
		Token: "token for authentication",
	})
}
