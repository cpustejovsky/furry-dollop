package main

import (
	"net/http"
	"strconv"
	"strings"
)

type UserStore interface {
	GetUser(id int) string
}

type UserServer struct {
	store UserStore
}

func (u *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/api/users/")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.Write([]byte("Failure to Retrieve ID"))
	} else {
		w.Write([]byte(u.store.GetUser(id)))
	}
}
