package main

import (
	"./app/routes"
	"net/http"
)

func main() {
	http.Handle("/", routes.GetRouter())
	http.ListenAndServe(":8080", nil)
}
