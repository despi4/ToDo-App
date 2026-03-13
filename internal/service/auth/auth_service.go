package authsvc

import (
	"context"
	"strings"
	authdomain "todo-app/internal/domain/auth"
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

func (auth *AuthService) Register(ctx context.Context, input authdomain.RegisterUser) error {
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

func (auth *AuthService) Login(ctx context.Context, email, password string) (authdomain.TokenPair, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	var tokenPair authdomain.TokenPair

	if email == "" || password == "" {
		return tokenPair, userdomain.ErrInvalidArgument
	}

	user, err := auth.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return tokenPair, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return tokenPair, userdomain.ErrUnauthorized
	}

	return tokenPair, nil
}

func (auth *AuthService) ChangePassword(ctx context.Context, userID uuid.UUID, old_password, new_password string) error {
	user, err := auth.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if old_password == "" || new_password == "" {
		return userdomain.ErrInvalidArgument
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

	new_hash, err := hashPassword(new_password)
	if err != nil {
		return err
	}

	if err = auth.repo.UpdatePasswordHash(ctx, userID, userdomain.PasswordHash(new_hash)); err != nil {
		return err
	}

	return nil
}

func RefreshToken(ctx context.Context, refreshToken string) (authdomain.TokenPair, error) {
	var tokenPair authdomain.TokenPair
	return tokenPair, nil
}

func passwordValidation(password string) error {
	if len(password) == 0 {
		return userdomain.ErrInvalidArgument
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
