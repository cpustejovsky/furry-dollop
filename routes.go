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
	u.HandleFunc("/{userID}", app.UpdateUser).Methods("PUT")
	u.HandleFunc("/{userID}", app.DeleteUser).Methods("DELETE")

	p := s.PathPrefix("/posts").Subrouter()
	p.HandleFunc("/new", app.AddPost).Methods("POST")
	p.HandleFunc("/", app.GetPosts).Methods("GET")
	p.HandleFunc("/{postID}", app.GetPostsById).Methods("GET")
	p.HandleFunc("/user/{userID}", app.GetPostsByUserId).Methods("GET")
	p.HandleFunc("/{postID}", app.UpdatePost).Methods("PUT")
	p.HandleFunc("/{postID}", app.DeletePost).Methods("DELETE")
	return r
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the API"))
}
