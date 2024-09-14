package store

import "fmt"

type NoUserError struct {
	Username string
}

func (e NoUserError) Error() string {
	return fmt.Sprintf("user %s does not exist", e.Username)
}
