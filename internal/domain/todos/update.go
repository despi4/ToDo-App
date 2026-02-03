package todos

type TodoUpdate struct {
	Status      TodoStatus
	Title       string
	Description *string
}
