package transports

import (
	"encoding/json"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"
	"github.com/doniacld/outdoorsight/internal/endpoints"
	"github.com/doniacld/outdoorsight/internal/endpoints/addspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/deletespot"
	"github.com/doniacld/outdoorsight/internal/endpoints/getapidoc"
	"github.com/doniacld/outdoorsight/internal/endpoints/getspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/updatespot"
	"github.com/doniacld/outdoorsight/internal/errors"

	"github.com/gorilla/mux"
)

// DecodeRequestGetAPIDoc decodes the request into the internal structure
func DecodeRequestGetAPIDoc(_ *http.Request) (getapidoc.GetAPIDocRequest, *errors.ODSError) {
	return getapidoc.GetAPIDocRequest{}, nil
}

// EncodeResponseGetAPIDoc encodes the response
func EncodeResponseGetAPIDoc(w http.ResponseWriter, resp getapidoc.GetAPIDocResponse) *errors.ODSError {
	// set the response parameters
	w.WriteHeader(getapidoc.GetAPIDocMeta.SuccessCode())
	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeHTML)

	if _, err := w.Write(resp); err != nil {
		return errors.NewFromError(http.StatusInternalServerError, err, "unable to encode getAPIDoc response")
	}
	return nil
}

// DecodeRequestAddSpot decodes the request into the internal structure
func DecodeRequestAddSpot(r *http.Request) (addspot.AddSpotRequest, *errors.ODSError) {
	// decode body
	var addSpotRequest addspot.AddSpotRequest
	if err := json.NewDecoder(r.Body).Decode(&addSpotRequest); err != nil {
		return addspot.AddSpotRequest{}, errors.NewFromError(http.StatusBadRequest, err, "unable to decode addSpotRequest")
	}
	defer r.Body.Close()

	// call validate request
	if err := addSpotRequest.Validate(); err != nil {
		return addspot.AddSpotRequest{}, err.Wrap("addSpotRequest is not valid")
	}
	return addSpotRequest, nil
}

// EncodeResponseAddSpot encodes the response into a JSON file
func EncodeResponseAddSpot(w http.ResponseWriter, resp addspot.AddSpotResponse) *errors.ODSError {
	w.WriteHeader(addspot.AddSpotMeta.SuccessCode())
	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return errors.NewFromError(http.StatusInternalServerError, err, "unable to encode addSpot response")
	}

	return nil
}

// DecodeRequestGetSpot decodes the request into the internal structure
func DecodeRequestGetSpot(r *http.Request) (getspot.GetSpotRequest, *errors.ODSError) {
	vars := mux.Vars(r)
	spotName := vars[endpoints.ParamSpotName]

	return getspot.GetSpotRequest{SpotName: spotName}, nil
}

// EncodeResponseGetSpot encodes the response into a JSON file
func EncodeResponseGetSpot(w http.ResponseWriter, resp getspot.GetSpotResponse) *errors.ODSError {
	w.WriteHeader(getspot.GetSpotMeta.SuccessCode())
	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return errors.NewFromError(http.StatusBadRequest, err, "unable to encode getSpot response")
	}

	return nil
}

// DecodeRequestDeleteSpot decodes the request into the internal structure
func DecodeRequestDeleteSpot(r *http.Request) (deletespot.DeleteSpotRequest, *errors.ODSError) {
	vars := mux.Vars(r)
	spotName := vars[endpoints.ParamSpotName]

	return deletespot.DeleteSpotRequest{SpotName: spotName}, nil
}

// EncodeResponseDeleteSpot encodes the response
func EncodeResponseDeleteSpot(w http.ResponseWriter, _ deletespot.DeleteSpotResponse) *errors.ODSError {
	w.WriteHeader(deletespot.DeleteSpotMeta.SuccessCode())
	return nil
}

// DecodeRequestUpdateSpot decodes the request into the internal structure
func DecodeRequestUpdateSpot(r *http.Request) (updatespot.UpdateSpotRequest, *errors.ODSError) {
	var updateSpotRequest updatespot.UpdateSpotRequest

	if err := json.NewDecoder(r.Body).Decode(&updateSpotRequest); err != nil {
		return updatespot.UpdateSpotRequest{}, errors.NewFromError(http.StatusBadRequest, err, "unable to decode updateSpotRequest")
	}
	defer r.Body.Close()

	// set the path parameter in the updateSpotRequest
	vars := mux.Vars(r)
	updateSpotRequest.Name = vars[endpoints.ParamSpotName]

	// call validate request
	if err := updateSpotRequest.Validate(); err != nil {
		return updatespot.UpdateSpotRequest{}, err.Wrap("updateSpotRequest is not valid")
	}
	return updateSpotRequest, nil
}

// EncodeResponseUpdateSpot encodes the response into a JSON file
func EncodeResponseUpdateSpot(w http.ResponseWriter, resp updatespot.UpdateSpotResponse) *errors.ODSError {
	w.WriteHeader(updatespot.UpdateSpotMeta.SuccessCode())
	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return errors.NewFromError(http.StatusInternalServerError, err, "unable to encode updateSpot response")
	}
	return nil
}
