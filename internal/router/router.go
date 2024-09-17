package router

import (
	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/handlers"
	"github.com/SpaceSlow/gophkeeper/internal/middlewares"
	"github.com/SpaceSlow/gophkeeper/internal/store"
)

func SetupRouter(db *store.DB) *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")
	public.POST("/register", handlers.RegisterHandler(db))
	public.POST("/login", handlers.LoginHandler(db))

	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware(db))
	protected.POST("/sensitive_records", handlers.SensitiveRecordUploadHandler(db))

	return router
}
