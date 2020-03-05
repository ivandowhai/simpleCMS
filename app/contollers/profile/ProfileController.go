package profile

import (
	"../../core"
	"../../models"
	"../../repositories/user"
	"fmt"
	"html/template"
	"net/http"
)

func ProfilePage(writer http.ResponseWriter, request *http.Request) {
	templ, err := template.ParseFiles("templates/default/profile/index.html")
	if err != nil {
		fmt.Println(err)
	}

	session, err := core.Store.Get(request, "user")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		User *models.User
	}{User: nil}

	if session.Values["userID"] != nil {
		user, err := user.GetById(session.Values["userID"].(uint64))

		if err != nil {
			fmt.Println(err)
		}

		data.User = user
	}

	templ.Execute(writer, data)
}
