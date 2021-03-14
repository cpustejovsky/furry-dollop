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

func TestPostModelGetByID(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	rows := mock.NewRows([]string{"post_id", "title", "body", "id"}).AddRow(testPost.ID, testPost.Title, testPost.Body, testPost.UserId)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT post_id, id, title, body FROM posts WHERE post_id = $1")).WithArgs(testhelper.TestPostUUID()).WillReturnRows(rows)

	m := psql.PostModel{db}

	user, err := m.GetById(testhelper.TestPostUUIDString)

	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}

	if !reflect.DeepEqual(user, &testPost) {
		t.Errorf("want:\n%v\ngot:\n%v\n", &testPost, user)
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
