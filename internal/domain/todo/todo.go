package todos

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TodoStatus string

const (
	TodoStatusNotStarted TodoStatus = "not_started"
	TodoStatusInProgress TodoStatus = "in_progress"
	TodoStatusDone       TodoStatus = "done"
)

// Todo model
type Todo struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Title       string
	Description string
	Status      TodoStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Domain Errors
var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
	ErrForbidden       = errors.New("forbidden")
)
