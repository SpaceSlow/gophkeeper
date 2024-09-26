// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	AuthScopes = "auth.Scopes"
)

// Defines values for SensitiveRecordTypeEnum.
const (
	Binary      SensitiveRecordTypeEnum = "binary"
	Credential  SensitiveRecordTypeEnum = "credential"
	PaymentCard SensitiveRecordTypeEnum = "payment-card"
	Text        SensitiveRecordTypeEnum = "text"
)

// CreateSensitiveRecordRequest defines model for CreateSensitiveRecordRequest.
type CreateSensitiveRecordRequest struct {
	Data     openapi_types.File      `json:"data"`
	Metadata string                  `json:"metadata"`
	Type     SensitiveRecordTypeEnum `json:"type"`
}

// CreateSensitiveRecordResponse defines model for CreateSensitiveRecordResponse.
type CreateSensitiveRecordResponse = SensitiveRecordWithData

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Errors errors must split by \r symbol
	Errors string `json:"errors"`
}

// GetSensitiveRecordResponse defines model for GetSensitiveRecordResponse.
type GetSensitiveRecordResponse = SensitiveRecordWithData

// ListSensitiveRecordResponse defines model for ListSensitiveRecordResponse.
type ListSensitiveRecordResponse struct {
	SensitiveRecords []SensitiveRecord `json:"sensitive_records"`
}

// ListSensitiveRecordTypeResponse defines model for ListSensitiveRecordTypeResponse.
type ListSensitiveRecordTypeResponse struct {
	SensitiveRecordTypes []SensitiveRecordType `json:"sensitive_record_types"`
}

// LoginUserRequest defines model for LoginUserRequest.
type LoginUserRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginUserResponse defines model for LoginUserResponse.
type LoginUserResponse struct {
	Token string `json:"token"`
}

// RegisterUserRequest defines model for RegisterUserRequest.
type RegisterUserRequest struct {
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
	Username         string `json:"username"`
}

// SensitiveRecord defines model for SensitiveRecord.
type SensitiveRecord struct {
	Id       int                     `json:"id"`
	Metadata string                  `json:"metadata"`
	Type     SensitiveRecordTypeEnum `json:"type"`
}

// SensitiveRecordType defines model for SensitiveRecordType.
type SensitiveRecordType struct {
	Id   int                     `json:"id"`
	Name SensitiveRecordTypeEnum `json:"name"`
}

// SensitiveRecordTypeEnum defines model for SensitiveRecordTypeEnum.
type SensitiveRecordTypeEnum string

// SensitiveRecordWithData defines model for SensitiveRecordWithData.
type SensitiveRecordWithData struct {
	Data     openapi_types.File      `json:"data"`
	Id       int                     `json:"id"`
	Metadata string                  `json:"metadata"`
	Type     SensitiveRecordTypeEnum `json:"type"`
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = LoginUserRequest

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody = RegisterUserRequest

// PostSensitiveRecordJSONRequestBody defines body for PostSensitiveRecord for application/json ContentType.
type PostSensitiveRecordJSONRequestBody = CreateSensitiveRecordRequest
