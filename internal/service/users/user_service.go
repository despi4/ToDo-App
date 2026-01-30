package usersvc

import "todo-app/internal/domain/users"

type UserService struct {
	repo *users.UserRepository
}

func NewUserService(repo *users.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
