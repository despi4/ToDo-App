package userdomain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByID(ctx context.Context, ID uuid.UUID) (user *User, err error)
	GetUserByEmail(ctx context.Context, email string) (user *User, err error)
	UpdateUser(ctx context.Context, ID uuid.UUID, userUpdate UpdateUser) (User, error)
	UpdatePasswordHash(ctx context.Context, ID uuid.UUID, new_hash PasswordHash) error
	DeleteUser(ctx context.Context, ID uuid.UUID) error
}

type UserService interface {
	Create(ctx context.Context, user User, password string) (User, error)
	GetByID(ctx context.Context, ID uuid.UUID) (user *User, err error)
	GetByEmail(ctx context.Context, email string) (user *User, err error)
	Update(ctx context.Context, ID uuid.UUID, userUpdate UpdateUser) (User, error)
	UpdatePassword(ctx context.Context, ID uuid.UUID, old_password, new_password string) error
	Delete(ctx context.Context, ID uuid.UUID) error
}
