package handlers

import (
	"net/http"
	"todo-app/internal/service"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

// POST
func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriterm, r *http.Request) error {
	if r.Body.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 400)
		return
	}
}

// GET
// func (h *TodoHandler) GetTodoHandler() error {

// }

// PUT
// func (h *TodoHandler) MarkIsDoneHandler() error {

// }

// DELETE
// func (h *TodoHandler) DeleteTodoHandler() error {

// }
