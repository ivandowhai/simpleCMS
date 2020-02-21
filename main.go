package main

import (
	"github.com/gorilla/mux"
	"net/http"
)
import "./app/contollers"

func main() {
	// TODO: use gorilla mux, move routing to separate module
	r := mux.NewRouter()
	r.HandleFunc("/", contollers.PostsList)
	r.HandleFunc("/posts", contollers.PostsList)
	r.HandleFunc("/posts/{postId}", contollers.ViewPost)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
