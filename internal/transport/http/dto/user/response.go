package userdto

import userdomain "todo-app/internal/domain/user"

type GetUserResponse struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Email     string          `json:"email"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Role      userdomain.Role `json:"role"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}
