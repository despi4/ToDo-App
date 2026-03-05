package userdomain

type UpdateUser struct {
	Name    *string
	Surname *string
	Email   *string
}

type RegisterUser struct {
	Name     string
	Surname  string
	Email    string
	Password string
}
