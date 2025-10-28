package repository

import (
	"errors"
	"log"

	"todo-app/internal/models"
)

// Отделить логику хранения данных от бизнес-логики

type TodoRepository interface {
	Create(todo *models.Todo) error       // create todo
	GetAll() ([]models.Todo, error)       // Get todo
	GetById(id int) (*models.Todo, error) // search todo by id
	UpdateStatus(todo *models.Todo) error // update todo
	Delete(id int) error                  // delete todo
}

var (
	ErrNotFound = errors.New("not found")
	ErrNilInput = errors.New("nil input")
)

// type Database map[int]*models.Todo
type Database struct {
	data   map[int]*models.Todo
	lastId int
}

func NewDatabase() *Database {
	return &Database{
		data:   make(map[int]*models.Todo),
		lastId: 0,
	}
}

// Method Create for create new todo
func (db *Database) Create(todo *models.Todo) error {
	if todo == nil {
		return ErrNilInput
	}

	(*db).lastId++
	todo.Id = (*db).lastId
	(*db).data[(*db).lastId] = todo

	log.Println((*db).data)

	return nil
}

// GetAll for get slice of todo
func (db *Database) GetAll() ([]models.Todo, error) {
	var todoList []models.Todo

	for _, data := range (*db).data {
		todoList = append(todoList, *data)
	}

	return todoList, nil
}

// GetById method return pointer of todo
func (db *Database) GetById(id int) (*models.Todo, error) {
	if _, ok := (*db).data[id]; ok {
		return (*db).data[id], nil
	}

	return nil, ErrNotFound
}

func (db *Database) UpdateStatus(todo *models.Todo) error {
	if todo == nil {
		return ErrNilInput
	}

	if _, ok := (*db).data[todo.Id]; ok {
		(*db).data[todo.Id] = &models.Todo{
			Id:        todo.Id,
			Title:     (*db).data[todo.Id].Title,
			Completed: todo.Completed,
			CreatedAt: (*db).data[todo.Id].CreatedAt,
		}
		log.Printf("Status of element by id %d updated", todo.Id)
		return nil
	}

	return ErrNotFound
}

// Delete for delete by id todo
func (db *Database) Delete(id int) error {
	if _, ok := (*db).data[id]; ok {
		delete((*db).data, id)
		return nil
	}

	return ErrNotFound
}
