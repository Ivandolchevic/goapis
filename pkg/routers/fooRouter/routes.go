package fooRouter

import (
	geomapHandler "github.com/Ivandolchevic/goapis/pkg/handlers/fooHandlers/geomapHandler"
	httpUtil "github.com/Ivandolchevic/goapis/pkg/utils/httpUtil"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc httpUtil.APIHandler
}

type Routes []Route

var fooRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		geomapHandler.GetAll,
	},
	Route{
		"GeomapGetAll",
		"GET",
		"/geomaps",
		geomapHandler.GetAll,
	},
	Route{
		"GeomapGet",
		"GET",
		"/geomap/{geomapId}",
		geomapHandler.Get,
	},
	Route{
		"GeomapPut",
		"PUT",
		"/geomaps",
		geomapHandler.Put,
	},
}
