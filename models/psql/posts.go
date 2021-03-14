package psql

import (
	"database/sql"
	"errors"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/google/uuid"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) GetById(id string) (*models.Post, error) {
	p := &models.Post{}
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, models.ErrInvalidID
	}
	stmt := `
	SELECT post_id, id, title, body 
	FROM posts 
	WHERE post_id = $1`
	err = m.DB.QueryRow(stmt, uuid).Scan(&p.ID, &p.UserId, &p.Title, &p.Body)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return p, nil
}

func (m *PostModel) Insert(userid, title, body string) error {
	stmt := `
	INSERT INTO posts (id, title, body) 
	VALUES($1, $2, $3)`

	_, err := m.DB.Exec(stmt, userid, title, body)
	if err != nil {
		return err
	}

	return nil
}
