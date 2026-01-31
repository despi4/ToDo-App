package todos

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TodoStatus string

const (
	TodoStatusNotStarted TodoStatus = "Not Started"
	TodoStatusInProgress TodoStatus = "In Progress"
	TodoStatusDone       TodoStatus = "Done"
)

// Todo model
type Todo struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Title     string
	Task      string
	Status    TodoStatus
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
