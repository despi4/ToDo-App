package app

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"todo-app/internal/db"
	"todo-app/internal/repository/postgre"
	usersvc "todo-app/internal/service/user"
	userhandler "todo-app/internal/transport/http/handler/user"
	"todo-app/internal/transport/http/middleware"

	"github.com/joho/godotenv"
)

const (
	pattern = "ui/templates/*.html"
)

func Run() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	tmpl := template.Must(template.ParseGlob(pattern))

	db, err := db.NewDB(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 2458b756-daef-4978-861d-557e0671979d
	// ivan@example.com
	repo := postgre.NewUserRepo(db)
	service := usersvc.NewUserService(repo)
	handler := userhandler.NewUserHandler(service, tmpl)
	router := http.NewServeMux()

	router.HandleFunc("POST /users", handler.Create)
	router.HandleFunc("GET /users/{id}", handler.GetByID)
	router.HandleFunc("GET /users/", handler.GetByEmail)
	router.HandleFunc("PATCH /users/{id}", handler.Update)
	router.HandleFunc("DELETE /users/{id}", handler.Delete)

	mux := middleware.SecureHeaders(router)

	log.Printf("Server started on : %s\n", port)
	err = http.ListenAndServe(":"+port, middleware.Recover(middleware.Logger(mux)))
	if err != nil {
		log.Fatal(err)
	}
}
