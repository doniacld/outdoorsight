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

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// DecodeRequestGetAPIDoc decodes the request into the internal structure
func DecodeRequestGetAPIDoc(_ *http.Request) (getapidoc.GetAPIDocRequest, error) {
	return getapidoc.GetAPIDocRequest{}, nil
}

// EncodeResponseGetAPIDoc encodes the response
func EncodeResponseGetAPIDoc(w http.ResponseWriter, resp getapidoc.GetAPIDocResponse) error {
	// set the response parameters
	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeHTML)
	w.WriteHeader(getapidoc.GetAPIDocMeta.SuccessCode())

	if _, err := w.Write(resp); err != nil {
		return errors.Wrap(err, "unable to encode getAPIDoc response")
	}
	return nil
}

// DecodeRequestAddSpot decodes the request into the internal structure
func DecodeRequestAddSpot(r *http.Request) (addspot.AddSpotRequest, error) {
	var addSpotRequest addspot.AddSpotRequest
	if err := json.NewDecoder(r.Body).Decode(&addSpotRequest); err != nil {
		return addspot.AddSpotRequest{}, errors.Wrap(err, "unable to decode addSpotRequest")
	}
	defer r.Body.Close()

	return addSpotRequest, nil
}

// EncodeResponseAddSpot encodes the response into a JSON file
func EncodeResponseAddSpot(w http.ResponseWriter, resp addspot.AddSpotResponse) error {
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return errors.Wrap(err, "unable to encode addSpot response")
	}

	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
	w.WriteHeader(addspot.AddSpotMeta.SuccessCode())
	return nil
}

// DecodeRequestGetSpot decodes the request into the internal structure
func DecodeRequestGetSpot(r *http.Request) (getspot.GetSpotRequest, error) {
	vars := mux.Vars(r)
	spotName := vars[endpoints.ParamSpotName]

	return getspot.GetSpotRequest{SpotName: spotName}, nil
}

// EncodeResponseGetSpot encodes the response into a JSON file
func EncodeResponseGetSpot(w http.ResponseWriter, resp getspot.GetSpotResponse) error {
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return errors.Wrap(err, "unable to encode getSpot response")
	}

	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
	w.WriteHeader(getspot.GetSpotMeta.SuccessCode())
	return nil
}

// DecodeRequestDeleteSpot decodes the request into the internal structure
func DecodeRequestDeleteSpot(r *http.Request) (deletespot.DeleteSpotRequest, error) {
	vars := mux.Vars(r)
	spotName := vars[endpoints.ParamSpotName]

	return deletespot.DeleteSpotRequest{SpotName: spotName}, nil
}

// EncodeResponseDeleteSpot encodes the response
func EncodeResponseDeleteSpot(w http.ResponseWriter, _ deletespot.DeleteSpotResponse) error {
	w.WriteHeader(getspot.GetSpotMeta.SuccessCode())
	return nil
}

// DecodeRequestUpdateSpot decodes the request into the internal structure
func DecodeRequestUpdateSpot(r *http.Request) (updatespot.UpdateSpotRequest, error) {
	var updateSpotRequest updatespot.UpdateSpotRequest

	vars := mux.Vars(r)
	updateSpotRequest.Name = vars[endpoints.ParamSpotName]

	if err := json.NewDecoder(r.Body).Decode(&updateSpotRequest); err != nil {
		return updatespot.UpdateSpotRequest{}, errors.Wrap(err, "unable to decode updateSpotRequest")
	}
	defer r.Body.Close()

	return updateSpotRequest, nil
}

// EncodeResponseUpdateSpot encodes the response into a JSON file
func EncodeResponseUpdateSpot(w http.ResponseWriter, resp updatespot.UpdateSpotResponse) error {
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return errors.Wrap(err, "unable to encode updateSpot response")
	}

	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
	w.WriteHeader(updatespot.UpdateSpotMeta.SuccessCode())
	return nil
}
