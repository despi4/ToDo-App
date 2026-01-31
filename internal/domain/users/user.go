package users

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// User model
type User struct {
	ID        uuid.UUID
	Name      string
	Surname   string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Domain Errors
var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argumenmt")
	ErrConflict        = errors.New("conflict")
	ErrForbidden       = errors.New("forbidden")
)
