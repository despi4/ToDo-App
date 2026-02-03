package todos

type TodoFilter struct {
	Search string
	Limit  string
	Offset string
	Status TodoStatus
}
