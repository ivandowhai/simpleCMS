package profile

import (
	"../../core"
	"../../models"
	"../../repositories"
	"fmt"
	"net/http"
)

func ProfilePage(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("profile/index", "main")
	userRepository := repositories.UserRepository{}
	postRepository := repositories.PostRepository{}

	session := core.SessionGet(request, "user")

	data := struct {
		User     *models.User
		Posts    []*models.Post
		IsLogged bool
	}{User: nil, Posts: []*models.Post{}, IsLogged: session.Values["userID"] != nil}

	if session.Values["userID"] != nil {
		userID := session.Values["userID"].(uint64)
		user, err := userRepository.GetById(userID)

		if err != nil {
			fmt.Println(err)
		}

		data.User = user
		data.Posts = postRepository.GetByUser(userID)
	}

	templ.ExecuteTemplate(writer, "base", data)
}
