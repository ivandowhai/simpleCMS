package main

import "net/http"
import "./app/contollers"

func main() {
	// TODO: use gorilla mux, move routing to separate module
	http.HandleFunc("/posts", contollers.PostsList)
	http.HandleFunc("/posts/view", contollers.ViewPost)
	http.ListenAndServe(":8080", nil)
}
