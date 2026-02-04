package userhandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	userdomain "todo-app/internal/domain/user"
	userdto "todo-app/internal/transport/http/dto/user"

	"github.com/google/uuid"
)

type UserHandler struct {
	service userdomain.UserService
}

func NewUserHandler(service userdomain.UserService) *UserHandler {
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
	defer r.Body.Close()

	var (
		name    = strings.TrimSpace(creatReq.Name)
		surname = strings.TrimSpace(creatReq.Surname)
		email   = strings.TrimSpace(creatReq.Email)
	)

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	user := userdomain.User{
		Name:    name,
		Surname: surname,
		Email:   email,
	}

	out, err := userHandler.service.Create(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	log.Println("New User created")
	log.Printf("ID: %s, Name: %s, Surname: %s, Email: %s, CreatedAt: %s, UpdteadAt: %s\n", out.ID, out.Name, out.Surname, out.Email, out.CreatedAt, out.UpdatedAt)
}

func (userHandler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[7:]

	id, err := uuid.Parse(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id == uuid.Nil {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	user, err := userHandler.service.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	log.Println("Get User By ID")
	log.Printf("ID: %s, Name: %s, Surname: %s, Email: %s, CreatedAt: %s, UpdteadAt: %s\n", user.ID, user.Name, user.Surname, user.Email, user.CreatedAt, user.UpdatedAt)
}

func (userHandler *UserHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	defer r.Body.Close()

	queryEmail := queryParams.Get("email")

	email := strings.TrimSpace(queryEmail)
	if email == "" {
		http.Error(w, "email is requied", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	user, err := userHandler.service.GetByEmail(ctx, email)
	if err != nil {
		http.Error(w, "service error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	log.Println("Get User By Email")
	log.Printf("ID: %s, Name: %s, Surname: %s, Email: %s, CreatedAt: %s, UpdteadAt: %s\n", user.ID, user.Name, user.Surname, user.Email, user.CreatedAt, user.UpdatedAt)
}

func (userHandler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[:7]

	id, err := uuid.Parse(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id == uuid.Nil {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var updateReq userdto.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	user, err := userHandler.service.Update(ctx, id, userdomain.UserUpdate(updateReq))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	log.Println("User updated")
	log.Println("Updated rows")

	if updateReq.Name != nil {
		log.Println("name: ", user.Name)
	}

	if updateReq.Surname != nil {
		log.Println("surname: ", user.Surname)
	}

	if updateReq.Email != nil {
		log.Println("email: ", user.Email)
	}
}

func (userHandler *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	defer r.Body.Close()

	queryID := queryParams.Get("id")

	id, err := uuid.Parse(queryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id != uuid.Nil {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	err = userHandler.service.Delete(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
