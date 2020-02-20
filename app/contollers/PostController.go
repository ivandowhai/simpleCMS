package contollers

import (
	"../models"
	"fmt"
	"html/template"
	"net/http"
)

func PostsList(writer http.ResponseWriter, request *http.Request) {
	// TODO: move path to config
	templ, err := template.ParseFiles("templates/default/post/index.html")
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Posts []*models.Post
	}{Posts: models.GetAllPosts()}

	templ.Execute(writer, data)
}

func ViewPost(writer http.ResponseWriter, request *http.Request) {
	// TODO: move path to config
	templ, err := template.ParseFiles("templates/default/post/view.html")
	if err != nil {
		fmt.Println(err)
	}

	post, err := models.GetPostById(1)

	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Post *models.Post
	}{post}

	templ.Execute(writer, data)
}
