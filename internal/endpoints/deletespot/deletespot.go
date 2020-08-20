package deletespot

import (
	"fmt"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpoints"

	"github.com/gorilla/mux"
)

// DeleteSpot returns all the details on a given spot
// Idempotent endpoint, it will return its success code even the name is not found.
func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spotName := vars[endpoints.ParamSpotName]

	fmt.Println(spotName)
	// set the response parameters
	w.WriteHeader(DeleteSpotMeta.SuccessCode())
}
