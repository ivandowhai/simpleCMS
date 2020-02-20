package main

import "net/http"
import "./app/contollers"

func main() {
	http.HandleFunc("/", contollers.PostsList)
	http.ListenAndServe(":8080", nil)
}
