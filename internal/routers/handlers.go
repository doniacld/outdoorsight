package routers

import (
	"context"
	"fmt"
	"github.com/doniacld/outdoorsight/internal/endpoints/addspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/deletespot"
	"github.com/doniacld/outdoorsight/internal/endpoints/getapidoc"
	"github.com/doniacld/outdoorsight/internal/endpoints/getspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/updatespot"
	transports "github.com/doniacld/outdoorsight/internal/routers/tansports"
	"net/http"
)

// GetAPIDocHandler is the handler of GetAPIDoc endpoint
func GetAPIDocHandler(w http.ResponseWriter, r *http.Request) {
	// decode request
	request, err := transports.DecodeRequestGetAPIDoc(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while decoding getAPIDoc request : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// call the endpoint
	response, err := getapidoc.GetAPIDoc(request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while calling getAPIDoc endpoint : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// encode the response
	err = transports.EncodeResponseGetAPIDoc(w, response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while encoding getAPIDoc response : %s", err.Error()), http.StatusBadRequest)
		return
	}
}

// AddSpotHandler is the handler of AddSpot endpoint
func AddSpotHandler(w http.ResponseWriter, r *http.Request) {
	// decode request
	request, err := transports.DecodeRequestAddSpot(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while decoding addSpot request : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// call the endpoint
	response, err := addspot.AddSpot(request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while calling addSpot endpoint : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// encode the response
	err = transports.EncodeResponseAddSpot(w, response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while encoding addSpot response : %s", err.Error()), http.StatusBadRequest)
		return
	}
}

// GetSpotHandler is the handler of GetSpot endpoint
func GetSpotHandler(w http.ResponseWriter, r *http.Request) {
	// decode request
	request, err := transports.DecodeRequestGetSpot(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while decoding getSpot request : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// call the endpoint
	response, err := getspot.GetSpot(request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while calling getSpot endpoint : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// encode the response
	err = transports.EncodeResponseGetSpot(w, response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while encoding getSpot response : %s", err.Error()), http.StatusBadRequest)
		return
	}
}

// DeleteSpotHandler is the handler of DeleteSpot
// Idempotent endpoint, it will return its success code even the name is not found.
func DeleteSpotHandler(w http.ResponseWriter, r *http.Request) {
	// decode request
	request, err := transports.DecodeRequestDeleteSpot(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while decoding deleteSpot request : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// call the endpoint
	response, err := deletespot.DeleteSpot(context.TODO(), request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while calling deleteSpot endpoint : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// encode the response
	err = transports.EncodeResponseDeleteSpot(w, response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while encoding deleteSpot response : %s", err.Error()), http.StatusBadRequest)
		return
	}
}

// UpdateSpotHandler returns all the details about a given spot
func UpdateSpotHandler(w http.ResponseWriter, r *http.Request) {
	// decode request
	request, err := transports.DecodeRequestUpdateSpot(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while decoding updateSpot request : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// call the endpoint
	response, err := updatespot.UpdateSpot(context.TODO(), request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while decoding updateSpot request : %s", err.Error()), http.StatusBadRequest)
		return
	}

	// encode the response
	err = transports.EncodeResponseUpdateSpot(w, response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while encoding updateSpot response : %s", err.Error()), http.StatusBadRequest)
		return
	}
}
