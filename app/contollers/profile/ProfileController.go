package profile

import (
	"../../core"
	"../../models"
	"../../repositories/post"
	"../../repositories/user"
	"fmt"
	"net/http"
)

func ProfilePage(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("profile/index")

	session := core.SessionGet(request, "user")

	data := struct {
		User  *models.User
		Posts []*models.Post
	}{User: nil, Posts: []*models.Post{}}

	if session.Values["userID"] != nil {
		userID := session.Values["userID"].(uint64)
		user, err := user.GetById(userID)

		if err != nil {
			fmt.Println(err)
		}

		data.User = user
		data.Posts = post.GetByUser(userID)
	}

	templ.ExecuteTemplate(writer, "base", data)
}
