package userdomain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// User model
type User struct {
	ID        uuid.UUID
	Role      Role
	Name      string
	Surname   string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Domain Errors
var (
	ErrUnauthorized    = errors.New("unauthorized")
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
	ErrForbidden       = errors.New("forbidden")
	ErrAlreadyExists   = errors.New("already exists")
	ErrEmailTaken      = errors.New("email taken")
)
