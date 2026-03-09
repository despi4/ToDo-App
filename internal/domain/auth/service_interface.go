package authdomain

import (
	"context"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(ctx context.Context, input RegisterUser) error
	Login(ctx context.Context, email string, password string) (TokenPair, error)
	ChangePassword(ctx context.Context, userID uuid.UUID, old_password, new_password string) error

	// должен принять refresh token и вернуть новую пару
	RefreshToken(ctx context.Context, refreshToken string) (TokenPair, error)
}
