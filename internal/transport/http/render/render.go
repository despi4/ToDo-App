package render

import (
	"html/template"
	"net/http"
	"strconv"
	pagedomain "todo-app/internal/domain/page"
)

func Render(w http.ResponseWriter, name pagedomain.WebPage, data *pagedomain.ErrorPage, tmpl *template.Template) {
	if data != nil {
		tmpl.Execute(w, data)
	}

	err := tmpl.ExecuteTemplate(w, string(name), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		code := strconv.Itoa(http.StatusInternalServerError)

		data = &pagedomain.ErrorPage{
			Title:      "Error " + code,
			Error:      "Internal Server Error",
			StatusCode: code,
		}

		tmpl.ExecuteTemplate(w, "error", data)
	}
	w.WriteHeader(http.StatusOK)
}
