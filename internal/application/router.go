package application

import (
	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/application/users"
	"github.com/SpaceSlow/gophkeeper/internal/handlers"
	"github.com/SpaceSlow/gophkeeper/internal/middlewares"
)

func SetupRouter(userRepo users.Repository) *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")

	userHandlers := users.SetupHandlers(userRepo)

	public.POST("/register", userHandlers.RegisterUser)
	public.POST("/login", userHandlers.LoginUser)

	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware(userRepo))
	sensitiveRecordHandler := handlers.NewSensitiveRecordHandler(db) // TODO: fix
	protected.POST("/sensitive_records", sensitiveRecordHandler.Upload())

	return router
}
