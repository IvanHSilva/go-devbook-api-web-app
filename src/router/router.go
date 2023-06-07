package router

import "github.com/gorilla/mux"

// Return a router with configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
