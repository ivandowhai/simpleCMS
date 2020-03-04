package auth

import (
	"../../models"
	"../../repositories/user"
	"fmt"
	"html/template"
	"net/http"
)

func RegisterPage(writer http.ResponseWriter, _ *http.Request) {
	templ, err := template.ParseFiles("templates/default/auth/register.html")
	if err != nil {
		fmt.Println(err)
	}

	data := struct{ Test string }{Test: ""}
	templ.Execute(writer, data)
}

func Register(writer http.ResponseWriter, request http.Request) {
	request.ParseForm()
	// TODO: hash password, role from settings
	newUser := models.User{Email: request.Form.Get("email"), Name: request.Form.Get("name"), Password: request.Form.Get("password"), Role: 2}
	user.CreateUser(newUser)

	// TODO: redirect with session message
}
