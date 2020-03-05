package auth

import (
	"../../core"
	"../../models"
	"../../repositories/user"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegisterPage(writer http.ResponseWriter, _ *http.Request) {
	templ := core.GetView("auth/register")

	data := struct{ Result string }{Result: ""}
	templ.Execute(writer, data)
}

func Register(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	// TODO: role from settings
	password, err := bcrypt.GenerateFromPassword([]byte(request.Form.Get("password")), 10)
	if err != nil {
		fmt.Println(err)
	}
	newUser := models.User{Email: request.Form.Get("email"), Name: request.Form.Get("name"), Password: string(password), Role: 2}
	user.CreateUser(newUser)

	templ := core.GetView("auth/register")

	data := struct{ Result string }{Result: "OK"}
	templ.Execute(writer, data)
}
