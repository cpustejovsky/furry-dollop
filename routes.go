package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", placeholder)
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", placeholderAPI)
	s.HandleFunc("/user/{userID}", app.getUser).Methods("GET")
	return r
}

func placeholder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func placeholderAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the API"))
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userID"]
	fmt.Println(id)
	u, err := app.users.Get(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Write([]byte(u.Name))
	}
}
