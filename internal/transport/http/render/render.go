package render

import (
	"html/template"
	"net/http"
	pagedomain "todo-app/internal/domain/page"
)

func Render(w http.ResponseWriter, name pagedomain.WebPage, tmpl *template.Template, data pagedomain.PageInfo) {
	err := "not found"

	switch name {
	case pagedomain.Register:
		data = pagedomain.PageInfo{
			Title: "Register",
		}
	case pagedomain.Index:
		data = pagedomain.PageInfo{
			Title: "Todo-App",
		}
	case pagedomain.Login:
		data = pagedomain.PageInfo{
			Title: "Login",
		}
	default:
		data = pagedomain.PageInfo{
			Title:        "Error",
			ErrorMessage: &err,
		}
	}

	tmpl.ExecuteTemplate(w, string(name), data)
}
