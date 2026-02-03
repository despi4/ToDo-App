package todos

import (
	"context"

	"github.com/google/uuid"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo Todo) (Todo, error)
	GetTodoByID(ctx context.Context, userID, ID uuid.UUID) (todo *Todo, err error)
	GetTodos(ctx context.Context, userID uuid.UUID, todoFilter TodoFilter) (todos []Todo, err error)
	UpdateTodo(ctx context.Context, userID uuid.UUID, ID uuid.UUID, todoUpdate TodoUpdate) (Todo, error)
	DeleteTodo(ctx context.Context, userID uuid.UUID, ID uuid.UUID) error
}
