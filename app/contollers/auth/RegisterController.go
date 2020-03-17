package auth

import (
	"../../core"
	"../../models"
	"../../repositories/user"
	"../../services"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
)

func RegisterPage(writer http.ResponseWriter, _ *http.Request) {
	templ := core.GetView("auth/register", "auth")

	data := struct{ Result string }{Result: ""}
	templ.ExecuteTemplate(writer, "base", data)
}

func Register(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	// TODO: role from settings
	password, err := bcrypt.GenerateFromPassword([]byte(request.Form.Get("password")), 10)
	if err != nil {
		fmt.Println(err)
	}

	confirmationCode := randStringBytes(30)

	newUser := models.User{Email: request.Form.Get("email"), Name: request.Form.Get("name"), Password: string(password), Role: 2, ConfirmationCode: confirmationCode}
	user.CreateUser(newUser)

	templ := core.GetView("auth/register", "auth")

	services.SendConfirmationEmail(request.Form.Get("email"), confirmationCode)

	data := struct{ Result string }{Result: "OK"}
	templ.ExecuteTemplate(writer, "base", data)
}

func ConfirmAccount(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("auth/confirm", "auth")
	params := mux.Vars(request)

	userModel, err := user.GetByEmail(params["email"])
	if err != nil {
		data := struct{ Error string }{Error: err.Error()}
		templ.ExecuteTemplate(writer, "base", data)
		return
	}

	if userModel.ConfirmationCode != params["confirmationCode"] {
		data := struct{ Error string }{Error: "Wrong code."}
		templ.ExecuteTemplate(writer, "base", data)
		return
	}

	user.Confirm(userModel)

	data := struct{ Error string }{Error: nil}
	templ.ExecuteTemplate(writer, "base", data)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
