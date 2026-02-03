package todosvc

import (
	"context"
	"todo-app/internal/domain/todos"

	"github.com/google/uuid"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo todos.Todo) (todos.Todo, error)
	GetTodoByID(ctx context.Context, userID, ID uuid.UUID) (todo *todos.Todo, err error)
	GetTodos(ctx context.Context, userID uuid.UUID, todoFilter todos.TodoFilter) (todos []todos.Todo, err error)
	UpdateTodo(ctx context.Context, userID uuid.UUID, ID uuid.UUID, todoUpdate todos.TodoUpdate) (todos.Todo, error)
	DeleteTodo(ctx context.Context, userID uuid.UUID, ID uuid.UUID) error
}

type TodoService struct {
	repo *TodoRepository
}

func NewTodoRepository(repo *TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}
