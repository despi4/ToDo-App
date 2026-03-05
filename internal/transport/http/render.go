package http

import (
	"html/template"
	"net/http"
)

type Page string

const (
	indexPage = "index"
	loginPage = "login"
	register  = "register"
	error     = "error"
)

const (
	pattern = "ui/templates/*.html"
)

func render(w http.ResponseWriter) {
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}
