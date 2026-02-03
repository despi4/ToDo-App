package usersvc

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"todo-app/internal/domain/user"

	"github.com/google/uuid"
)

var emailRegex = regexp.MustCompile(`(?i)^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)

type UserService struct {
	repo users.UserRepository
}

func NewUserService(repo users.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (userSvc *UserService) Create(ctx context.Context, user users.User) (users.User, error) {
	if err := validateUser(&user.Name, &user.Surname, &user.Email); err != nil {
		return users.User{}, err
	}

	user, err := userSvc.repo.CreateUser(ctx, user)
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}

func (userSvc *UserService) GetByID(ctx context.Context, ID uuid.UUID) (user *users.User, err error) {
	if ID == uuid.Nil {
		return nil, users.ErrInvalidArgument
	}

	user, err = userSvc.repo.GetUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userSvc *UserService) GetByEmail(ctx context.Context, email string) (user *users.User, err error) {
	err = validateUser(nil, nil, &email)
	if err != nil {
		return nil, err
	}

	user, err = userSvc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userSvc *UserService) Update(ctx context.Context, ID uuid.UUID, userUpdate users.UserUpdate) (users.User, error) {
	if ID == uuid.Nil {
		return users.User{}, users.ErrInvalidArgument
	}

	if userUpdate.Name == nil && userUpdate.Email == nil && userUpdate.Surname == nil {
		user, err := userSvc.GetByID(ctx, ID)
		if err != nil {
			return users.User{}, err
		}

		return *user, nil
	}

	if err := validateUser(userUpdate.Name, userUpdate.Surname, userUpdate.Email); err != nil {
		return users.User{}, err
	}

	user, err := userSvc.repo.UpdateUser(ctx, ID, userUpdate)
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}

func (userSvc *UserService) Delete(ctx context.Context, ID uuid.UUID) error {
	if ID == uuid.Nil {
		return users.ErrInvalidArgument
	}

	err := userSvc.repo.DeleteUser(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func validateUser(name, surname, email *string) error {
	if name != nil {
		*name = strings.TrimSpace(*name)
		if len(*name) == 0 || len(*name) > 30 {
			return fmt.Errorf("name is incorrect: %w", users.ErrInvalidArgument)
		}
	}

	if surname != nil {
		*surname = strings.TrimSpace(*surname)
		if len(*surname) == 0 || len(*surname) > 30 {
			return fmt.Errorf("surname is incorrect: %w", users.ErrInvalidArgument)
		}
	}

	if email != nil {
		*email = strings.TrimSpace(*email)
		switch {
		case len(*email) == 0:
			return fmt.Errorf("email is incorrect: %w", users.ErrInvalidArgument)
		case len(*email) > 254:
			return fmt.Errorf("email is incorrect: %w", users.ErrInvalidArgument)
		case !emailRegex.Match([]byte(*email)):
			return fmt.Errorf("email is incorrect: %w", users.ErrInvalidArgument)
		}
	}

	return nil
}
