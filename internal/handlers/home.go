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
		return
	}

	if len(todos) != 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// PUT
func (h *TodoHandler) MarkIsDoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type updateRequest struct {
		Id int `json:"id"`
	}

	var req updateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateStatusTodo(req.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo updated succesfully"})
}

// DELETE
func (h *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type deleteRequest struct {
		Id int `json:"id"`
	}

	var req deleteRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTodo(req.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo deleted succesfully"))
}

/*
Зачем использовать конструкторы

1)Инкапсуляция: вся логика создания объекта сосредоточена в одном месте.
2)Упрощение использования: снаружи не нужно знать, как правильно создать TodoHandler, — просто вызываешь NewTodoHandler(service).
3)Контроль зависимостей: можно быть уверенным, что объект всегда будет создан с нужными зависимостями (service в данном случае).
4)Готовность к тестированию: легко подменить зависимости на тестовые (например, мок-сервис).
*/

// json.NewDecoder(...) - создаёт декодер JSON, который умеет читать JSON-данные из потока
// json.NewDecoder(...).Decode(...) - декодирует (распаковывает) этот JSON в твою структуру Go.
