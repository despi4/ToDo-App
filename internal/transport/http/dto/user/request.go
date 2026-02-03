package userdto

type CreateUserRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// Unmarshaling(Десериализация) - JSON string в Go struct, для принятие данных
// Marshaling(Сериализаций) - Go Struct в string JSON, для отправки данных

type UpdateUserRequest struct {
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Email   *string `json:"email"`
}
