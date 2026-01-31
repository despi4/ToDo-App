package users

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByID(ctx context.Context) error
	GetUserByEmail(ctx context.Context) error
	UpdateUser(ctx context.Context, name string, surname string, email string) error
	DeleteUser(ctx context.Context) error
}
