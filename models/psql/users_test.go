package psql_test

import (
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/cpustejovsky/furry-dollop/models/psql"
	"github.com/cpustejovsky/furry-dollop/testhelper"
)

var testUser = models.User{
	ID:        testhelper.TestUUID(),
	Name:      "Harry Haskell",
	Email:     "harry@example.com",
	Expertise: "Haskell",
}

func TestUserModelGet(t *testing.T) {
	tests := []struct {
		name      string
		userID    string
		wantUser  *models.User
		wantError error
	}{
		{
			name:      "Valid ID",
			userID:    testhelper.TestUUIDString,
			wantUser:  &testUser,
			wantError: nil,
		},
		{
			name:      "Invalid UUID",
			userID:    "hello1999",
			wantUser:  nil,
			wantError: models.ErrInvalidID,
		},
		{
			name:      "Non-existent ID",
			userID:    "38cb90c7-c7a7-40ea-a79e-e7f1aebadf82",
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := testhelper.NewMockDB(t)
			rows := mock.NewRows([]string{"id", "username", "email", "expertise"}).AddRow(testUser.ID, testUser.Name, testUser.Email, testUser.Expertise)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, email, expertise FROM users WHERE id = $1")).WithArgs(testhelper.TestUUID()).WillReturnRows(rows)

			m := psql.UserModel{db}

			user, err := m.Get(tt.userID)

			if err != tt.wantError {
				if !strings.Contains(err.Error(), "arguments do not match") {
					t.Errorf("want %v; got %s", tt.wantError, err)
				}
			}

			if !reflect.DeepEqual(user, tt.wantUser) {
				t.Errorf("want %v; got %v", tt.wantUser, user)
			}
		})
	}
}

func TestUserModelInsert(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (username, email, expertise)")).
		WithArgs(testUser.Name, testUser.Email, testUser.Expertise).
		WillReturnResult(sqlmock.NewResult(1, 1))

	m := psql.UserModel{db}

	err := m.Insert(testUser.Name, testUser.Email, testUser.Expertise)
	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}

}

func TestUserModelUpdate(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	rows := mock.NewRows([]string{"id", "username", "email", "expertise"}).AddRow(testUser.ID, testUser.Name, testUser.Email, testUser.Expertise)
	mock.ExpectExec(regexp.QuoteMeta("UPDATE users")).
		WithArgs(testUser.ID, testUser.Name, testUser.Email, testUser.Expertise).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, email, expertise FROM users WHERE id = $1")).
		WithArgs(testhelper.TestUUID()).
		WillReturnRows(rows)
	m := psql.UserModel{db}

	user, err := m.Update(testhelper.TestUUIDString, testUser.Name, testUser.Email, testUser.Expertise)
	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}

	if !reflect.DeepEqual(user, &testUser) {
		t.Errorf("want %v; got %v", testUser, user)
	}
}

func TestUserModelDelete(t *testing.T) {
	db, mock := testhelper.NewMockDB(t)
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = $1")).
		WithArgs(testhelper.TestUUID()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	m := psql.UserModel{db}

	err := m.Delete(testhelper.TestUUIDString)
	if err != nil {
		t.Errorf("expected nil, instead got following error:\n%v", err)
	}
}
