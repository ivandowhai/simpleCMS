package routes

import (
	"../contollers"
	"../contollers/auth"
	"../contollers/profile"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", contollers.PostsList)
	r.HandleFunc("/posts", contollers.PostsList)
	r.HandleFunc("/posts/view/{postId}", contollers.ViewPost)
	r.HandleFunc("/posts/create", contollers.CreatePost)
	r.HandleFunc("/posts/store", contollers.StorePost)
	r.HandleFunc("/posts/edit/{postId}", contollers.EditPost)
	r.HandleFunc("/posts/update/{postId}", contollers.UpdatePost)
	r.HandleFunc("/posts/delete/{postId}", contollers.DeletePost)
	r.HandleFunc("/register", auth.RegisterPage)
	r.HandleFunc("/register-submit", auth.Register)
	r.HandleFunc("/login", auth.LoginPage)
	r.HandleFunc("/login-submit", auth.Login)
	r.HandleFunc("/profile", profile.ProfilePage)

	return r
}
