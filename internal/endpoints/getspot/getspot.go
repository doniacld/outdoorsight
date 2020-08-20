package getspot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/db"
	"github.com/doniacld/outdoorsight/internal/endpointdef"
	"github.com/doniacld/outdoorsight/internal/endpoints"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// GetSpot returns all the details on a given spot
func GetSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotName := vars[endpoints.ParamSpotName]

	// call the database
	osDB := db.New()
	spotDetails, err := osDB.GetSpot(context.TODO(), spotName)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("unable to get spot %s", spotName)))
	}

	// set the response parameters
	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeJSON)
	w.WriteHeader(GetSpotMeta.SuccessCode())

	if err := json.NewEncoder(w).Encode(spotDetails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
