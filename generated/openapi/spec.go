// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/Go2mv0dAAD//56C8ZtrLQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}