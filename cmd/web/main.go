package main

import (
	"todo-app/internal/app"
)

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

func main() {
	app.Run()
}
