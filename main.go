package main

import (
	"log"
	"net/http"
)

type InMemoryUserStore struct{}

func (i *InMemoryUserStore) GetUser(id int) string {
	if id == 1 {
		return "Ryan Dahl"
	} else if id == 2 {
		return "Rob Pike"
	} else {
		return "No User Found"
	}
}

func main() {
	server := &UserServer{&InMemoryUserStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
