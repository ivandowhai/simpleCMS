package main

import (
	"./app/routes"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./templates/default/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.Handle("/", routes.GetRouter())
	http.ListenAndServe(":8080", nil)
}
