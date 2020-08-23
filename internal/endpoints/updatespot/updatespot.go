package updatespot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/db"
	"github.com/doniacld/outdoorsight/internal/endpoints"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// UpdateSpot returns all the details about a given spot
func UpdateSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotName := vars[endpoints.ParamSpotName]

	// decode the body into spotDetails structure
	var spotDetailsDB db.SpotDetails
	if err := json.NewDecoder(r.Body).Decode(&spotDetailsDB); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call the database
	osDB := db.New()
	if err := osDB.UpdateSpot(context.TODO(), spotName, spotDetailsDB); err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("unable to update spot %s", spotName)))
	}

	// get the updated spot
	if spotName != spotDetailsDB.Name {
		spotName = spotDetailsDB.Name
	}
	res, err := osDB.GetSpot(context.TODO(), spotName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set the response parameters
	w.WriteHeader(UpdateSpotMeta.SuccessCode())
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
