package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Println("Server started on:8081...")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

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
