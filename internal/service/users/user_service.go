package usersvc

import (
	"context"
	"todo-app/internal/domain/users"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user users.User) (users.User, error)
	GetUserByID(ctx context.Context, ID uuid.UUID) (user *users.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user *users.User, err error)
	UpdateUser(ctx context.Context, ID uuid.UUID, userUpdate users.UserUpdate) (users.User, error)
	DeleteUser(ctx context.Context, ID uuid.UUID) error
}

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
