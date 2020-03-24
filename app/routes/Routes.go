package routes

import (
	"../contollers"
	"../contollers/admin"
	"../contollers/auth"
	"../contollers/profile"
	"../core"
	"github.com/gorilla/handlers"
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

		var settings = core.GetSettings()

		handler = handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"}),
			handlers.AllowedOrigins([]string{settings.FrontHost}),
			handlers.AllowedMethods([]string{route.Method, "Options"}),
		)(handler)

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
		Pattern: "/post",
		Handler: contollers.PostsList,
	},
	Route{
		Name:    "ViewPost",
		Method:  "GET",
		Pattern: "/post/view/{postId}",
		Handler: contollers.ViewPost,
	},
	Route{
		Name:       "CreatePost",
		Method:     "GET",
		Pattern:    "/post/create",
		Handler:    contollers.CreatePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "StorePost",
		Method:     "POST",
		Pattern:    "/post/store",
		Handler:    contollers.StorePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "EditPost",
		Method:     "GET",
		Pattern:    "/post/edit/{postId}",
		Handler:    contollers.EditPost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "UpdatePost",
		Method:     "POST",
		Pattern:    "/post/update/{postId}",
		Handler:    contollers.UpdatePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "DeletePost",
		Method:     "GET",
		Pattern:    "/post/delete/{postId}",
		Handler:    contollers.DeletePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "Profile",
		Method:     "GET",
		Pattern:    "/profile",
		Handler:    profile.ProfilePage,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed},
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
	Route{
		Name:    "Logout",
		Method:  "GET",
		Pattern: "/logout",
		Handler: auth.Logout,
	},
	Route{
		Name:    "ConfirmAccount",
		Method:  "GET",
		Pattern: "/confirm",
		Handler: auth.ConfirmAccount,
	},
	Route{
		Name:       "AdminIndex",
		Method:     "GET",
		Pattern:    "/admin",
		Handler:    admin.AdminIndex,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminOrModeratorRole},
	},
	Route{
		Name:       "AdminUsers",
		Method:     "GET",
		Pattern:    "/admin/users",
		Handler:    admin.UsersList,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminRole},
	},
	Route{
		Name:       "AdminUserEdit",
		Method:     "GET",
		Pattern:    "/admin/users/edit/{userId}",
		Handler:    admin.UserEdit,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminRole},
	},
	Route{
		Name:       "AdminUserUpdate",
		Method:     "POST",
		Pattern:    "/admin/users/update/{userId}",
		Handler:    admin.UserUpdate,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminRole},
	},
	Route{
		Name:       "AdminUserDelete",
		Method:     "GET",
		Pattern:    "/admin/users/delete/{userId}",
		Handler:    admin.UserDelete,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminRole},
	},
}
