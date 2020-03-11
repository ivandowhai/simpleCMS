package routes

import (
	"../core"
	"net/http"
)

func isUserLoggedMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session := core.SessionGet(request, "user")
		if session.Values["userID"] == nil {
			http.Redirect(writer, request, "/profile", http.StatusSeeOther)
		}
		handler.ServeHTTP(writer, request)
	})
}

func hasAuthorRole(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session := core.SessionGet(request, "user")
		if !core.CanUserPost(session.Values["userRole"].(uint8)) {
			http.Redirect(writer, request, "/profile", http.StatusSeeOther)
		}
		handler.ServeHTTP(writer, request)
	})
}
