package auth

import (
	"../../core"
	"../../repositories"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
	// TODO: log all errors

	request.ParseForm()

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

	session.Values["userID"] = user.ID
	session.Values["userRole"] = user.Role
	session.Values["isUserConfirmed"] = user.ConfirmationCode == ""

	err = session.Save(request, writer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	session := core.SessionGet(request, "user")
	session.Values["userID"] = nil
	session.Values["userRole"] = nil
	session.Save(request, writer)

	http.Redirect(writer, request, "/", http.StatusFound)
}
