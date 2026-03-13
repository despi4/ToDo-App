package app

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/repository/postgre"
	authsvc "todo-app/internal/service/auth"
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
	access_secret, refresh_secret := os.Getenv("JWT_ACCESS_SECRET"), os.Getenv("JWT_REFRESH_SECRET")

	tmpl := template.Must(template.ParseGlob(pattern))

	db, err := db.NewDB(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 2458b756-daef-4978-861d-557e0671979d
	// ivan@example.com
	repo := postgre.NewUserRepo(db)
	jwtService := authsvc.NewJWTService(
		[]byte(access_secret),
		[]byte(refresh_secret),
		time.Minute*15,
		time.Hour*167,
	)

	_ = authsvc.NewAuthService(repo, *jwtService)

	service := usersvc.NewUserService(repo)
	handler := userhandler.NewUserHandler(service, tmpl)
	router := http.NewServeMux()

	router.HandleFunc("GET /auth/register", handler.Register)
	router.HandleFunc("POST /auth/register", handler.Register)
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

func InitDB() {
	cmd := exec.Command("docker", "compose", "up", "-d")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal("Error docker compose up")
	}

	log.Println("Success")
}
