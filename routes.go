package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/gorilla/mux"
)

func (app *application) Routes() *mux.Router {
	r := mux.NewRouter()
	r.Use(app.session.Enable)
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", helloWorld)
	s.HandleFunc("/login", app.AuthenticateUser).Methods("POST")
	s.HandleFunc("/logout", app.Logout).Methods("GET")

	u := s.PathPrefix("/users").Subrouter()
	u.HandleFunc("/{userID}", app.GetUser).Methods("GET")
	u.HandleFunc("/new", app.AddUser).Methods("POST")
	u.HandleFunc("/{userID}", app.UpdateUser).Methods("PUT")
	u.HandleFunc("/delete", app.DeleteUser).Methods("DELETE")
	u.Use(app.authenticate)

	p := s.PathPrefix("/posts").Subrouter()
	p.HandleFunc("/new", app.AddPost).Methods("POST")
	p.HandleFunc("/all", app.GetPosts).Methods("GET")
	p.HandleFunc("/{postID}", app.GetPostsById).Methods("GET")
	p.HandleFunc("/user/{userID}", app.GetPostsByUserId).Methods("GET")
	p.HandleFunc("/{postID}", app.UpdatePost).Methods("PUT")
	p.HandleFunc("/{postID}", app.DeletePost).Methods("DELETE")
	p.Use(app.authenticate)
	return r
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the API"))
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exists := app.session.Exists(r, "authenticatedUserID")
		if !exists {
			next.ServeHTTP(w, r)
			return
		}
		id := app.session.GetString(r, "authenticatedUserID")
		app.infoLog.Println(id)
		_, err := app.users.Get(id)
		if errors.Is(err, models.ErrNoRecord) {
			app.session.Remove(r, "authenticatedUserID")
			next.ServeHTTP(w, r)
			return
		} else if err != nil {
			app.errorLog.Println(err)
			return
		}
		infoLog.Println("user is authenticated")
		ctx := context.WithValue(r.Context(), contextKeyIsAuthenticated, true)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
