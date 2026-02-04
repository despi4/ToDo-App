package userdomain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByID(ctx context.Context, ID uuid.UUID) (user *User, err error)
	GetUserByEmail(ctx context.Context, email string) (user *User, err error)
	UpdateUser(ctx context.Context, ID uuid.UUID, userUpdate UserUpdate) (User, error)
	DeleteUser(ctx context.Context, ID uuid.UUID) error
}

type UserService interface {
	Create(ctx context.Context, user User) (User, error)
	GetByID(ctx context.Context, ID uuid.UUID) (user *User, err error)
	GetByEmail(ctx context.Context, email string) (user *User, err error)
	Update(ctx context.Context, ID uuid.UUID, userUpdate UserUpdate) (User, error)
	Delete(ctx context.Context, ID uuid.UUID) error
}
