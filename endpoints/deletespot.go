package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteSpot returns all the details on a given spot
func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotName := vars[ParamSpotName]

	fmt.Println(spotName)
	// set the response parameters
	w.WriteHeader(DeleteSpotMeta.SuccessCode())
}
