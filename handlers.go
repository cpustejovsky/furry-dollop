package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/gorilla/mux"
)

func (app *application) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userID"]
	u, err := app.users.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Write([]byte(u.Name))
	}
}

type FormUser struct {
	Name      string
	Email     string
	Expertise string
}

func (app *application) AddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user FormUser
	err := decoder.Decode(&user)
	if err != nil {
		w.Write([]byte("We're sorry, something has gone wrong with adding user."))
		app.errorLog.Println(err)
	}
	err = app.users.Insert(user.Name, user.Email, user.Expertise)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			w.Write([]byte("email address is already in use"))
			app.errorLog.Println(err)

		} else {
			w.Write([]byte("We're sorry, something has gone wrong with adding user."))
			app.errorLog.Println(err)
		}
		return
	}
	w.Write([]byte("User added!"))
}

func (app *application) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userID"]
	decoder := json.NewDecoder(r.Body)
	var user FormUser
	err := decoder.Decode(&user)
	if err != nil {
		w.Write([]byte("We're sorry, something has gone wrong with updating user."))
		app.errorLog.Println(err)
	}
	u, err := app.users.Update(id, user.Name, user.Email, user.Expertise)
	if err != nil {
		w.Write([]byte("We're sorry, something has gone wrong with updating user."))
		app.errorLog.Println(err)
	}
	b, err := json.Marshal(u)
	if err != nil {
		w.Write([]byte("We're sorry, something has gone wrong with updating user."))
		app.errorLog.Println(err)
	}
	w.Write(b)
}
