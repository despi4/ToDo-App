package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"todo-app/internal/config"
	"todo-app/internal/handlers"
	"todo-app/internal/repository"
	"todo-app/internal/service"
)

func main() {
	err := LoadDotEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")

	dbl := repository.NewDatabase()
	service := service.NewTodoService(dbl)
	handler := handlers.NewTodoHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", handler.CreateTodoHandler)
	mux.HandleFunc("/todos/", handler.GetTodoHandler)
	mux.HandleFunc("/todos/getbyid", handler.GetTodoByIdHandler)
	mux.HandleFunc("/todos/update", handler.MarkIsDoneHandler)
	mux.HandleFunc("/todos/delete", handler.DeleteTodoHandler)

	log.Print("Server started on 8081...")
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadDotEnv() error {
	file, err := os.Open("./.env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		subLine := strings.SplitN(line, "=", 2)
		if len(subLine) != 2 {
			continue
		}

		var (
			key   = strings.TrimSpace(subLine[0])
			value = strings.Trim(strings.TrimSpace(subLine[1]), "\"")
		)

		os.Setenv(key, value)
	}

	return scanner.Err()
}

// URL - endpoints | http://librarian.com/books
// URI - уникальный адресс рессурса | http://librarian.com/books?author=Gogol (GET parametr)

// API - интерфейс для взаимодействия между программами

// Поток данных
// HTTP Request >> handlers/ >> service/ >> repository/ >> models/

// 1. Хэндлер получает JSON → создаёт models.Todo
// 2. Вызывает service.CreateTodo(title)

// 3. Сервис:
//    - проверяет: title не пустой, не длинный и т.д.
//    - создаёт todo := models.Todo{Title: title, Done: false}
//    - вызывает repo.Create(&todo)

// 4. Репозиторий:
//    - выполняет INSERT
//    - возвращает ошибку, если БД отказалась (например, дубль)
//    - НЕ проверяет title!
