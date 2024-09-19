package users

type Repository interface {
	ExistUsername(username string) (bool, error)
	RegisterUser(username, passwordHash string) error
	FetchPasswordHash(username string) (string, error)
	FetchUserID(username string) (int, error)
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
