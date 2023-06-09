package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Represent all routes of the API
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	AuthRequire bool
}

// Insert all routes in router
func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, routesPosts...)

	for _, route := range routes {
		if route.AuthRequire {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function))).
				Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).
				Methods(route.Method)
		}
	}
	return r
}
