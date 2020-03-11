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

	templ.ExecuteTemplate(writer, "base", data)
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

	templ.ExecuteTemplate(writer, "base", data)
}

func CreatePost(writer http.ResponseWriter, request *http.Request) {
	session := core.SessionGet(request, "user")

	templ := core.GetView("post/create")

	data := struct {
		UserID uint64
	}{session.Values["userID"].(uint64)}

	templ.ExecuteTemplate(writer, "base", data)
}

func StorePost(writer http.ResponseWriter, request *http.Request) {
	session := core.SessionGet(request, "user")

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
	// TODO: handle not found, 404 page, check user is author
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Post *models.Post
	}{post}

	templ.ExecuteTemplate(writer, "base", data)
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
