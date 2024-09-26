package application

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	ginmiddleware "github.com/oapi-codegen/gin-middleware"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
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

	server := &httpServer{}
	server.UserHandlers = users.SetupHandlers(userRepo, cfg)
	server.SensitiveRecordHandlers = sensitive_records.SetupHandlers(sensitiveRecordRepo)

	spec, _ := openapi.GetSwagger()
	m := crypto.NewJWTMiddleware(cfg.SecretKey(), userRepo)

	validator := ginmiddleware.OapiRequestValidatorWithOptions(spec,
		&ginmiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
					return m.AuthenticateRequest(ctx, input)
				},
			},
		})

	api := router.Group("/api", router.Handlers...)
	api.Use(validator)
	openapi.RegisterHandlers(api, server)

	return router
}
