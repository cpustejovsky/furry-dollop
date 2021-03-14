package testhelper

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

var TestPasswordString = "testpw"

func TestPassword() []byte {
	hp, err := bcrypt.GenerateFromPassword([]byte(TestPasswordString), 12)
	if err != nil {
		fmt.Println("Something went wrong")
	}
	return hp
}
