package application

import (
	"github.com/SpaceSlow/gophkeeper/internal/application/middlewares"
	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/application/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/internal/application/users"
)

func SetupRouter(userRepo users.Repository, sensitiveRecordRepo sensitive_records.Repository) *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")

	userHandlers := users.SetupHandlers(userRepo)
	public.POST("/register", userHandlers.RegisterUser)
	public.POST("/login", userHandlers.LoginUser)

	protected := router.Group("/api")
	sensitiveRecordHandlers := sensitive_records.SetupHandlers(sensitiveRecordRepo)
	protected.Use(middlewares.AuthMiddleware(userRepo))
	protected.POST("/sensitive_records", sensitiveRecordHandlers.UploadSensitiveRecord)

	return router
}
