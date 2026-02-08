package userdto

type GetUserResponse struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}
