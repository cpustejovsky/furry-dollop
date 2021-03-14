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

func (m *PostModel) Get() (*[]models.Post, error) {
	var posts []models.Post
	stmt := `
	SELECT *
	FROM posts
	`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := &models.Post{}
		err = rows.Scan(&p.ID, &p.UserId, &p.Title, &p.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *p)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &posts, nil
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
