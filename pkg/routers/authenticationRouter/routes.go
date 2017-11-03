package authenticationRouter

import (
	basicAuthenticationHandler "github.com/Ivandolchevic/goapis/pkg/handlers/authenticationHandlers/basicAuthenticationHandler"
	httpUtil "github.com/Ivandolchevic/goapis/pkg/utils/httpUtil"
)

// Route gives the struct of a route toward a method
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc httpUtil.APIHandler
}

// Routes is an array of Route
type Routes []Route

var authenticationroutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		basicAuthenticationHandler.GetAll,
	},
	Route{
		"BasicsPut",
		"PUT",
		"/basicauthentications",
		basicAuthenticationHandler.Put,
	},
	Route{
		"Index",
		"GET",
		"/basicauthentications",
		basicAuthenticationHandler.GetAll,
	},
	Route{
		"BasicsGet",
		"GET",
		"/basicauthentications/{id}",
		basicAuthenticationHandler.Get,
	},
}
