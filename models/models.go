package models

import "errors"

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type User struct {
	ID        int
	Name      string
	Email     string
	Expertise string
}


