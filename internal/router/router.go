package router

import (
	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")
	public.POST("/register", handlers.RegisterHandler)

	return router
}
