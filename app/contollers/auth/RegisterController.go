package auth

import (
	"../../core"
	"../../helpers"
	"../../models"
	"../../repositories"
	"../../services"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(writer http.ResponseWriter, request *http.Request) {
	logger := core.Logger{}
	logger.Init()

	var requestBody registerRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		response := core.ErrorResponse{Error: err.Error()}
		core.MakeErrorResponse(writer, &response)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	confirmationCode := helpers.RandStringBytes(30)
	userRepository := repositories.UserRepository{}

	// TODO: role from settings
	newUser := models.User{
		Email:    requestBody.Email,
		Name:     requestBody.Name,
		Password: string(password), Role: 2,
		ConfirmationCode: sql.NullString{String: helpers.RandStringBytes(30), Valid: true},
	}
	userRepository.CreateUser(newUser)

	go services.SendConfirmationEmail(request.Form.Get("email"), confirmationCode)

	response := core.SuccessResponse{Data: struct {
		Result string
	}{Result: "Confirmation code was sent to your email"}}
	core.MakeSuccessResponse(writer, &response)
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
