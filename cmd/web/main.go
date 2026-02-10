package main

import "todo-app/internal/app"

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

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	msg := "Hello"

// 	w.Write([]byte(msg))
// }

// func OtherHandler(w http.ResponseWriter, r *http.Request) {
// 	msg := "Goodbye"

// 	w.Write([]byte(msg))
// }

// func Logger(next http.Handler) http.Handler {
// 	function := func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("addr:%s method:%s uri:%s proto:%s", r.RemoteAddr, r.Method, r.RequestURI, r.Proto)

// 		next.ServeHTTP(w, r)
// 	}

// 	return http.HandlerFunc(function)
// }

// func Router() {
// 	http.HandleFunc("/salam", Handler)

// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/goodbye", OtherHandler)

// 	m := Logger(mux)

// 	mux.HandleFunc("/salam", Handler)

// 	http.ListenAndServe(":8080", m)
// }

// func main() {
// 	Router()
// }

func main() {
	app.Run()
}
