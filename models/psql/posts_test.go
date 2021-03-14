package psql_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/cpustejovsky/furry-dollop/models/psql"

	"github.com/cpustejovsky/furry-dollop/testhelper"
)

var testPost = models.Post{
	UserId: testhelper.TestUserUUID(),
	ID:     testhelper.TestPostUUID(),
	Title:  "FP FTW",
	Body:   "Why be practical when you can be PURE!",
}

func TestPostModelInsert(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO posts (id, title, body)")).
		WithArgs(testPost.UserId, testPost.Title, testPost.Body).
		WillReturnResult(sqlmock.NewResult(1, 1))

	m := psql.PostModel{db}

	err := m.Insert(testhelper.TestUserUUIDString, testPost.Title, testPost.Body)
	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}
}

func TestPostModelGetAll(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	rows := mock.NewRows([]string{"post_id", "title", "body", "id"}).AddRow(testPost.ID, testPost.Title, testPost.Body, testPost.UserId)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM posts")).WillReturnRows(rows)

	m := psql.PostModel{db}

	posts, err := m.GetAll()

	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}

	if !reflect.DeepEqual(posts, &[]models.Post{testPost}) {
		t.Errorf("want:\n%v\ngot:\n%v\n", &testPost, posts)
	}

}

func TestPostModelGetByID(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	rows := mock.NewRows([]string{"post_id", "title", "body", "id"}).AddRow(testPost.ID, testPost.Title, testPost.Body, testPost.UserId)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT post_id, title, body, id FROM posts WHERE post_id = $1")).WithArgs(testhelper.TestPostUUID()).WillReturnRows(rows)

	m := psql.PostModel{db}

	post, err := m.GetById(testhelper.TestPostUUIDString)

	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}

	if !reflect.DeepEqual(post, &testPost) {
		t.Errorf("want:\n%v\ngot:\n%v\n", &testPost, post)
	}

}

func TestPostModelGetByUserID(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	rows := mock.NewRows([]string{"post_id", "title", "body", "id"}).AddRow(testPost.ID, testPost.Title, testPost.Body, testPost.UserId)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT post_id, title, body, id FROM posts WHERE id = $1")).WithArgs(testhelper.TestUserUUID()).WillReturnRows(rows)

	m := psql.PostModel{db}

	posts, err := m.GetByUserId(testhelper.TestUserUUIDString)

	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}

	if !reflect.DeepEqual(posts, &[]models.Post{testPost}) {
		t.Errorf("want:\n%v\ngot:\n%v\n", &[]models.Post{testPost}, posts)
	}

}

func TestPostModelUpdate(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	rows := mock.NewRows([]string{"post_id", "title", "body", "id"}).AddRow(testPost.ID, testPost.Title, testPost.Body, testPost.UserId)
	mock.ExpectExec(regexp.QuoteMeta("UPDATE posts")).
		WithArgs(testhelper.TestPostUUIDString, testPost.Title, testPost.Body).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT post_id, title, body, id FROM posts WHERE post_id = $1")).
		WithArgs(testhelper.TestPostUUID()).
		WillReturnRows(rows)
	m := psql.PostModel{db}

	post, err := m.Update(testhelper.TestPostUUIDString, testPost.Title, testPost.Body)
	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}

	if !reflect.DeepEqual(post, &testPost) {
		t.Errorf("want %v; got %v", testUser, post)
	}
}

func TestPostModelDelete(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM posts WHERE post_id = $1")).
		WithArgs(testhelper.TestPostUUID()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	m := psql.PostModel{db}

	err := m.Delete(testhelper.TestPostUUIDString)
	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}
}
