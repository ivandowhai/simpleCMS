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
