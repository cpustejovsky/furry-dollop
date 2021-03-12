package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	users map[int]string
}

func (s *StubUserStore) GetUser(id int) string {
	users := s.users[id]
	return users
}

func TestGETUsers(t *testing.T) {
	server := &UserServer{
		&StubUserStore{
			map[int]string{
				1: "Ryan Dahl",
				2: "Rob Pike",
			},
		},
	}
	tests := []struct {
		name string
		id   int
		want string
	}{
		{
			"Get Ryan Dahl by Id",
			1,
			"Ryan Dahl",
		},
		{
			"Get Rob Pike by Id",
			2,
			"Rob Pike",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/users/%v", tt.id), nil)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			got := response.Body.String()

			AssertEqual(t, got, tt.want)
		})
	}

}
