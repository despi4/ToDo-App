package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type message struct {
	Msg string
}

func main() {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		return
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}

		

		err = tmpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			return
		}
	})

	fmt.Println("Server starting : 5000")
	http.ListenAndServe("localhost:5000", mux)
}
