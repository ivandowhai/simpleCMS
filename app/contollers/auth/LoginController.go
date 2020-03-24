package auth

import (
	"../../core"
	"../../repositories"
	"../../services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func LoginPage(writer http.ResponseWriter, request *http.Request) {
	session := core.SessionGet(request, "user")
	templ := core.GetView("auth/login", "auth")

	data := struct{ Result string }{""}
	if len(session.Flashes()) > 0 {
		data.Result = session.Flashes()[0].(string)
	}

	templ.ExecuteTemplate(writer, "base", data)
}

func Login(writer http.ResponseWriter, request *http.Request) {
	session := core.SessionGet(request, "user")
	logger := core.Logger{}
	logger.Init()
	loginService := services.LoginService{}

	err := request.ParseForm()
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	userRepository := repositories.UserRepository{}

	user, err := userRepository.GetByEmail(request.Form.Get("email"))
	if err != nil {
		session.AddFlash(err.Error())
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Form.Get("password"))); err != nil {
		session.AddFlash("password is wrong")
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
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
