package admin

import (
	"../../core"
	"../../models"
	"../../repositories"
	"../../services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AdminIndex(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("admin/index", "admin")
	session := core.SessionGet(request, "user")

	data := struct {
		IsAdmin bool
	}{IsAdmin: core.IsAdmin(session.Values["userRole"].(uint8))}

	templ.ExecuteTemplate(writer, "base", data)
}

func UsersList(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("admin/users", "admin")
	usersRepository := repositories.UserRepository{}

	data := struct {
		Users []*models.User
		Roles []core.Role
	}{Users: usersRepository.GetAll(), Roles: core.AllRoles}

	templ.ExecuteTemplate(writer, "base", data)
}

func UserEdit(writer http.ResponseWriter, request *http.Request) {
	logger := core.Logger{}
	logger.Init()
	templ := core.GetView("admin/user_edit", "admin")
	userRepository := repositories.UserRepository{}

	ID, err := strconv.ParseUint(mux.Vars(request)["userId"], 10, 16)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	user, err := userRepository.GetById(ID)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	data := struct {
		User  *models.User
		Roles []core.Role
	}{User: user, Roles: core.AllRoles}

	templ.ExecuteTemplate(writer, "base", data)
}

func UserUpdate(writer http.ResponseWriter, request *http.Request) {
	userService := services.UserService{}
	request.ParseForm()

	userService.ChangeRole(mux.Vars(request)["userId"], request.Form.Get("role"))

	http.Redirect(writer, request, "/admin/users", http.StatusSeeOther)
}

func UserDelete(writer http.ResponseWriter, request *http.Request) {
	userService := services.UserService{}

	userService.Delete(mux.Vars(request)["userId"])

	http.Redirect(writer, request, "/admin/users", http.StatusSeeOther)
}
