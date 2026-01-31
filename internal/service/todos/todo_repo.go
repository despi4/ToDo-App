package todosvc

import (
	"context"
	"todo-app/internal/domain/todos"

	"github.com/google/uuid"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo todos.Todo) (Todo, err error)
	GetTodoByID(ctx context.Context, userID, ID uuid.UUID) (todo *todos.Todo, err error)
	GetTodos(ctx context.Context, userID uuid.UUID, todoFilter TodoFilter) (todos []todos.Todo, err error)
	UpdateTodo(ctx context.Context, userID uuid.UUID, ID uuid.UUID, todoUpdate TodoUpdate) error
	DeleteTodo(ctx context.Context, userID uuid.UUID, ID uuid.UUID) error
}

type TodoUpdate struct {
	Status      todos.TodoStatus
	Title       string
	Description string
}

type TodoFilter struct {
	Search string
	Limit  string
	Offset string
	Status todos.TodoStatus
}
