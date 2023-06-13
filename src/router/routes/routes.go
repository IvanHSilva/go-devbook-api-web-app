package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
