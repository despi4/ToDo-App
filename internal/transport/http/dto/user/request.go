package userdto

import userdomain "todo-app/internal/domain/user"

// Unmarshaling(Десериализация) - JSON string в Go struct, для принятие данных
// Marshaling(Сериализаций) - Go Struct в string JSON, для отправки данных

type CreateUserRequest struct {
	Name    string           `json:"name"`
	Surname string           `json:"surname"`
	Email   string           `json:"email"`
	Role    *userdomain.Role `json:"role"`
}

type UpdateUserRequest struct {
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Email   *string `json:"email"`
}
