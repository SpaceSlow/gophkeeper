package users

import (
	"errors"
	"fmt"
)

var ErrUserValidation = errors.New("validation error")
var ErrUserLogin = errors.New("incorrect username or password")

type NoUserError struct {
	username string
}

func NewNoUserError(username string) error {
	return NoUserError{username: username}
}

func (e NoUserError) Error() string {
	return fmt.Sprintf("user %s does not exist", e.username)
}

type RegisteredUserError struct {
	username string
}

func NewRegisteredUserError(username string) error {
	return RegisteredUserError{username: username}
}

func (e RegisteredUserError) Error() string {
	return fmt.Sprintf("user %s is already registered", e.username)
}
