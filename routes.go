package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) Routes() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", helloWorld)
	u := s.PathPrefix("/users").Subrouter()
	u.HandleFunc("/{userID}", app.GetUser).Methods("GET")
	u.HandleFunc("/new", app.AddUser).Methods("POST")
	u.HandleFunc("/{userID}", app.UpdateUser).Methods("PATCH")
	u.HandleFunc("/{userID}", app.DeleteUser).Methods("DELETE")
	// n := s.PathPrefix("/notes").Subrouter()
	// n.HandleFunc("/", app.GetNotes).Methods("GET")
	// n.HandleFunc("/{noteID}", app.GetNoteById).Methods("GET")
	// n.HandleFunc("/{noteID}/{userID", app.GetNotesByUserId).Methods("GET")
	// n.HandleFunc("/new", app.AddNote).Methods("POST")
	// n.HandleFunc("/{noteID}", app.UpdateNote).Methods("PATCH")
	// n.HandleFunc("/{noteID}", app.DeleteNote).Methods("DELETE")
	return r
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the API"))
}
