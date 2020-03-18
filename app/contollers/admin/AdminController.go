package admin

import (
	"../../core"
	"../../models"
	"../../repositories"
	"net/http"
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
