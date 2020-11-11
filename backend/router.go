package backend

import (
	"net/http"
	"u_admin/logger"

	"github.com/gorilla/mux"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"users",
		"GET",
		"/users",
		GetAllUsers,
	},

	Route{
		"create user",
		"POST",
		"/users",
		AddUser,
	},

	Route{
		"get user",
		"GET",
		"/users/{id}",
		GetUser,
	},

	Route{
		"update user",
		"PUT",
		"/users/{id}",
		UpdateUser,
	},

	Route{
		"delete user",
		"DELETE",
		"/users/{id}",
		DeleteUser,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
