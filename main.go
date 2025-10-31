package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Попробуйте эту строку подключения
	connStr := "host=localhost port=8080 user=postgres password=postgres dbname=TodoAppWeb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	fmt.Println("Successfully connected!")
}
