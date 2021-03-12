package main

import (
	"net/http"
	"strconv"
	"strings"
)

func UserServer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "api/users/"))
	if err != nil {
		w.Write([]byte("Failure to Retrieve ID"))
	}
	if id == 1 {
		w.Write([]byte("Ryan Dahl"))
	} else if id == 2 {
		w.Write([]byte("Rob Pike"))
	}
}
