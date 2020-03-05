package main

import (
	"./app/core"
	"net/http"
)

func main() {
	http.Handle("/", core.GetRouter())
	http.ListenAndServe(":8080", nil)
}
