package repository

import (
	"errors"
	"todo-app/internal/models"
)

// Отделить логику хранения данных от бизнес-логики

type TodoRepository interface {
	Create(todo *models.Todo) error // create todo
	GetAll()                        // Get todo
	GetById()                       // search todo by id
	Update()                        // update todo
	Delete()                        // delete todo
}

type Database map[int]*models.Todo

func (db *Database) Create(todo *models.Todo) error {
	if todo == nil {
		return errors.New("todo must not be nil")
	}

	(*db)[todo.Id] = todo
	
	return nil
}
