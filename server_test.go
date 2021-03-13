package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	testhelper "github.com/cpustejovsky/furry-dollop/testhelper"
)

type StubStore struct {
	posts map[string]string
}

func (s *StubStore) GetPostsByUserID(id string) (string, error) {
	user, ok := s.posts[id]
	if !ok {
		return "", errors.New("User Not Found")
	} else {
		return user, nil
	}
}

func TestGETPostsByUserID(t *testing.T) {
	server := &Server{
		&StubStore{
			map[string]string{
				"1": "Ryan Dahl's posts",
				"2": "Rob Pike's posts",
			},
		},
	}
	tests := []struct {
		name string
		id   string
		want string
	}{
		{
			"Get Ryan Dahl's posts by Id",
			"1",
			"Ryan Dahl's posts",
		},
		{
			"Get Rob Pike by Id",
			"2",
			"Rob Pike's posts",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newGetPostsRequest(tt.id)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			got := response.Body.String()
			gotStatus := response.Code
			testhelper.AssertStatus(t, gotStatus, http.StatusOK)
			testhelper.AssertEqual(t, got, tt.want)
		})
	}
	t.Run("404 Error When No User is Found", func(t *testing.T) {
		req := newGetPostsRequest("555")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Code

		testhelper.AssertStatus(t, got, http.StatusNotFound)
	})
}

// func TestPOSTposts(t *testing.T) {
// 	store := StubStore{
// 		map[string]string{},
// 	}
// 	server := &Server{&store}

// 	t.Run("it returns accepted on POST", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/user/%s", id), nil)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		testhelper.AssertStatus(t, response.Code, http.StatusAccepted)
// 	})
// }

func newGetPostsRequest(id string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/posts/user/%s", id), nil)
	return request
}
