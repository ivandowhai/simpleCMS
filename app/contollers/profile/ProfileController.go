package profile

import (
	"../../core"
	"../../models"
	"../../repositories"
	"net/http"
)

func ProfilePage(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("profile/index", "main")
	userRepository := repositories.UserRepository{}
	postRepository := repositories.PostRepository{}
	logger := core.Logger{}
	logger.Init()

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
			logger.WriteLog(err.Error(), "error")
		}

		data.User = user
		data.Posts = postRepository.GetByUser(userID)
	}

	templ.ExecuteTemplate(writer, "base", data)
}
