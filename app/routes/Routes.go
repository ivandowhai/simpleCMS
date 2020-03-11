package routes

import (
	"../contollers"
	"../contollers/auth"
	"../contollers/profile"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware []mux.MiddlewareFunc
}

type Routes []Route

func GetRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler = route.Handler
		for _, middleware := range route.Middleware {
			handler = middleware(handler)
		}

		r.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return r
}

var routes = Routes{
	Route{
		Name:    "Index",
		Method:  "GET",
		Pattern: "/",
		Handler: contollers.PostsList,
	},
	Route{
		Name:    "PostsList",
		Method:  "GET",
		Pattern: "/posts",
		Handler: contollers.PostsList,
	},
	Route{
		Name:    "ViewPost",
		Method:  "GET",
		Pattern: "/posts/view/{postId}",
		Handler: contollers.ViewPost,
	},
	Route{
		Name:       "CreatePost",
		Method:     "GET",
		Pattern:    "/posts/create",
		Handler:    contollers.CreatePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole},
	},
	Route{
		Name:       "StorePost",
		Method:     "POST",
		Pattern:    "/posts/store",
		Handler:    contollers.StorePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole},
	},
	Route{
		Name:       "EditPost",
		Method:     "GET",
		Pattern:    "/posts/edit/{postId}",
		Handler:    contollers.EditPost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole},
	},
	Route{
		Name:       "UpdatePost",
		Method:     "POST",
		Pattern:    "/posts/update/{postId}",
		Handler:    contollers.UpdatePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole},
	},
	Route{
		Name:       "DeletePost",
		Method:     "GET",
		Pattern:    "/posts/delete/{postId}",
		Handler:    contollers.DeletePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole},
	},
	Route{
		Name:       "Profile",
		Method:     "GET",
		Pattern:    "/profile",
		Handler:    profile.ProfilePage,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware},
	},
	Route{
		Name:    "RegisterPage",
		Method:  "GET",
		Pattern: "/register",
		Handler: auth.RegisterPage,
	},
	Route{
		Name:    "Register",
		Method:  "POST",
		Pattern: "/register-submit",
		Handler: auth.Register,
	},
	Route{
		Name:    "LoginPage",
		Method:  "GET",
		Pattern: "/login",
		Handler: auth.LoginPage,
	},
	Route{
		Name:    "Login",
		Method:  "POST",
		Pattern: "/login-submit",
		Handler: auth.Login,
	},
}
