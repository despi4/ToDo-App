package render

import (
	"html/template"
	"net/http"
	pagedomain "todo-app/internal/domain/page"
)

func Render(w http.ResponseWriter, name pagedomain.WebPage, tmpl *template.Template) {
	err := tmpl.Execute(w, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
