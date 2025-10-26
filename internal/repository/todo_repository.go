package repository

import (
	"errors"
	"todo-app/internal/models"
)

// Отделить логику хранения данных от бизнес-логики

type TodoRepository interface {
	Create(todo *models.Todo) error       // create todo
	GetAll() ([]models.Todo, error)       // Get todo
	GetById(id int) (*models.Todo, error) // search todo by id
	Update()                              // update todo
	Delete()                              // delete todo
}

type Database map[int]*models.Todo

func (db *Database) Create(todo *models.Todo) error {
	if todo == nil {
		return errors.New("todo must not be nil")
	}

	if len(*db) == 0 {
		(*db)[todo.Id] = todo
		return nil
	}

	(*db)[todo.Id] = todo

	return nil
}

func (db *Database) GetAll() ([]models.Todo, error) {
	var todoList []models.Todo

	for _, data := range *db {
		todoList = append(todoList, *data)
	}

	return todoList, nil
}

func (db *Database) GetById(id int) (*models.Todo, error) {
	if id < 1 {
		return nil, errors.New("id can not be non-positive")
	}

	if _, ok := (*db)[id]; ok {
		return (*db)[id], nil
	}

	return nil, errors.New("database does not have this id")
}
