// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostLoginWithBody request with any body
	PostLoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostLogin(ctx context.Context, body PostLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostRegisterWithBody request with any body
	PostRegisterWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostRegister(ctx context.Context, body PostRegisterJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListSensitiveRecords request
	ListSensitiveRecords(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostSensitiveRecordWithBody request with any body
	PostSensitiveRecordWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostSensitiveRecord(ctx context.Context, body PostSensitiveRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteSensitiveRecordWithID request
	DeleteSensitiveRecordWithID(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FetchSensitiveRecordWithID request
	FetchSensitiveRecordWithID(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostSensitiveRecordDataWithBody request with any body
	PostSensitiveRecordDataWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostLoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostLoginRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostLogin(ctx context.Context, body PostLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostLoginRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostRegisterWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostRegisterRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostRegister(ctx context.Context, body PostRegisterJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostRegisterRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListSensitiveRecords(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListSensitiveRecordsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSensitiveRecordWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSensitiveRecordRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSensitiveRecord(ctx context.Context, body PostSensitiveRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSensitiveRecordRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteSensitiveRecordWithID(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteSensitiveRecordWithIDRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FetchSensitiveRecordWithID(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFetchSensitiveRecordWithIDRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSensitiveRecordDataWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSensitiveRecordDataRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostLoginRequest calls the generic PostLogin builder with application/json body
func NewPostLoginRequest(server string, body PostLoginJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostLoginRequestWithBody(server, "application/json", bodyReader)
}

// NewPostLoginRequestWithBody generates requests for PostLogin with any type of body
func NewPostLoginRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/login")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPostRegisterRequest calls the generic PostRegister builder with application/json body
func NewPostRegisterRequest(server string, body PostRegisterJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostRegisterRequestWithBody(server, "application/json", bodyReader)
}

// NewPostRegisterRequestWithBody generates requests for PostRegister with any type of body
func NewPostRegisterRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/register")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListSensitiveRecordsRequest generates requests for ListSensitiveRecords
func NewListSensitiveRecordsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/sensitive_records")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostSensitiveRecordRequest calls the generic PostSensitiveRecord builder with application/json body
func NewPostSensitiveRecordRequest(server string, body PostSensitiveRecordJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostSensitiveRecordRequestWithBody(server, "application/json", bodyReader)
}

// NewPostSensitiveRecordRequestWithBody generates requests for PostSensitiveRecord with any type of body
func NewPostSensitiveRecordRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/sensitive_records")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteSensitiveRecordWithIDRequest generates requests for DeleteSensitiveRecordWithID
func NewDeleteSensitiveRecordWithIDRequest(server string, id int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/sensitive_records/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFetchSensitiveRecordWithIDRequest generates requests for FetchSensitiveRecordWithID
func NewFetchSensitiveRecordWithIDRequest(server string, id int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/sensitive_records/%s/data", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostSensitiveRecordDataRequestWithBody generates requests for PostSensitiveRecordData with any type of body
func NewPostSensitiveRecordDataRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/sensitive_records/%s/data", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// PostLoginWithBodyWithResponse request with any body
	PostLoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostLoginResponse, error)

	PostLoginWithResponse(ctx context.Context, body PostLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*PostLoginResponse, error)

	// PostRegisterWithBodyWithResponse request with any body
	PostRegisterWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostRegisterResponse, error)

	PostRegisterWithResponse(ctx context.Context, body PostRegisterJSONRequestBody, reqEditors ...RequestEditorFn) (*PostRegisterResponse, error)

	// ListSensitiveRecordsWithResponse request
	ListSensitiveRecordsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListSensitiveRecordsResponse, error)

	// PostSensitiveRecordWithBodyWithResponse request with any body
	PostSensitiveRecordWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSensitiveRecordResponse, error)

	PostSensitiveRecordWithResponse(ctx context.Context, body PostSensitiveRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSensitiveRecordResponse, error)

	// DeleteSensitiveRecordWithIDWithResponse request
	DeleteSensitiveRecordWithIDWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteSensitiveRecordWithIDResponse, error)

	// FetchSensitiveRecordWithIDWithResponse request
	FetchSensitiveRecordWithIDWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*FetchSensitiveRecordWithIDResponse, error)

	// PostSensitiveRecordDataWithBodyWithResponse request with any body
	PostSensitiveRecordDataWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSensitiveRecordDataResponse, error)
}

type PostLoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LoginUserResponse
	JSON400      *ErrorResponse
	JSON401      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostLoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostLoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostRegisterResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *ErrorResponse
	JSON409      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostRegisterResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostRegisterResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListSensitiveRecordsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListSensitiveRecordResponse
	JSON401      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r ListSensitiveRecordsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListSensitiveRecordsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostSensitiveRecordResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *CreateSensitiveRecordResponse
	JSON400      *ErrorResponse
	JSON401      *ErrorResponse
	JSON422      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostSensitiveRecordResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostSensitiveRecordResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteSensitiveRecordWithIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON401      *ErrorResponse
	JSON403      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r DeleteSensitiveRecordWithIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteSensitiveRecordWithIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type FetchSensitiveRecordWithIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON401      *ErrorResponse
	JSON403      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r FetchSensitiveRecordWithIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FetchSensitiveRecordWithIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostSensitiveRecordDataResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *ErrorResponse
	JSON401      *ErrorResponse
	JSON403      *ErrorResponse
	JSON409      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostSensitiveRecordDataResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostSensitiveRecordDataResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostLoginWithBodyWithResponse request with arbitrary body returning *PostLoginResponse
func (c *ClientWithResponses) PostLoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostLoginResponse, error) {
	rsp, err := c.PostLoginWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostLoginResponse(rsp)
}

func (c *ClientWithResponses) PostLoginWithResponse(ctx context.Context, body PostLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*PostLoginResponse, error) {
	rsp, err := c.PostLogin(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostLoginResponse(rsp)
}

// PostRegisterWithBodyWithResponse request with arbitrary body returning *PostRegisterResponse
func (c *ClientWithResponses) PostRegisterWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostRegisterResponse, error) {
	rsp, err := c.PostRegisterWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostRegisterResponse(rsp)
}

func (c *ClientWithResponses) PostRegisterWithResponse(ctx context.Context, body PostRegisterJSONRequestBody, reqEditors ...RequestEditorFn) (*PostRegisterResponse, error) {
	rsp, err := c.PostRegister(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostRegisterResponse(rsp)
}

// ListSensitiveRecordsWithResponse request returning *ListSensitiveRecordsResponse
func (c *ClientWithResponses) ListSensitiveRecordsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListSensitiveRecordsResponse, error) {
	rsp, err := c.ListSensitiveRecords(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListSensitiveRecordsResponse(rsp)
}

// PostSensitiveRecordWithBodyWithResponse request with arbitrary body returning *PostSensitiveRecordResponse
func (c *ClientWithResponses) PostSensitiveRecordWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSensitiveRecordResponse, error) {
	rsp, err := c.PostSensitiveRecordWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSensitiveRecordResponse(rsp)
}

func (c *ClientWithResponses) PostSensitiveRecordWithResponse(ctx context.Context, body PostSensitiveRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSensitiveRecordResponse, error) {
	rsp, err := c.PostSensitiveRecord(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSensitiveRecordResponse(rsp)
}

// DeleteSensitiveRecordWithIDWithResponse request returning *DeleteSensitiveRecordWithIDResponse
func (c *ClientWithResponses) DeleteSensitiveRecordWithIDWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteSensitiveRecordWithIDResponse, error) {
	rsp, err := c.DeleteSensitiveRecordWithID(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteSensitiveRecordWithIDResponse(rsp)
}

// FetchSensitiveRecordWithIDWithResponse request returning *FetchSensitiveRecordWithIDResponse
func (c *ClientWithResponses) FetchSensitiveRecordWithIDWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*FetchSensitiveRecordWithIDResponse, error) {
	rsp, err := c.FetchSensitiveRecordWithID(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFetchSensitiveRecordWithIDResponse(rsp)
}

// PostSensitiveRecordDataWithBodyWithResponse request with arbitrary body returning *PostSensitiveRecordDataResponse
func (c *ClientWithResponses) PostSensitiveRecordDataWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSensitiveRecordDataResponse, error) {
	rsp, err := c.PostSensitiveRecordDataWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSensitiveRecordDataResponse(rsp)
}

// ParsePostLoginResponse parses an HTTP response from a PostLoginWithResponse call
func ParsePostLoginResponse(rsp *http.Response) (*PostLoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostLoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LoginUserResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParsePostRegisterResponse parses an HTTP response from a PostRegisterWithResponse call
func ParsePostRegisterResponse(rsp *http.Response) (*PostRegisterResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostRegisterResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}

// ParseListSensitiveRecordsResponse parses an HTTP response from a ListSensitiveRecordsWithResponse call
func ParseListSensitiveRecordsResponse(rsp *http.Response) (*ListSensitiveRecordsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListSensitiveRecordsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListSensitiveRecordResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParsePostSensitiveRecordResponse parses an HTTP response from a PostSensitiveRecordWithResponse call
func ParsePostSensitiveRecordResponse(rsp *http.Response) (*PostSensitiveRecordResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostSensitiveRecordResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest CreateSensitiveRecordResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseDeleteSensitiveRecordWithIDResponse parses an HTTP response from a DeleteSensitiveRecordWithIDWithResponse call
func ParseDeleteSensitiveRecordWithIDResponse(rsp *http.Response) (*DeleteSensitiveRecordWithIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteSensitiveRecordWithIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	}

	return response, nil
}

// ParseFetchSensitiveRecordWithIDResponse parses an HTTP response from a FetchSensitiveRecordWithIDWithResponse call
func ParseFetchSensitiveRecordWithIDResponse(rsp *http.Response) (*FetchSensitiveRecordWithIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FetchSensitiveRecordWithIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParsePostSensitiveRecordDataResponse parses an HTTP response from a PostSensitiveRecordDataWithResponse call
func ParsePostSensitiveRecordDataResponse(rsp *http.Response) (*PostSensitiveRecordDataResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostSensitiveRecordDataResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}
