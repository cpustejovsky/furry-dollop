package psql

import (
	"database/sql"
	"errors"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/google/uuid"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Get(id string) (*models.User, error) {
	u := &models.User{}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, models.ErrInvalidID
	}
	stmt := `
	SELECT id, username, email, expertise 
	FROM users 
	WHERE id = $1`
	err = m.DB.QueryRow(stmt, uuid).Scan(&u.ID, &u.Name, &u.Email, &u.Expertise)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return u, nil
}
