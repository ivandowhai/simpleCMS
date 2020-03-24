package routes

import (
	"../core"
	"../services"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func isUserLoggedMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokenString := request.Header.Get("token")
		if tokenString == "" {
			// If the cookie is not set, return an unauthorized status
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := &services.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return token, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			writer.WriteHeader(http.StatusUnauthorized)
			return
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

func isUserConfirmed(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session := core.SessionGet(request, "user")
		if session.Values["isUserConfirmed"] == false {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
		}
		handler.ServeHTTP(writer, request)
	})
}

func hasAdminRole(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session := core.SessionGet(request, "user")
		if !core.IsAdmin(session.Values["userRole"].(uint8)) {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
		}
		handler.ServeHTTP(writer, request)
	})
}

func hasAdminOrModeratorRole(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session := core.SessionGet(request, "user")
		if !core.IsAdminOrModer(session.Values["userRole"].(uint8)) {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
		}
		handler.ServeHTTP(writer, request)
	})
}
