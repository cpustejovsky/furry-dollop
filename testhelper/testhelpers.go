package testhelper

import (
	"database/sql"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var TestUserUUIDString string = "af60a115-c4d1-4262-8ecb-3a36f65aa6b3"

func TestUserUUID() uuid.UUID {
	uuid, _ := uuid.Parse("af60a115-c4d1-4262-8ecb-3a36f65aa6b3")
	return uuid
}

var TestPostUUIDString string = "f630fe58-ca59-4179-a3e1-9aaa08ea0f97"

func TestPostUUID() uuid.UUID {
	uuid, _ := uuid.Parse("f630fe58-ca59-4179-a3e1-9aaa08ea0f97")
	return uuid
}

//NewMockDB creates a new sqlmock db for testing
func NewMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

//AssertEqual takes two interfaces and determines if they are equal
func AssertEqual(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%v\nwant\n%v\n", got, want)
	}
}

//AssertStatus takes two status codes and compares them
func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

// ErrorContains checks if the error message in out contains the text in
// want.
//
// This is safe when out is nil. Use an empty string for want if you want to
// test that err is nil.
//pulled from: https://github.com/Teamwork/test/blob/7949982/test.go#L12-L25
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
