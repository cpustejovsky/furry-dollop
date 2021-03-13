package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", placeholder)
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", placeholderAPI)
	return r
}

func placeholder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func placeholderAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the API"))
}
