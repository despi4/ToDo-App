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

type PasswordHash string

// User model
type User struct {
	ID           uuid.UUID
	Role         Role
	Name         string
	Surname      string
	Email        string
	PasswordHash PasswordHash
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Domain Errors
var (
	ErrShortPassword     = errors.New("password must contain at least 6 characters")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrNotFound          = errors.New("not found")
	ErrInvalidArgument   = errors.New("invalid argument")
	ErrConflict          = errors.New("conflict")
	ErrForbidden         = errors.New("forbidden")
	ErrAlreadyExists     = errors.New("already exists")
	ErrEmailTaken        = errors.New("email taken")
	ErrInvalidCredential = errors.New("invalid credential")
	ErrSamePassword = errors.New("same password")
)
