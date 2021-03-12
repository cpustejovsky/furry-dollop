package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETUsers(t *testing.T) {
	t.Run("returns Ryan Dahl's sp", func(t *testing.T) {
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
				request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("api/users/%v", tt.id), nil)
				response := httptest.NewRecorder()

				UserServer(response, request)

				got := response.Body.String()

				if got != tt.want {
					t.Errorf("got %q, want %q", got, tt.want)
				}
			})
		}
	})
}
