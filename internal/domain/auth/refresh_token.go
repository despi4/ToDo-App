package authdomain

import (
	"time"

	"github.com/google/uuid"
)

type TokenHash string

type RefreshToken struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	TokenHash  TokenHash
	ExpiresAt  time.Time
	Revoked    bool
	ReplacedBy *TokenHash
	CreatedAT  time.Time
}
