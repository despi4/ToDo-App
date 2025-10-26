package service

import (
	"errors"
	"todo-app/internal/models"
	"todo-app/internal/repository"
)

// Бизнес-логика приложения
// Этот файл предназначен для работы с задачами (create, read, update, delete)

type TodoService struct {
	repo repository.TodoRepository
}

// Конструктор для инициализации сервиса
func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(title string) error {
	// 1. Валидация (title не должен быть пустым)
	// 2. Создание структуры Todo
	// 3. Вызов repo.Create(todo)
	// 4. Возврат ошибки или nil
	return errors.New("not implemented yet")
}

// Получение списка задач с фильтром
func (s *TodoService) GetTodos(filter string) ([]models.Todo, error) {
	// 1. Получить все задачи repo.GetAll()
	// 2. Отфильтровать (all, active, completed)
	// 3. Вернуть результат
	return nil, errors.New("not implemented yet")
}
