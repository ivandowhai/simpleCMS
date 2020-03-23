package contollers

import (
	"../core"
	"../models"
	"../repositories"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func PostsList(writer http.ResponseWriter, _ *http.Request) {
	postRepository := repositories.PostRepository{}

	posts := postRepository.GetAll()

	jsonResponse, err := json.Marshal(posts)
	if err != nil {
		fmt.Println(err)
	}
	core.MakeResponse(writer, jsonResponse)
}

func ViewPost(writer http.ResponseWriter, request *http.Request) {
	logger := core.Logger{}
	logger.Init()
	postRepository := repositories.PostRepository{}
	ID, err := strconv.ParseUint(mux.Vars(request)["postId"], 10, 16)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	post, err := postRepository.GetOne(ID)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	jsonResponse, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	}
	core.MakeResponse(writer, jsonResponse)
}

func CreatePost(writer http.ResponseWriter, request *http.Request) {
	session := core.SessionGet(request, "user")

	templ := core.GetView("post/create", "main")

	data := struct {
		IsLogged bool
	}{session.Values["userID"] != nil}

	templ.ExecuteTemplate(writer, "base", data)
}

func StorePost(writer http.ResponseWriter, request *http.Request) {
	session := core.SessionGet(request, "user")
	postRepository := repositories.PostRepository{}

	request.ParseForm()
	newPost := models.Post{Title: request.Form.Get("title"), Content: request.Form.Get("content"), AuthorID: session.Values["userID"].(uint64)}

	postRepository.Create(newPost)

	http.Redirect(writer, request, "/profile", http.StatusSeeOther)
}

func EditPost(writer http.ResponseWriter, request *http.Request) {
	logger := core.Logger{}
	logger.Init()
	session := core.SessionGet(request, "user")
	postRepository := repositories.PostRepository{}
	ID, err := strconv.ParseUint(mux.Vars(request)["postId"], 10, 16)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	templ := core.GetView("post/edit", "main")

	post, err := postRepository.GetOne(ID)
	// TODO: handle not found, 404 page, check user is author
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	data := struct {
		Post     *models.Post
		IsLogged bool
	}{Post: post, IsLogged: session.Values["userID"] != nil}

	templ.ExecuteTemplate(writer, "base", data)
}

func UpdatePost(writer http.ResponseWriter, request *http.Request) {
	logger := core.Logger{}
	logger.Init()
	postRepository := repositories.PostRepository{}
	ID, err := strconv.ParseUint(mux.Vars(request)["postId"], 10, 16)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	request.ParseForm()
	postRepository.Update(ID, request.Form.Get("title"), request.Form.Get("content"))

	http.Redirect(writer, request, "/post/view/"+mux.Vars(request)["postId"], http.StatusSeeOther)
}

func DeletePost(writer http.ResponseWriter, request *http.Request) {
	logger := core.Logger{}
	logger.Init()
	postRepository := repositories.PostRepository{}
	ID, err := strconv.ParseUint(mux.Vars(request)["postId"], 10, 16)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	postRepository.Delete(ID)

	http.Redirect(writer, request, "/profile", http.StatusSeeOther)
}
