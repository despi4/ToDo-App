package userhandler

import (
	"context"
	"encoding/json"
	"errors"
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
		w.Header().Set("content-type", "application-json")

		errRes := userdto.ErrorResponse{
			Error:     userdomain.ErrInvalidArgument.Error(),
			ErrorCode: http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("User Not Created: %s", err)

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
		w.Header().Set("content-type", "application-json")

		errRes := userdto.ErrorResponse{
			Error:    "Internal Server Error",
			ErrorCode: http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("User not created: %s", err)

		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("New User Created by email=%s\n", out.Email)
}

func (userHandler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	url := strings.Split(r.URL.Path, "/")
	id, err := uuid.Parse(url[len(url)-1])
	if err != nil {
		w.Header().Set("content-type", "application-json")
		errRes := userdto.ErrorResponse{
			Error:     userdomain.ErrInvalidArgument.Error(),
			ErrorCode: http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("Can't Get User by id: %s", err)

		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	user, err := userHandler.service.GetByID(ctx, id)
	if err != nil {
		w.Header().Set("content-type", "application-json")

		var (
			err_msg string
			status  int
		)

		if errors.Is(err, userdomain.ErrNotFound) {
			err_msg = userdomain.ErrNotFound.Error()
			status = http.StatusNotFound
		} else {
			err_msg = "Internal Server Error"
			status = http.StatusNotFound
		}

		errRes := userdto.ErrorResponse{
			Error:     err_msg,
			ErrorCode: status,
		}

		w.WriteHeader(status)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("Can't Get User by id: %s", err)

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
	json.NewEncoder(w).Encode(&getRes)

	log.Println("Get User by id")
}

func (userHandler *UserHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application-json")
	queryParams := r.URL.Query()
	defer r.Body.Close()

	queryEmail := queryParams.Get("email")

	email := strings.TrimSpace(queryEmail)
	if email == "" {
		w.Header().Set("content-type", "application-json")
		errRes := userdto.ErrorResponse{
			Error:     userdomain.ErrInvalidArgument.Error(),
			ErrorCode: http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("Can't Get User by email: %s", userdomain.ErrInvalidArgument)

		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	user, err := userHandler.service.GetByEmail(ctx, email)
	if err != nil {
		w.Header().Set("content-type", "application-json")

		var (
			err_msg string
			status  int
		)

		if errors.Is(err, userdomain.ErrNotFound) {
			err_msg = userdomain.ErrNotFound.Error()
			status = http.StatusNotFound
		} else {
			err_msg = "Internal Server Error"
			status = http.StatusInternalServerError
		}

		errRes := userdto.ErrorResponse{
			Error:     err_msg,
			ErrorCode: status,
		}

		w.WriteHeader(status)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("Can't Get User by email(%s): %s",email, err)

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
	json.NewEncoder(w).Encode(&getRes)

	log.Printf("Get User by email=%s\n", user.Email)
}

func (userHandler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application-json")
	url := strings.Split(r.URL.Path, "/")
	defer r.Body.Close()

	id, err := uuid.Parse(url[len(url)-1])
	if err != nil {
		w.Header().Set("content-type", "application-json")
		errRes := userdto.ErrorResponse{
			Error:     userdomain.ErrInvalidArgument.Error(),
			ErrorCode: http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("User not Updated: %s", err)

		return
	}

	var updateReq userdto.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		w.Header().Set("content-type", "application-json")
		errRes := userdto.ErrorResponse{
			Error:     userdomain.ErrInvalidArgument.Error(),
			ErrorCode: http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("User not Updated: %s", err)

		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	user, err := userHandler.service.Update(ctx, id, userdomain.UserUpdate(updateReq))
	if err != nil {
		w.Header().Set("content-type", "application-json")

		var (
			err_msg string
			status  int
		)

		if errors.Is(err, userdomain.ErrNotFound) {
			err_msg = userdomain.ErrNotFound.Error()
			status = http.StatusNotFound
		} else {
			err_msg = "Internal Server Error"
			status = http.StatusInternalServerError
		}

		errRes := userdto.ErrorResponse{
			Error:     err_msg,
			ErrorCode: status,
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("User not Updated: %s", err)
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
	json.NewEncoder(w).Encode(&getRes)

	log.Printf("User Updated by email=%s\n", user.Email)
}

func (userHandler *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	id, err := uuid.Parse(url[len(url)-1])
	defer r.Body.Close()

	if err != nil {
		w.Header().Set("content-type", "application-json")
		errRes := userdto.ErrorResponse{
			Error:     userdomain.ErrInvalidArgument.Error(),
			ErrorCode: http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("User not Deleted: %s", err)

		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	err = userHandler.service.Delete(ctx, id)
	if err != nil {
		w.Header().Set("content-type", "application-json")
		var (
			err_msg string
			status  int
		)

		if errors.Is(err, userdomain.ErrNotFound) {
			err_msg = userdomain.ErrNotFound.Error()
			status = http.StatusNotFound
		} else {
			err_msg = "Internal Server Error"
			status = http.StatusInternalServerError
		}

		errRes := userdto.ErrorResponse{
			Error:     err_msg,
			ErrorCode: status,
		}

		w.WriteHeader(status)

		json.NewEncoder(w).Encode(&errRes)

		log.Printf("User not deleted: %s", err)

		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("User Deleted by id=%s", id)
}
