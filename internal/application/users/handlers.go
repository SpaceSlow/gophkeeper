package users

import "context"

type Repository interface {
	ExistUsername(ctx context.Context, username string) (bool, error)
	RegisterUser(ctx context.Context, username, passwordHash string) error
	FetchPasswordHash(ctx context.Context, username string) (string, error)
	FetchUserID(ctx context.Context, username string) (int, error)
	Close()
}

type UserHandlers struct {
	repo Repository
}

func SetupHandlers(repo Repository) UserHandlers {
	return UserHandlers{
		repo: repo,
	}
}
