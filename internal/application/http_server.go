package application

import (
	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/application/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/internal/application/users"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

type httpServer struct {
	users.UserHandlers
	sensitive_records.SensitiveRecordHandlers
}

func SetupHTTPServer(userRepo UserRepository, sensitiveRecordRepo SensitiveRecordRepository, cfg users.ConfigProvider) *gin.Engine {
	router := gin.Default()

	server := httpServer{}
	server.UserHandlers = users.SetupHandlers(userRepo, cfg)
	server.SensitiveRecordHandlers = sensitive_records.SetupHandlers(sensitiveRecordRepo)

	public := router.Group("/api")
	public.POST("/register", server.RegisterUser)
	public.POST("/login", server.LoginUser)

	protected := router.Group("/api")
	public.GET("/sensitive_record_types", server.ListSensitiveRecordTypes)
	protected.Use(crypto.AuthMiddleware(userRepo, cfg))
	protected.POST("/sensitive_records", server.PostSensitiveRecord)

	return router
}
