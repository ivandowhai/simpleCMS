package contollers

import (
	"../models"
	"../repositories/post"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

func PostsList(writer http.ResponseWriter, request *http.Request) {
	// TODO: move path to config
	templ, err := template.ParseFiles("templates/default/post/index.html")
	if err != nil {
		fmt.Println(err)
	}

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
	templ, err := template.ParseFiles("templates/default/post/view.html")
	if err != nil {
		fmt.Println(err)
	}

	post, err := models.GetPostById(ID)
	// TODO: handle not found
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Post *models.Post
	}{post}

	templ.Execute(writer, data)
}
