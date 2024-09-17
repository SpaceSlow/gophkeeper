package crypto

import (
	"errors"
	"fmt"
)

var ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
var ErrInvalidToken = errors.New("token is not valid")
var ErrInvalidPasswordHash = errors.New("invalid password hash layout")

type UnknownHashAlgError struct {
	Alg string
}

func (e UnknownHashAlgError) Error() string {
	return fmt.Sprintf("unknown password hash algorithm: %s", e.Alg)
}
