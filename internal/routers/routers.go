package routers

import (
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"
	"github.com/doniacld/outdoorsight/internal/endpoints/addspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/deletespot"
	"github.com/doniacld/outdoorsight/internal/endpoints/getapidoc"
	"github.com/doniacld/outdoorsight/internal/endpoints/getspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/updatespot"

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
		getapidoc.GetAPIDocMeta,
		getapidoc.GetAPIDoc,
	},
	{
		getspot.GetSpotMeta,
		getspot.GetSpot,
	},
	{
		addspot.AddSpotMeta,
		addspot.AddSpot,
	},
	{
		deletespot.DeleteSpotMeta,
		deletespot.DeleteSpot,
	},
	{
		updatespot.UpdateSpotMeta,
		updatespot.UpdateSpot,
	},
}
