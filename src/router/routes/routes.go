package routes

import "net/http"

// Represent all routes of the API
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	AuthRequire bool
}
