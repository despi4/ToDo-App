package handlers

import (
	"encoding/json"
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
func (h *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type createRequest struct {
		Title string `json:"title"`
	}

	var req createRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = h.service.CreateTodo(req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo created succesfully"})
}

// GET
func (h *TodoHandler) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filter := r.URL.Query().Get("filter")
	todos, err := h.service.GetTodo(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// PUT
// func (h *TodoHandler) MarkIsDoneHandler() error {

// }

// DELETE
// func (h *TodoHandler) DeleteTodoHandler() error {

// }
