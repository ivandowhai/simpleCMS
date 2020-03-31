package auth

import (
	"../../core"
	"../../repositories"
	"../../services"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(writer http.ResponseWriter, request *http.Request) {
	var requestBody loginRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		response := core.ErrorResponse{Error: err.Error()}
		core.MakeErrorResponse(writer, &response)
		return
	}

	loginService := services.LoginService{}
	userRepository := repositories.UserRepository{}

	user, err := userRepository.GetByEmail(requestBody.Email)
	if err != nil {
		response := core.ErrorResponse{Error: err.Error()}
		core.MakeErrorResponse(writer, &response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		response := core.ErrorResponse{Error: err.Error()}
		core.MakeErrorResponse(writer, &response)
		return
	}

	err = loginService.Login(user)
	if err != nil {
		response := core.ErrorResponse{Error: err.Error()}
		core.MakeErrorResponse(writer, &response)
	} else {
		response := core.SuccessResponse{Data: struct {
			Token          string
			ExpirationTime time.Time
		}{loginService.TokenString, loginService.ExpirationTime}}
		core.MakeSuccessResponse(writer, &response)
	}
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	logger := core.Logger{}
	logger.Init()

	session := core.SessionGet(request, "user")
	session.Values["userID"] = nil
	session.Values["userRole"] = nil
	err := session.Save(request, writer)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	http.Redirect(writer, request, "/", http.StatusFound)
}
