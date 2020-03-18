package admin

import (
	"../../core"
	"net/http"
)

func AdminIndex(writer http.ResponseWriter, request *http.Request) {
	templ := core.GetView("admin/index", "admin")

	data := struct {
		Test string
	}{""}

	templ.ExecuteTemplate(writer, "base", data)
}
