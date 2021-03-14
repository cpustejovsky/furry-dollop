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
	// p := s.PathPrefix("/posts").Subrouter()
	// p.HandleFunc("/", app.GetPosts).Methods("GET")
	// p.HandleFunc("/{postID}", app.GetPostById).Methods("GET")
	// p.HandleFunc("/{postID}/{userID", app.GetPostsByUserId).Methods("GET")
	// p.HandleFunc("/new", app.AddPost).Methods("POST")
	// p.HandleFunc("/{postID}", app.UpdatePost).Methods("PATCH")
	// p.HandleFunc("/{postID}", app.DeletePost).Methods("DELETE")
	return r
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the API"))
}
