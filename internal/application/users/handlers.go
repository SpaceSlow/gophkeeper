package users

import "time"

type Repository interface {
	ExistUsername(username string) (bool, error)
	RegisterUser(username, passwordHash string) error
	FetchPasswordHash(username string) (string, error)
	FetchUserID(username string) (int, error)
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
