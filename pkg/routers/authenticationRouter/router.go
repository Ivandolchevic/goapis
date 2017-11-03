package authenticationRouter

import (
	"github.com/gorilla/mux"

	interceptors "github.com/Ivandolchevic/goapis/pkg/interceptors"
	httpUtil "github.com/Ivandolchevic/goapis/pkg/utils/httpUtil"
)

// NewAuthenticationRouter provides a router for the authentication API
func NewAuthenticationRouter() *mux.Router {

	router := httpUtil.NewRouter()

	for _, route := range authenticationroutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(interceptors.Authentifier(route.HandlerFunc))
	}

	return router
}
