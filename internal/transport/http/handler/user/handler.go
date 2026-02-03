package userhandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	users "todo-app/internal/domain/user"
	usersvc "todo-app/internal/service/user"
	userdto "todo-app/internal/transport/http/dto/user"
)

type UserHandler struct {
	service usersvc.UserService
}

func NewUserHandler(service usersvc.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (userHandler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var creatReq userdto.CreateUserRequest

	// 1) Unmarshaling
	if err := json.NewDecoder(r.Body).Decode(&creatReq); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	var (
		name    = strings.TrimSpace(creatReq.Name)
		surname = strings.TrimSpace(creatReq.Surname)
		email   = strings.TrimSpace(creatReq.Email)
	)

	// 2) Validation
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if surname == "" {
		http.Error(w, "surname is required", http.StatusBadRequest)
		return
	}

	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	user := users.User{
		Name:    name,
		Surname: surname,
		Email:   email,
	}

	out, err := userHandler.service.Create(ctx, user)
	if err != nil {
		http.Error(w, "service error", http.StatusInternalServerError)
		return
	}

	log.Println(out)
}
