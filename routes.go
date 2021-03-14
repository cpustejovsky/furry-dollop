package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) Routes() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", helloWorld)
	s.HandleFunc("/user/{userID}", app.GetUser).Methods("GET")
	s.HandleFunc("/user/new", app.AddUser).Methods("POST")
	s.HandleFunc("/user/{userID}", app.UpdateUser).Methods("PATCH")
	return r
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the API"))
}
