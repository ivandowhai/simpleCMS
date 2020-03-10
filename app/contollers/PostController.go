package contollers

import (
	"../core"
	"../models"
	"../repositories/post"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func PostsList(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("post/index")

	data := struct {
		Posts []*models.Post
	}{Posts: post.GetAll()}

	templ.Execute(writer, data)
}

func ViewPost(writer http.ResponseWriter, request *http.Request) {
	ID, err := strconv.ParseUint(mux.Vars(request)["postId"], 10, 16)
	if err != nil {
		fmt.Println(err)
	}

	templ := core.GetView("post/view")

	post, err := post.GetOne(ID)
	// TODO: handle not found
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Post *models.Post
	}{post}

	templ.Execute(writer, data)
}

func CreatePost(writer http.ResponseWriter, request *http.Request) {
	// TODO: to middleware
	session, err := core.Store.Get(request, "user")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.Values["userID"] == nil || session.Values["userRole"].(uint8) != 3 {
		http.Redirect(writer, request, "/profile", http.StatusSeeOther)
	}

	templ := core.GetView("post/create")

	data := struct {
		UserID uint64
	}{session.Values["userID"].(uint64)}

	templ.Execute(writer, data)
}

func StorePost(writer http.ResponseWriter, request *http.Request) {
	session, err := core.Store.Get(request, "user")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.Values["userID"] == nil || !core.CanUserPost(session.Values["userRole"].(uint8)) {
		http.Redirect(writer, request, "/profile", http.StatusSeeOther)
	}

	request.ParseForm()
	newPost := models.Post{Title: request.Form.Get("title"), Content: request.Form.Get("content"), AuthorID: session.Values["userID"].(uint64)}

	post.Create(newPost)

	http.Redirect(writer, request, "/profile", http.StatusSeeOther)
}

func EditPost(writer http.ResponseWriter, request *http.Request) {
	ID, err := strconv.ParseUint(mux.Vars(request)["postId"], 10, 16)
	if err != nil {
		fmt.Println(err)
	}

	templ := core.GetView("post/edit")

	post, err := post.GetOne(ID)
	// TODO: handle not found
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Post *models.Post
	}{post}

	templ.Execute(writer, data)
}

func UpdatePost(writer http.ResponseWriter, request *http.Request) {
	ID, err := strconv.ParseUint(mux.Vars(request)["postId"], 10, 16)
	if err != nil {
		fmt.Println(err)
	}

	request.ParseForm()
	post.Update(ID, request.Form.Get("title"), request.Form.Get("content"))

	http.Redirect(writer, request, "/posts/view/"+mux.Vars(request)["postId"], http.StatusSeeOther)
}

func DeletePost(writer http.ResponseWriter, request *http.Request) {
	ID, err := strconv.ParseUint(mux.Vars(request)["postId"], 10, 16)
	if err != nil {
		fmt.Println(err)
	}

	post.Delete(ID)

	http.Redirect(writer, request, "/profile", http.StatusSeeOther)
}
