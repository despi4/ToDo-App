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

/*
Если не вызывать метод w.WriteHeader()
напрямую, тогда первый вызов w.Write()
автоматически отправит пользователю код состояния 200 OK.
Поэтому, если вы хотите вернуть другой код состояния,
вызовите один раз метод w.WriteHeader() перед любым вызовом w.Write()
*/

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
	log.Printf("New User Created by email=%s\n", out.Email)
}

func (userHandler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
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

	var getRes userdto.GetUserResponse

	getRes = userdto.GetUserResponse{
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		CreatedAt: time.Time(user.CreatedAt).Format(time.DateTime),
		UpdatedAt: time.Time(user.UpdatedAt).Format(time.DateTime),
	}

	// w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&getRes); err != nil {
		http.Error(w, "invalid something", http.StatusBadRequest)
	}
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

	var getRes userdto.GetUserResponse

	getRes = userdto.GetUserResponse{
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		CreatedAt: time.Time(user.CreatedAt).Format(time.DateTime),
		UpdatedAt: time.Time(user.UpdatedAt).Format(time.DateTime),
	}

	// w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&getRes); err != nil {
		http.Error(w, "invalid something", http.StatusBadRequest)
	}

	log.Printf("Get User by email=%s\n", user.Email)
}

func (userHandler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[:7]
	defer r.Body.Close()

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

	var getRes userdto.GetUserResponse

	getRes = userdto.GetUserResponse{
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		CreatedAt: time.Time(user.CreatedAt).Format(time.DateTime),
		UpdatedAt: time.Time(user.UpdatedAt).Format(time.DateTime),
	}

	// w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&getRes); err != nil {
		http.Error(w, "invalid something", http.StatusBadRequest)
	}

	log.Printf("User Updated by email=%s\n", user.Email)
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
	log.Printf("User Deleted by email=%s", id)
}
