package routers

import (
	"context"
	"github.com/doniacld/outdoorsight/internal/db"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpoints/addspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/deletespot"
	"github.com/doniacld/outdoorsight/internal/endpoints/getapidoc"
	"github.com/doniacld/outdoorsight/internal/endpoints/getspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/updatespot"
	"github.com/doniacld/outdoorsight/internal/errors"
	"github.com/doniacld/outdoorsight/internal/routers/tansports"
)

// GetAPIDocHandler is the handler of GetAPIDoc endpoint
func GetAPIDocHandler(w http.ResponseWriter, r *http.Request) {
	// decode request
	request, err := transports.DecodeRequestGetAPIDoc(r)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// call the endpoint
	response, err := getapidoc.GetAPIDoc(request)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// encode the response
	err = transports.EncodeResponseGetAPIDoc(w, response)
	if err != nil {
		err.HTTPError(w)
		return
	}
}

// AddSpotHandler is the handler of AddSpot endpoint
func AddSpotHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	// decode request
	request, err := transports.DecodeRequestAddSpot(r)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// create a new connection to database
	odsDB, dbErr := db.New()
	if dbErr != nil {
		er := errors.NewFromError(http.StatusInternalServerError, err, "error while creating a new connection to the database")
		er.HTTPError(w)
		return
	}

	// call the endpoint
	response, err := addspot.AddSpot(ctx, request, odsDB)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// encode the response
	err = transports.EncodeResponseAddSpot(w, response)
	if err != nil {
		err.HTTPError(w)
		return
	}
}

// GetSpotHandler is the handler of GetSpot endpoint
func GetSpotHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	// decode request
	request, err := transports.DecodeRequestGetSpot(r)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// create a new connection to database
	odsDB, dbErr := db.New()
	if dbErr != nil {
		er := errors.NewFromError(http.StatusInternalServerError, err, "error while creating a new connection to the database")
		er.HTTPError(w)
		return
	}

	// call the endpoint
	response, err := getspot.GetSpot(ctx, request, odsDB)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// encode the response
	err = transports.EncodeResponseGetSpot(w, response)
	if err != nil {
		err.HTTPError(w)
		return
	}
}

// DeleteSpotHandler is the handler of DeleteSpot
// Idempotent endpoint, it will return its success code even the name is not found.
func DeleteSpotHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	// decode request
	request, err := transports.DecodeRequestDeleteSpot(r)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// create a new connection to database
	odsDB, dbErr := db.New()
	if dbErr != nil {
		er := errors.NewFromError(http.StatusInternalServerError, err, "error while creating a new connection to the database")
		er.HTTPError(w)
		return
	}

	// call the endpoint
	response, err := deletespot.DeleteSpot(ctx, request, odsDB)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// encode the response
	err = transports.EncodeResponseDeleteSpot(w, response)
	if err != nil {
		err.HTTPError(w)
		return
	}
}

// UpdateSpotHandler returns all the details about a given spot
func UpdateSpotHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// decode request
	request, err := transports.DecodeRequestUpdateSpot(r)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// create a new connection to database
	odsDB, dbErr := db.New()
	if dbErr != nil {
		er := errors.NewFromError(http.StatusInternalServerError, err, "error while creating a new connection to the database")
		er.HTTPError(w)
		return
	}

	// call the endpoint
	response, err := updatespot.UpdateSpot(ctx, request, odsDB)
	if err != nil {
		err.HTTPError(w)
		return
	}

	// encode the response
	err = transports.EncodeResponseUpdateSpot(w, response)
	if err != nil {
		err.HTTPError(w)
		return
	}
}
