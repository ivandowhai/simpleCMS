package auth

import (
	"../../core"
	"../../repositories/user"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

func LoginPage(writer http.ResponseWriter, _ *http.Request) {
	templ, err := template.ParseFiles("templates/default/auth/login.html")
	if err != nil {
		fmt.Println(err)
	}

	data := struct{ Result string }{Result: ""}
	templ.Execute(writer, data)
}

func Login(writer http.ResponseWriter, request *http.Request) {
	// TODO: log all errors
	templ, err := template.ParseFiles("templates/default/auth/login.html")
	if err != nil {
		fmt.Println(err)
	}

	request.ParseForm()

	data := struct {
		Result string
		UserID uint64
	}{Result: "", UserID: 0}

	user, err := user.GetByEmail(request.Form.Get("email"))
	if err != nil {
		fmt.Println(err.Error())
		data.Result = err.Error()
		templ.Execute(writer, data)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Form.Get("password"))); err != nil {
		fmt.Println(err.Error())
		data.Result = "password is wrong"
		templ.Execute(writer, data)
		return
	}

	session, err := core.Store.Get(request, "user")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["userID"] = user.ID
	err = session.Save(request, writer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Result = "OK"
	templ.Execute(writer, data)
}
