// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

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

// CreateSensitiveRecordDataResponse defines model for CreateSensitiveRecordDataResponse.
type CreateSensitiveRecordDataResponse struct {
	Message string `json:"message"`
}

// CreateSensitiveRecordRequest defines model for CreateSensitiveRecordRequest.
type CreateSensitiveRecordRequest struct {
	Metadata string                  `json:"metadata"`
	Type     SensitiveRecordTypeEnum `json:"type"`
}

// CreateSensitiveRecordResponse defines model for CreateSensitiveRecordResponse.
type CreateSensitiveRecordResponse = SensitiveRecord

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Errors errors must split by \r symbol
	Errors string `json:"errors"`
}

// ListSensitiveRecordResponse defines model for ListSensitiveRecordResponse.
type ListSensitiveRecordResponse struct {
	SensitiveRecords []SensitiveRecord `json:"sensitive_records"`
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

// SensitiveRecordTypeEnum defines model for SensitiveRecordTypeEnum.
type SensitiveRecordTypeEnum string

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = LoginUserRequest

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody = RegisterUserRequest

// PostSensitiveRecordJSONRequestBody defines body for PostSensitiveRecord for application/json ContentType.
type PostSensitiveRecordJSONRequestBody = CreateSensitiveRecordRequest
