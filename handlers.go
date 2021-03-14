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
		return
	}
	b, err := json.Marshal(u)
	if err != nil {
		w.Write([]byte("Something has gone wrong with updating user."))
		app.errorLog.Println(err)
	}
	w.Write(b)
}

func (app *application) AddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user struct {
		Name      string
		Email     string
		Expertise string
		Password  string
	}
	err := decoder.Decode(&user)
	if err != nil {
		w.Write([]byte("Something has gone wrong with adding user."))
		app.errorLog.Println(err)
	}
	err = app.users.Insert(user.Name, user.Email, user.Expertise, user.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			w.Write([]byte("email address is already in use"))
			app.errorLog.Println(err)

		} else {
			w.Write([]byte("Something has gone wrong with adding user."))
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
	var user struct {
		Name      string
		Email     string
		Expertise string
	}
	err := decoder.Decode(&user)
	if err != nil {
		w.Write([]byte("Something has gone wrong with updating user."))
		app.errorLog.Println(err)
	}
	u, err := app.users.Update(id, user.Name, user.Email, user.Expertise)
	if err != nil {
		w.Write([]byte("Something has gone wrong with updating user."))
		app.errorLog.Println(err)
	}
	b, err := json.Marshal(u)
	if err != nil {
		w.Write([]byte("Something has gone wrong with updating user."))
		app.errorLog.Println(err)
	}
	w.Write(b)
}

func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userID"]
	err := app.users.Delete(id)
	if err != nil {
		w.Write([]byte("Something has gone wrong with deleting user."))
		app.errorLog.Println(err)
		return
	}
	w.Write([]byte("User deleted"))
}

//NOTE ROUTES

type FormPost struct {
	UserID string
	Title  string
	Body   string
}

func (app *application) AddPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var post FormPost
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte("Something has gone wrong with adding post."))
		app.errorLog.Println(err)
		return
	}
	err = app.posts.Insert(post.UserID, post.Title, post.Body)
	if err != nil {
		w.Write([]byte("Something has gone wrong with adding post."))
		app.errorLog.Println(err)
		return
	}
	w.Write([]byte("Post added!"))
}

func (app *application) GetPosts(w http.ResponseWriter, r *http.Request) {
	n, err := app.posts.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		w.Write([]byte("Something has gone wrong with fetching posts."))
		return
	}
	b, err := json.Marshal(n)
	if err != nil {
		w.Write([]byte("Something has gone wrong with fetching posts."))
		app.errorLog.Println(err)
	}
	w.Write(b)
}

func (app *application) GetPostsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["postID"]
	n, err := app.posts.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Something has gone wrong with fetching post."))
		return
	}
	b, err := json.Marshal(n)
	if err != nil {
		w.Write([]byte("Something has gone wrong with fetching post."))
		app.errorLog.Println(err)
	}
	w.Write(b)
}

func (app *application) GetPostsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userID"]
	n, err := app.posts.GetByUserId(id)
	if err != nil {
		app.errorLog.Println(err)
		w.Write([]byte("Something has gone wrong with fetching posts by user."))
		return
	}
	b, err := json.Marshal(n)
	if err != nil {
		app.errorLog.Println(err)
		w.Write([]byte("Something has gone wrong with formatting posts by user."))
	}
	w.Write(b)
}

func (app *application) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["postID"]
	decoder := json.NewDecoder(r.Body)
	var post struct {
		Title string
		Body  string
	}
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte("Something has gone wrong with updating post."))
		app.errorLog.Println(err)
	}
	u, err := app.posts.Update(id, post.Title, post.Body)
	if err != nil {
		w.Write([]byte("Something has gone wrong with updating post."))
		app.errorLog.Println(err)
	}
	b, err := json.Marshal(u)
	if err != nil {
		w.Write([]byte("Something has gone wrong with showing updated post."))
		app.errorLog.Println(err)
	}
	w.Write(b)
}

func (app *application) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["postID"]
	err := app.posts.Delete(id)
	if err != nil {
		w.Write([]byte("Something has gone wrong with deleting post."))
		app.errorLog.Println(err)
		return
	}
	w.Write([]byte("Post deleted"))
}
