package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doniacld/outdoorsight/db"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// UpdateSpot returns all the details on a given spot
func UpdateSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotName := vars[ParamSpotName]

	// decode the body into spotDetails structure
	var spotDetailsDB db.SpotDetails
	if err := json.NewDecoder(r.Body).Decode(&spotDetailsDB); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call the database
	osDB := db.New()
	err := osDB.UpdateSpot(context.TODO(), spotName, spotDetailsDB)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("unable to update spot %s", spotName)))
	}

	// set the response parameters
	w.WriteHeader(UpdateSpotMeta.SuccessCode())
}
