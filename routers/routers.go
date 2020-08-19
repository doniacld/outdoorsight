package routers

import (
	"net/http"

	"github.com/doniacld/outdoorsight/endpointdef"
	"github.com/doniacld/outdoorsight/endpoints"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router with all the routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		router.
			Methods(route.Verb()).
			Path(route.Path()).
			Name(route.TraceName()).
			Handler(route.HandlerFunc)
	}
	return router
}

// Route holds all the information of a route
type Route struct {
	endpointdef.Meta
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// routes contains all the available endpoints
var routes = Routes{
	Route{
		endpoints.GetSpotMeta,
		endpoints.GetSpot,
	},
	{
		endpoints.AddSpotMeta,
		endpoints.AddSpot,
	},
	{
		endpoints.DeleteSpotMeta,
		endpoints.DeleteSpot,
	},
	{
		endpoints.UpdateSpotMeta,
		endpoints.UpdateSpot,
	},
}
