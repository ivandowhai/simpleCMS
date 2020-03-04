package main

import (
	"github.com/gorilla/mux"
	"net/http"
)
import "./app/contollers"
import "./app/contollers/auth"

func main() {
	// TODO: use gorilla mux, move routing to separate module
	r := mux.NewRouter()
	r.HandleFunc("/", contollers.PostsList)
	r.HandleFunc("/posts", contollers.PostsList)
	r.HandleFunc("/posts/{postId}", contollers.ViewPost)
	r.HandleFunc("/register", auth.RegisterPage)
	r.HandleFunc("/register-submit", auth.Register)
	r.HandleFunc("/login", auth.LoginPage)
	r.HandleFunc("/login-submit", auth.Login)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
