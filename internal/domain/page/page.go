package pagedomain

type WebPage string

const (
	Index    WebPage = "index"
	Register WebPage = "register"
	Login    WebPage = "login"
	Error    WebPage = "error"
)

type ErrorPage struct {
	Title      string
	Error      string
	StatusCode string
}
