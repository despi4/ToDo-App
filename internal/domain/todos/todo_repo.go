package todos

import "context"

type TodoRepository interface {
	CreateTodo(ctx context.Context) error
	GetTodoByID(ctx context.Context) error
	GetTodoByUserID(ctx context.Context) error
	GetTodos(ctx context.Context) error
	UpdateTodo(ctx context.Context, title string, status TodoStatus) error
	DeleteTodo(ctx context.Context) error
}
