package main

import (
	"log"
	"net/http"
	handlers "todo-app/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)

	log.Println("Server started on:8081...")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
