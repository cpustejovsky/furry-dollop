package models

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidID          = errors.New("models: invalid UUID")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Post struct {
	UserId uuid.UUID
	Id     uuid.UUID
	Title  string
	Body   string
}

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Expertise string
}
