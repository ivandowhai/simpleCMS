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
			handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Accept", "Accept-Encoding", "content-type"}),
			handlers.AllowedOrigins([]string{settings.FrontHost}),
			handlers.AllowedMethods([]string{route.Method, http.MethodOptions}),
		)(handler)

		r.Methods(route.Method, http.MethodOptions).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return r
}

var routes = Routes{
	Route{
		Name:    "PostsList",
		Method:  http.MethodGet,
		Pattern: "/post",
		Handler: contollers.PostsList,
	},
	Route{
		Name:    "ViewPost",
		Method:  http.MethodGet,
		Pattern: "/post/view/{postId}",
		Handler: contollers.ViewPost,
	},
	Route{
		Name:       "CreatePost",
		Method:     http.MethodGet,
		Pattern:    "/post/create",
		Handler:    contollers.CreatePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "StorePost",
		Method:     http.MethodPost,
		Pattern:    "/post/store",
		Handler:    contollers.StorePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "EditPost",
		Method:     http.MethodGet,
		Pattern:    "/post/edit/{postId}",
		Handler:    contollers.EditPost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "UpdatePost",
		Method:     http.MethodPost,
		Pattern:    "/post/update/{postId}",
		Handler:    contollers.UpdatePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "DeletePost",
		Method:     http.MethodGet,
		Pattern:    "/post/delete/{postId}",
		Handler:    contollers.DeletePost,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, hasAuthorRole, isUserConfirmed},
	},
	Route{
		Name:       "Profile",
		Method:     http.MethodGet,
		Pattern:    "/profile",
		Handler:    profile.ProfilePage,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed},
	},
	Route{
		Name:    "Register",
		Method:  http.MethodPost,
		Pattern: "/register",
		Handler: auth.Register,
	},
	Route{
		Name:    "Login",
		Method:  http.MethodPost,
		Pattern: "/login",
		Handler: auth.Login,
	},
	Route{
		Name:    "Logout",
		Method:  http.MethodGet,
		Pattern: "/logout",
		Handler: auth.Logout,
	},
	Route{
		Name:    "ConfirmAccount",
		Method:  http.MethodGet,
		Pattern: "/confirm",
		Handler: auth.ConfirmAccount,
	},
	Route{
		Name:       "AdminIndex",
		Method:     http.MethodGet,
		Pattern:    "/admin",
		Handler:    admin.AdminIndex,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminOrModeratorRole},
	},
	Route{
		Name:       "AdminUsers",
		Method:     http.MethodGet,
		Pattern:    "/admin/users",
		Handler:    admin.UsersList,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminRole},
	},
	Route{
		Name:       "AdminUserEdit",
		Method:     http.MethodGet,
		Pattern:    "/admin/users/edit/{userId}",
		Handler:    admin.UserEdit,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminRole},
	},
	Route{
		Name:       "AdminUserUpdate",
		Method:     http.MethodPost,
		Pattern:    "/admin/users/update/{userId}",
		Handler:    admin.UserUpdate,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminRole},
	},
	Route{
		Name:       "AdminUserDelete",
		Method:     http.MethodGet,
		Pattern:    "/admin/users/delete/{userId}",
		Handler:    admin.UserDelete,
		Middleware: []mux.MiddlewareFunc{isUserLoggedMiddleware, isUserConfirmed, hasAdminRole},
	},
}
