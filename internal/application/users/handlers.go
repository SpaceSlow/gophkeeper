package users

import (
	"time"

	"github.com/SpaceSlow/gophkeeper/internal/domain/users"
)

type Repository interface {
	ExistUsername(username string) (bool, error)
	RegisterUser(username, passwordHash string) error
	FetchUser(username string) (*users.User, error)
	ExistUser(userID int) (bool, error)
	Close()
}

type ConfigProvider interface {
	SecretKey() string
	KeyLen() int
	PasswordIterationNum() int
	TokenLifetime() time.Duration
}

type UserHandlers struct {
	repo Repository
	cfg  ConfigProvider
}

func SetupHandlers(repo Repository, cfg ConfigProvider) UserHandlers {
	return UserHandlers{
		repo: repo,
		cfg:  cfg,
	}
}
