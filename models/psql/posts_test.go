package psql_test

import (
	"reflect"
	"regexp"
	"testing"

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

func TestPostModelGet(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	rows := mock.NewRows([]string{"post_id", "id", "title", "body"}).AddRow(testPost.ID, testPost.UserId, testPost.Title, testPost.Body)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT post_id, id, title, body FROM posts WHERE post_id = $1")).WithArgs(testhelper.TestPostUUID()).WillReturnRows(rows)

	m := psql.PostModel{db}

	user, err := m.Get(testhelper.TestPostUUIDString)

	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}

	if !reflect.DeepEqual(user, &testPost) {
		t.Errorf("want:\n%v\ngot:\n%v\n", &testPost, user)
	}

}
