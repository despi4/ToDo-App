package postgre

import (
	"context"
	"os/user"
	usersvc "todo-app/internal/service/users"

	"github.com/google/uuid"
)

// Repository — это не “часть базы данных”
// Repository — это Adapter, который делает базу данных совместимой с бизнес-логикой

type UserRepo struct {
	db *DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user user.User) (user.User, error) {
	var err error

	return user, err
}

func (u *UserRepo) GetUserByID(ctx context.Context, ID uuid.UUID) (*user.User, error)

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*user.User, error)

func (u *UserRepo) UpdateUser(ctx context.Context, ID uuid.UUID, userUpdate usersvc.UserUpdate) (user.User, error)

func (u *UserRepo) DeleteUser(ctx context.Context, ID uuid.UUID) error 
