// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Login user
	// (POST /login)
	PostLogin(c *gin.Context)
	// Register user
	// (POST /register)
	RegisterUser(c *gin.Context)
	// Returns sensitive record types
	// (GET /sensitive_record_types)
	ListSensitiveRecordTypes(c *gin.Context)
	// Returns sensitive records
	// (GET /sensitive_records)
	ListSensitiveRecords(c *gin.Context)
	// Create a new sensitive record
	// (POST /sensitive_records)
	PostSensitiveRecord(c *gin.Context, params PostSensitiveRecordParams)
	// Create file of sensitive record
	// (POST /sensitive_records/files)
	UploadFile(c *gin.Context)
	// Get file of sensitive record
	// (GET /sensitive_records/files/{hash})
	DownloadFile(c *gin.Context, hash string)
	// Delete sensitive record with {id}
	// (DELETE /sensitive_records/{id})
	DeleteSensitiveRecordWithID(c *gin.Context, id int)
	// Returns data for sensitive record with {id}
	// (GET /sensitive_records/{id})
	SensitiveRecordDataWithID(c *gin.Context, id int)
	// Create data for sensitive record with {id}
	// (POST /sensitive_records/{id})
	CreateSensitiveRecordDataWithID(c *gin.Context, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostLogin(c)
}

// RegisterUser operation middleware
func (siw *ServerInterfaceWrapper) RegisterUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.RegisterUser(c)
}

// ListSensitiveRecordTypes operation middleware
func (siw *ServerInterfaceWrapper) ListSensitiveRecordTypes(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListSensitiveRecordTypes(c)
}

// ListSensitiveRecords operation middleware
func (siw *ServerInterfaceWrapper) ListSensitiveRecords(c *gin.Context) {

	c.Set(AuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListSensitiveRecords(c)
}

// PostSensitiveRecord operation middleware
func (siw *ServerInterfaceWrapper) PostSensitiveRecord(c *gin.Context) {

	var err error

	c.Set(AuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostSensitiveRecordParams

	// ------------- Optional query parameter "type" -------------

	err = runtime.BindQueryParameter("form", true, false, "type", c.Request.URL.Query(), &params.Type)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter type: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostSensitiveRecord(c, params)
}

// UploadFile operation middleware
func (siw *ServerInterfaceWrapper) UploadFile(c *gin.Context) {

	c.Set(AuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UploadFile(c)
}

// DownloadFile operation middleware
func (siw *ServerInterfaceWrapper) DownloadFile(c *gin.Context) {

	var err error

	// ------------- Path parameter "hash" -------------
	var hash string

	err = runtime.BindStyledParameterWithOptions("simple", "hash", c.Param("hash"), &hash, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter hash: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DownloadFile(c, hash)
}

// DeleteSensitiveRecordWithID operation middleware
func (siw *ServerInterfaceWrapper) DeleteSensitiveRecordWithID(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteSensitiveRecordWithID(c, id)
}

// SensitiveRecordDataWithID operation middleware
func (siw *ServerInterfaceWrapper) SensitiveRecordDataWithID(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SensitiveRecordDataWithID(c, id)
}

// CreateSensitiveRecordDataWithID operation middleware
func (siw *ServerInterfaceWrapper) CreateSensitiveRecordDataWithID(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(AuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateSensitiveRecordDataWithID(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/login", wrapper.PostLogin)
	router.POST(options.BaseURL+"/register", wrapper.RegisterUser)
	router.GET(options.BaseURL+"/sensitive_record_types", wrapper.ListSensitiveRecordTypes)
	router.GET(options.BaseURL+"/sensitive_records", wrapper.ListSensitiveRecords)
	router.POST(options.BaseURL+"/sensitive_records", wrapper.PostSensitiveRecord)
	router.POST(options.BaseURL+"/sensitive_records/files", wrapper.UploadFile)
	router.GET(options.BaseURL+"/sensitive_records/files/:hash", wrapper.DownloadFile)
	router.DELETE(options.BaseURL+"/sensitive_records/:id", wrapper.DeleteSensitiveRecordWithID)
	router.GET(options.BaseURL+"/sensitive_records/:id", wrapper.SensitiveRecordDataWithID)
	router.POST(options.BaseURL+"/sensitive_records/:id", wrapper.CreateSensitiveRecordDataWithID)
}