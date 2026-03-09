package authdomain

type RegisterUser struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}
