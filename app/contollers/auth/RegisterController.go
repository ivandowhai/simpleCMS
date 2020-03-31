package auth

import (
	"../../core"
	"../../helpers"
	"../../models"
	"../../repositories"
	"../../services"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	logger := core.Logger{}
	logger.Init()

	err := request.ParseForm()
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Form.Get("password")), 10)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	confirmationCode := helpers.RandStringBytes(30)
	userRepository := repositories.UserRepository{}

	// TODO: role from settings
	newUser := models.User{
		Email:    request.Form.Get("email"),
		Name:     request.Form.Get("name"),
		Password: string(password), Role: 2,
		ConfirmationCode: sql.NullString{String: helpers.RandStringBytes(30), Valid: true},
	}
	userRepository.CreateUser(newUser)

	templ := core.GetView("auth/register", "auth")

	go services.SendConfirmationEmail(request.Form.Get("email"), confirmationCode)

	data := struct{ Result string }{Result: "OK"}
	templ.ExecuteTemplate(writer, "base", data)
}

func ConfirmAccount(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("auth/confirm", "auth")

	userRepository := repositories.UserRepository{}
	userModel, err := userRepository.GetByEmail(request.URL.Query().Get("email"))
	fmt.Println(userModel)
	if err != nil {
		data := struct{ Error string }{Error: err.Error()}
		templ.ExecuteTemplate(writer, "base", data)
		return
	}

	if userModel.ConfirmationCode.String != request.URL.Query().Get("code") {
		data := struct{ Error string }{Error: "Wrong code."}
		templ.ExecuteTemplate(writer, "base", data)
		return
	}

	userRepository.Confirm(userModel)

	data := struct{ Error string }{Error: ""}
	templ.ExecuteTemplate(writer, "base", data)
}
