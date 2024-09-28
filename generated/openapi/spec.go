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

	"H4sIAAAAAAAC/+xYUW/bNhD+KwQ3YC+K7TTZw/y0rk2HbMNaNO06IDUKWjpb7CSSPZ6cGIH++0BSsi1L",
	"stO0djt0TzYo8e54930f73THY50brUCR5eM7buMUcuH/PkEQBFegrCS5gJcQa0xewocCLLnnBrUBJAn+",
	"7RxIJIKE+w+3IjcZ8DG3Ogcm1UxjLkhqxcRUF8RsbZSht8ojTkvjNxBKNedlvXDHv0eY8TH/brgOdFhF",
	"OdwK7tXSwIUqcl6WEUf4UEiEhI+vg61oHeNk5U9P30NMzl/Pca3RykL7vDJpnPR0tDIpFcEc0Nn8qpMi",
	"Vx72ZOYCUWN/JsA99v8SsDFK487Ex9U6ywtLzJpMEpsu2du3yOwyn+qsfbyt+Cq7XRH9IS3du1KrvL4L",
	"eQ3lI8jtR2ZyXQAuEMWyFXHbU2fwei7VawvYyyUjrL1xDhuwqVdPu4BRWEAlcmhuqVdP9yZ7tT9ae98T",
	"fF++Sf8Dyv/Z6TG81uXjJcylJcDPnSME4yievPtK8tsVUVc+tmH4LWtRn0F3KP97zY1Y5qDoJBYhbrgl",
	"HvGpVAKXPOIxQgKKpMg2HGxUzUJcoKTllTtCyK8oKHW/UxAI+Mznio/5b29e8SjcmM5GeLrOVEpkeOlM",
	"uvy6/bFWJGIPZsiFzFyB5hoHciGUXgwk/exWB1i4k2YyhopfAXj8mUYGSRH7QvkKSPJFvLoR8zkg+1Wb",
	"9HcAA8hO2HMD6vGLS3Y2GPGILwBtkOXTwWgwcru1ASWM5GN+NhgNzjwyKfUHHmaO5h5oOpCvKe1eBZjD",
	"NJOKUQqbrhdSsBruTKiEbQDeodaHf5nwMX+hLXlTPKABLP2ik2WdKlDetTAmk+HQw/dWq3Wbsg99LaUt",
	"m7gjLMAvBCnzR380Gh3CfyWWPoBmLq+KOAZrZ0XGVtnhEU9BJBCu1b9PLm6NRLAnj2cE2C5HIsiRmL1+",
	"9YTdpKCY11YGYVeNUR/urAav23NC0ivSNglckOefMQ/NBqIjB5dqITKZrGAzrCHDbOG8gr97z0enxwwp",
	"1ogQ0wq+TOMa14kGq34gBrfSemH6MeRr2wi59zPmexmmFbOAC6fHTmiKPHeStEknVwwxt/WNYfnEvTnE",
	"6kbsZ2R9Z34eUtbWDsTLrgv+/tTspU+dpRotXwDAUpmCgvefjuf9woEQ1vT5VDw2wNQDyc6ueg6d0KQC",
	"lWWC+a6Z6VmrvbBuMS4QQVHttInKjo7f8kNK944J474ifmzJ+lMz16m45iY4CPfAA9FQtUF8fF03QNeT",
	"ctKESShsq5obkGk/m5RRj4a9yEQMTDAFNy2bbUnrFK7tVvkw+rXzy0hZKVkDmKeH9v0QaH5Zffw6iHH+",
	"6NHx4vjL5SAEALcxmLoUB2FnAEoPnfYwtFPgh3cyKUOcGRC0I37q15nIMubmty6hZzeSUuYNbdM37N4C",
	"9htJ6eVTP5qgyIF8R3x9x91w4scVHtXTkazZvm4gNhvf7YHYJWuLpOc7u4tvR9erOu6s3W5939EHPAQa",
	"z4Di9EsiY5dS6piATiwhiLxZ9NWstfr20Bq0/mt9xHkgydHi6AZJ3SnK5OC9jUfrTOOnsKG723ltMi0S",
	"FrDRzwrSD+h7ngoSh6PGfRqqT2JFT//0f1/zMUQ9O2ocKOcp2RqtrXbj2BPx4wxBJMvweWbPlXMENbkf",
	"1/c1Zd6XcxvoXGBWfWW24+Ew07HIUm1pKIzkzn1la/ssz2uy2OpbfzVvV4oQZvwy2rHNZ61r6KtMtB+V",
	"k/LfAAAA//9D8F9paR4AAA==",
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
