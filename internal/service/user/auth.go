package usersvc

import (
	"context"
	userdomain "todo-app/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

// Composition

type AuthService struct {
	repo userdomain.UserRepository
}

func NewAuthService(repo userdomain.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (auth *AuthService) Register(ctx context.Context, input userdomain.RegisterUser) (userdomain.User, error) {
	return userdomain.User{}, nil
}

func (auth *AuthService) Login(ctx context.Context, email string, password string) (string, error) {
	return email, nil
}

func HashPassword(password string) (string, error) {
	// bcrypt cost в Go — это параметр, который задаёт насколько “дорого” (медленно) будет вычисляться хэш пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
