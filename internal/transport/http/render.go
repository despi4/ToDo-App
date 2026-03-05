package http

import (
	"html/template"
	"net/http"
	domainpage "todo-app/internal/domain/page"
)

func render(w http.ResponseWriter, name domainpage.Page, tmpl *template.Template) {

}
