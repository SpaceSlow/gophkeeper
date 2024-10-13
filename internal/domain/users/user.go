package users

import (
	"fmt"

	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

type User struct {
	id               int
	username         string
	password         string
	repeatedPassword string
	passwordHash     string
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

func NewUserWithHash(id int, name, passwordHash string) (*User, error) {
	return &User{
		id:           id,
		username:     name,
		passwordHash: passwordHash,
	}, nil
}

func CreateUser(name, password string) (*User, error) {
	return NewUser(0, name, password)
}

func (u *User) Id() int {
	return u.id
}

func (u *User) GeneratePasswordHash(keyLen, passwordIterationNum int) (string, error) {
	return crypto.GenerateHash(u.password, keyLen, passwordIterationNum)
}

func (u *User) CheckPasswordHash(password string, keyLen int) (bool, error) {
	return crypto.IsValid(password, u.passwordHash, keyLen)
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
