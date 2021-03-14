package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) Routes() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", placeholderAPI)
	s.HandleFunc("/user/{userID}", app.GetUser).Methods("GET")
	s.HandleFunc("/user/new", app.AddUser).Methods("POST")
	return r
}

func placeholderAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the API"))
}
