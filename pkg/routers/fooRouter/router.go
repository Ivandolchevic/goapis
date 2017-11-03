package fooRouter

import (
	"github.com/gorilla/mux"

	interceptors "github.com/Ivandolchevic/goapis/pkg/interceptors"
)

// NewFooRouter define the router for the foo API
func NewFooRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range fooRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(interceptors.Authentifier(interceptors.Logger(route.HandlerFunc)))
	}

	return router
}
