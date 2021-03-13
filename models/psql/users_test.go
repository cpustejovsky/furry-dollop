package psql_test

import (
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/cpustejovsky/furry-dollop/models/psql"
	"github.com/cpustejovsky/furry-dollop/testhelper"
)

func TestUserModelGet(t *testing.T) {
	var testUser = models.User{
		ID:        testhelper.TestUUID(),
		Name:      "Harry Haskell",
		Email:     "harry@example.com",
		Expertise: "Haskell",
	}

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
			mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, email, expertise FROM users WHERE account_id = $1")).WithArgs(testhelper.TestUUID()).WillReturnRows(rows)

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
