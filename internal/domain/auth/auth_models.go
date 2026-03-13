package authdomain

import (
	"errors"
)

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

type TokenType int

const (
	AccessTkn TokenType = iota
	RefreshTkn
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)
