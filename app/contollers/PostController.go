package contollers

import (
	"../core"
	"../models"
	"../repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func PostsList(writer http.ResponseWriter, _ *http.Request) {
	postRepository := repositories.PostRepository{}

	posts := postRepository.GetAll()

	response := core.SuccessResponse{Data: posts}
	core.MakeSuccessResponse(writer, &response)
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

	response := core.SuccessResponse{Data: post}
	core.MakeSuccessResponse(writer, &response)
}

func StorePost(writer http.ResponseWriter, request *http.Request) {
	postRepository := repositories.PostRepository{}

	newPost := models.Post{}
	err := json.NewDecoder(request.Body).Decode(&newPost)
	if err != nil {
		response := core.ErrorResponse{Error: err.Error()}
		core.MakeErrorResponse(writer, &response)
		return
	}

	request.ParseForm()
	userId, _ := strconv.ParseInt(request.Form.Get("userId"), 10, 64)
	newPost.AuthorID = uint64(userId)

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
