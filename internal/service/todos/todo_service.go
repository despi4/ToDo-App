package todosvc

import "todo-app/internal/domain/todos"

type TodoService struct {
	repo *todos.TodoRepository
}

func NewTodoRepository(repo *todos.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}
