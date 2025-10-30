package main

import (
	"log"
	"net/http"

	"todo-app/internal/handlers"
	"todo-app/internal/repository"
	"todo-app/internal/service"
)

func main() {
	db := repository.NewDatabase()
	service := service.NewTodoService(db)
	handler := handlers.NewTodoHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", handler.CreateTodoHandler)
	mux.HandleFunc("/todos/", handler.GetTodoHandler)
	mux.HandleFunc("/todos/getbyid", handler.GetTodoByIdHandler)
	mux.HandleFunc("/todos/update", handler.MarkIsDoneHandler)
	mux.HandleFunc("/todos/delete", handler.DeleteTodoHandler)

	log.Print("Server started on 8081...")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
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
