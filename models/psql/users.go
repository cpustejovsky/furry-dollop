package psql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

func (m *UserModel) Insert(name, email, expertise string) error {
	stmt := `
	INSERT INTO users (username, email, expertise) 
	VALUES($1, $2, $3)`

	_, err := m.DB.Exec(stmt, name, email, expertise)
	if err != nil {
		var postgresError *pq.Error
		if errors.As(err, &postgresError) {
			if strings.EqualFold(string(postgresError.Code), "23505") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

func (m *UserModel) Update(id, name, email, expertise string) (*models.User, error) {
	u := &models.User{}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	stmt := `
	UPDATE users
	SET username = $2, email = $3, expertise = $4
	WHERE id = $1`
	_, err = m.DB.Exec(stmt, uuid, name, email, expertise)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	stmt = `
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
