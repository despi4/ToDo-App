package users

import (
	"errors"

	"github.com/google/uuid"
)

// User model
type User struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Email   string
}

// Domain Errors
var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argumenmt")
)
