package auth

import (
	"../../core"
	"../../helpers"
	"../../models"
	"../../repositories"
	"../../services"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type confirmRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
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
	var requestBody confirmRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		response := core.ErrorResponse{Error: err.Error()}
		core.MakeErrorResponse(writer, &response)
		return
	}

	userRepository := repositories.UserRepository{}
	userModel, err := userRepository.GetByEmail(requestBody.Email)
	if err != nil {
		response := core.ErrorResponse{Error: err.Error()}
		core.MakeErrorResponse(writer, &response)
		return
	}

	if userModel.ConfirmationCode.String != requestBody.Code {
		response := core.ErrorResponse{Error: "Wrong code."}
		core.MakeErrorResponse(writer, &response)
		return
	}

	userRepository.Confirm(userModel)

	response := core.SuccessResponse{Data: struct {
		Result string
	}{Result: "Your account was confirmed"}}
	core.MakeSuccessResponse(writer, &response)
}
