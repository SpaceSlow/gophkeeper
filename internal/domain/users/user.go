package users

import (
	"fmt"

	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

type User struct {
	id       int
	username string
	password string
}

func NewUser(id int, name, password string) (*User, error) {
	if err := validateUsername(name); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	return &User{
		id:       id,
		username: name,
		password: password,
	}, nil
}

func CreateUser(name, password string) (*User, error) {
	return NewUser(0, name, password)
}

func (u *User) GeneratePasswordHash(keyLen, passwordIterationNum int) (string, error) {
	return crypto.GenerateHash(u.password, keyLen, passwordIterationNum)
}

func (u *User) CheckPasswordHash(passwordHash string, keyLen int) (bool, error) {
	return crypto.IsValid(u.password, passwordHash, keyLen)
}

func validateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("%w: name is required", ErrUserValidation)
	}
	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("%w: password is required", ErrUserValidation)
	}
	return nil
}
