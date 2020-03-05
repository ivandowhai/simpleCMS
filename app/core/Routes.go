package core

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
	r.HandleFunc("/posts/{postId}", contollers.ViewPost)
	r.HandleFunc("/register", auth.RegisterPage)
	r.HandleFunc("/register-submit", auth.Register)
	r.HandleFunc("/login", auth.LoginPage)
	r.HandleFunc("/login-submit", auth.Login)
	r.HandleFunc("/profile", profile.ProfilePage)

	return r
}
