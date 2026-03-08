package usersvc

import (
	"context"
	"strings"
	userdomain "todo-app/internal/domain/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo userdomain.UserRepository
}

func NewAuthService(repo userdomain.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (auth *AuthService) Register(ctx context.Context, input userdomain.RegisterUser) error {
	user := userdomain.User{
		Name:    strings.TrimSpace(input.Name),
		Surname: strings.TrimSpace(input.Surname),
		Email:   strings.ToLower(strings.TrimSpace(input.Email)),
	}

	if user.Name == "" || user.Surname == "" || user.Email == "" {
		return userdomain.ErrInvalidArgument
	}

	err := passwordValidation(input.Password)
	if err != nil {
		return err
	}

	hash_password, err := hashPassword(input.Password)
	if err != nil {
		return err
	}

	user.PasswordHash = userdomain.PasswordHash(hash_password)

	_, err = auth.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (auth *AuthService) Login(ctx context.Context, email, password string) error {
	email = strings.ToLower(strings.TrimSpace(email))

	if email == "" {
		return userdomain.ErrInvalidArgument
	}

	user, err := auth.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return userdomain.ErrUnauthorized
	}

	return nil
}

func (auth *AuthService) ChangePassword(ctx context.Context, ID uuid.UUID, old_password, new_password string) error {
	user, err := auth.repo.GetUserByID(ctx, ID)
	if err != nil {
		return err
	}

	if old_password == new_password {
		return userdomain.ErrSamePassword
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(old_password)); err != nil {
		return userdomain.ErrInvalidCredential
	}

	err = passwordValidation(new_password)
	if err != nil {
		return err
	}

	new_hash, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err = auth.repo.UpdatePasswordHash(ctx, ID, userdomain.PasswordHash(new_hash)); err != nil {
		return err
	}

	return nil
}

func passwordValidation(password string) error {
	if len(password) == 0 {
		return userdomain.ErrInvalidCredential
	} else if len(password) < 6 {
		return userdomain.ErrShortPassword
	}

	return nil
}

func hashPassword(password string) (string, error) {
	// bcrypt cost в Go — это параметр, который задаёт насколько “дорого” (медленно) будет вычисляться хэш пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
