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

type Todo struct {
	Id        uuid.UUID
	Name      string
	Task      string
	Status    TodoStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argumenmt")
)
