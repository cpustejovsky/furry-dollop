package psql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Authenticate(email, password string) (string, error) {
	var id uuid.UUID
	var hashedPassword []byte
	stmt := `
	SELECT id, hashed_password 
	FROM users 
	WHERE email = $1`
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "error", models.ErrInvalidCredentials
		} else {
			return "error", err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "error", models.ErrInvalidCredentials
		} else {
			return "error", err
		}
	}

	return id.String(), nil
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

func (m *UserModel) Insert(name, email, expertise, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	stmt := `
	INSERT INTO users (username, email, expertise, hashed_password) 
	VALUES($1, $2, $3, $4)`

	_, err = m.DB.Exec(stmt, name, email, expertise, hashedPassword)
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

func (m *UserModel) Delete(id string) error {
	sqlStatement := `
	DELETE FROM users
	WHERE id = $1`
	_, err := m.DB.Exec(sqlStatement, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrNoRecord
		} else {
			return err
		}
	}
	return nil
}
