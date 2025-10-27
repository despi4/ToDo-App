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

var (
	ErrNilTitle       = errors.New("title can not be title")
	ErrOverheadTittle = errors.New("title length can not be over 100")
)

func validateTitle(title string) error {
	if len(title) == 0 {
		// нельзя создавать задачу без названия
		return ErrNilTitle
	} else if len(title) > 100 {
		return ErrOverheadTittle
	}

	return nil
}

// Create
func (s *TodoService) CreateTodo(title string) error {
	// 1. Валидация (title не должен быть пустым)
	// 2. Создание структуры Todo
	// 3. Вызов repo.Create(todo)
	// 4. Возврат ошибки или nil

	err := validateTitle(title)
	if err != nil {
		return err
	}

	newTodo := models.Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	err = s.repo.Create(&newTodo)
	if err != nil {
		return err
	}

	return nil
}

// Update
func (s *TodoService) UpdateStatusTodo(id int) error {
	newStatusTodo := models.Todo{
		Id:        id,
		Completed: true,
	}

	err := s.repo.UpdateStatus(&newStatusTodo)
	if err != nil {
		return err
	}

	return nil
}

// Получение списка задач с фильтром Read
func (s *TodoService) GetTodo(filter string) ([]models.Todo, error) {
	// 1. Получить все задачи repo.GetAll()
	// 2. Отфильтровать (all, active, completed)
	// 3. Вернуть результат

	todoList, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if !(strings.EqualFold("all", filter) || strings.EqualFold("completed", filter) || strings.EqualFold("active", filter)) {
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

func (s *TodoService) DeleteTodo(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
