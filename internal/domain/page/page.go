package pagedomain

type WebPage string

const (
	Index    WebPage = "index"
	Register WebPage = "register"
	Login    WebPage = "login"
	Error    WebPage = "error"
)

type PageInfo struct {
	Title        string
	Message      *string
	ErrorMessage *string
	StatusCode   *int
}
