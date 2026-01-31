package todosvc

type TodoService struct {
	repo *TodoRepository
}

func NewTodoRepository(repo *TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}
