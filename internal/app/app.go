package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"todo-app/internal/repository/postgre"
	usersvc "todo-app/internal/service/user"
	userhandler "todo-app/internal/transport/http/handler/user"

	"github.com/joho/godotenv"
)

func Init() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")

	db, err := postgre.NewDB(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	repo := postgre.NewUserRepo(db)
	service := usersvc.NewUserService(repo)
	handler := userhandler.NewUserHandler(service)
	router := http.NewServeMux()

	router.HandleFunc("POST /users", handler.Create)

	log.Println("Server started on : 3000")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
