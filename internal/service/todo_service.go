package service

import (
	"errors"
	"strings"
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

// Create
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

// Update
func (s *TodoService) MarkIsDone(title string) error {
	return nil
}

// Получение списка задач с фильтром Read
func (s *TodoService) GetTodos(filter string) ([]models.Todo, error) {
	// 1. Получить все задачи repo.GetAll()
	// 2. Отфильтровать (all, active, completed)
	// 3. Вернуть результат

	todoList := s.repo.GetAll()

	if strings.EqualFold("all", filter) || strings.EqualFold("completed", filter) || strings.EqualFold("active", filter) {
		return nil, errors.New("incorrect filter")
	}

	var newTodoList []models.Todo

	if strings.EqualFold("all", filter) {
		return todoList, nil
	} else if strings.EqualFold("active", filter) {
		for _, todo := range todoList {
			if !todo.Completed {
				newTodoList = append(todoList, todo)
			}
		}

		return newTodoList, nil
	}

	for _, todo := range todoList {
		if todo.Completed {
			newTodoList = append(newTodoList, todo)
		}
	}

	return todoList, nil
}
