package service

import (
	"errors"
	"time"
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

	if len(title) != 0 {
		// нельзя создавать задачу без названия
		return errors.New("title can not be empty")
	} else if len(title) > 100 {
		return errors.New("title length can not be over 100")
	}

	newTodo := models.Todo{
		Id:        1,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	s.repo.Create(&newTodo)

	return nil
}

// Получение списка задач с фильтром
func (s *TodoService) GetTodos(filter string) ([]models.Todo, error) {
	// 1. Получить все задачи repo.GetAll()
	// 2. Отфильтровать (all, active, completed)
	// 3. Вернуть результат
	return nil, errors.New("not implemented yet")
}
