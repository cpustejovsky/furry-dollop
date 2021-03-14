package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/gorilla/mux"
)

func (app *application) AuthenticateAndReturnID(w http.ResponseWriter, r *http.Request, email, password string) {
	id, err := app.users.Authenticate(email, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			app.errorLog.Println("email address or password was incorrect")
		} else {
			app.errorLog.Println(err)
		}
		return
	}
	app.session.Put(r, "authenticatedUserID", id)
	fmt.Fprint(w, app.session.GetString(r, "authenticatedUserID"))
}

func (app *application) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user struct {
		Email    string
		Password string
	}
	err := decoder.Decode(&user)
	if err != nil {
		w.Write([]byte("Something has gone wrong logging user in."))
		app.errorLog.Println(err)
	}
	app.AuthenticateAndReturnID(w, r, user.Email, user.Password)
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")
	app.session.Put(r, "flash", "You've been logged out successfully")
	app.infoLog.Println("logging user out")
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
		w.Write([]byte("Something has gone wrong with adding user."))
		app.errorLog.Println(err)
		return
	}
	w.Write([]byte("User added!"))
	app.AuthenticateAndReturnID(w, r, user.Email, user.Password)
}

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
	}
	app.errorLog.Println(err)
	w.Write(b)
}

//TODO: bring back update functionality in a way that won't overwrite SQL data
// func (app *application) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["userID"]
// 	decoder := json.NewDecoder(r.Body)
// 	var user struct {
// 		Name      string
// 		Email     string
// 		Expertise string
// 	}
// 	err := decoder.Decode(&user)
// 	if err != nil {
// 		w.Write([]byte("Something has gone wrong with updating user."))
// 		app.errorLog.Println(err)
// 	}
// 	u, err := app.users.Update(id, user.Name, user.Email, user.Expertise)
// 	if err != nil {
// 		w.Write([]byte("Something has gone wrong with updating user."))
// 		app.errorLog.Println(err)
// 	}
// 	b, err := json.Marshal(u)
// 	if err != nil {
// 		w.Write([]byte("Something has gone wrong with updating user."))
// 		app.errorLog.Println(err)
// 	}
// 	w.Write(b)
// }

func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uuid := app.session.GetString(r, "authenticatedUserID")
	err := app.users.Delete(uuid)
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
	uuid := app.session.GetString(r, "authenticatedUserID")
	decoder := json.NewDecoder(r.Body)
	var post FormPost
	err := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte("Something has gone wrong with adding post."))
		app.errorLog.Println(err)
		return
	}
	err = app.posts.Insert(uuid, post.Title, post.Body)
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
