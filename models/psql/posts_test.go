package psql_test

import (
	"testing"

	"github.com/cpustejovsky/furry-dollop/models"

	"github.com/cpustejovsky/furry-dollop/testhelper"
)

var testNote = models.Post{
	UserId: testhelper.TestUserUUID(),
	ID:     testhelper.TestPostUUID(),
	Title:  "FP FTW",
	Body:   "Why be practical when you can be PURE!",
}

func TestNoteModelGet(t *testing.T) {

}
