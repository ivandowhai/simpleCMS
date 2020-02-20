package contollers

import (
	"../models"
	"fmt"
	"html/template"
	"net/http"
)

func PostsList(writer http.ResponseWriter, request *http.Request) {
	templ, err := template.ParseFiles("templates/default/index.html")
	if err != nil {
		fmt.Println(err)
	}

	data := struct {
		Posts []*models.Post
	}{Posts: models.GetAllPosts()}

	templ.Execute(writer, data)
}
