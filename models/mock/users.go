package mock

import (
	"github.com/cpustejovsky/furry-dollop/models"
)

var mockUser = &models.User{
	ID:        1,
	Name:      "Harry",
	Email:     "harry@example.com",
	Expertise: "Haskell",
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, expertise string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
