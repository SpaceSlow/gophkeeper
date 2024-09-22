package users

import (
	"errors"
	"fmt"
)

var ErrUserValidation = errors.New("validation error")

type NoUserError struct {
	Username string
}

func (e NoUserError) Error() string {
	return fmt.Sprintf("user %s does not exist", e.Username)
}
