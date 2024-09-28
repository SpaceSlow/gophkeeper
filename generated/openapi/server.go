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
	PostRegister(c *gin.Context)
	// Returns sensitive records
	// (GET /sensitive_records)
	ListSensitiveRecords(c *gin.Context)
	// Create a new sensitive record
	// (POST /sensitive_records)
	PostSensitiveRecord(c *gin.Context)
	// Delete sensitive record with {id}
	// (DELETE /sensitive_records/{id})
	DeleteSensitiveRecordWithID(c *gin.Context, id int)
	// Returns data for sensitive record with {id}
	// (GET /sensitive_records/{id})
	FetchSensitiveRecordWithID(c *gin.Context, id int)
	// Upload binary data of sensitive record
	// (POST /sensitive_records/{id})
	PostSensitiveRecordData(c *gin.Context, id int)
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

// PostRegister operation middleware
func (siw *ServerInterfaceWrapper) PostRegister(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostRegister(c)
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

	c.Set(AuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostSensitiveRecord(c)
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

// FetchSensitiveRecordWithID operation middleware
func (siw *ServerInterfaceWrapper) FetchSensitiveRecordWithID(c *gin.Context) {

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

	siw.Handler.FetchSensitiveRecordWithID(c, id)
}

// PostSensitiveRecordData operation middleware
func (siw *ServerInterfaceWrapper) PostSensitiveRecordData(c *gin.Context) {

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

	siw.Handler.PostSensitiveRecordData(c, id)
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
	router.POST(options.BaseURL+"/register", wrapper.PostRegister)
	router.GET(options.BaseURL+"/sensitive_records", wrapper.ListSensitiveRecords)
	router.POST(options.BaseURL+"/sensitive_records", wrapper.PostSensitiveRecord)
	router.DELETE(options.BaseURL+"/sensitive_records/:id", wrapper.DeleteSensitiveRecordWithID)
	router.GET(options.BaseURL+"/sensitive_records/:id", wrapper.FetchSensitiveRecordWithID)
	router.POST(options.BaseURL+"/sensitive_records/:id", wrapper.PostSensitiveRecordData)
}
