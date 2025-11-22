package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
