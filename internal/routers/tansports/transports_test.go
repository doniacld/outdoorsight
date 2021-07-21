package transports

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/doniacld/outdoorsight/internal/endpoints"
	"github.com/doniacld/outdoorsight/internal/endpoints/addspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/deletespot"
	"github.com/doniacld/outdoorsight/internal/endpoints/getspot"
	"github.com/doniacld/outdoorsight/internal/endpoints/updatespot"
	"github.com/doniacld/outdoorsight/internal/spot"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDecodeRequestAddSpot(t *testing.T) {
	// nominal case
	body := `{"name":"LesSurplombs","routes":[{"name":"AlineLaMaline","level":"5a","points":5,"information":"Not so easy"}],"metadata":{"info":"Park at the canoe spot"}}`
	r := httptest.NewRequest(http.MethodPost, "/spots", strings.NewReader(body))

	request, _ := DecodeRequestAddSpot(r)
	assert.Equal(t, request, addspot.AddSpotRequest{
		Name:     "LesSurplombs",
		Routes:   []spot.Route{{Name: "AlineLaMaline", Level: "5a", Points: 5, Information: "Not so easy"}},
		Metadata: map[string]interface{}{"info": "Park at the canoe spot"}})

	// invalid case : a space is present in the name
	body = `{"name":"Les Surplombs","routes":[{"name":"AlineLaMaline","level":"5a","points":5,"information":"Not so easy"}],"metadata":{"info":"Park at the canoe spot"}}`
	r = httptest.NewRequest(http.MethodPost, "/spots", strings.NewReader(body))

	_, err := DecodeRequestAddSpot(r)
	assert.Contains(t, err.Message, "spot name 'Les Surplombs' contains at least one space")
	assert.Equal(t, http.StatusBadRequest, err.HTTPCode)

	// invalid case : no body
	r = httptest.NewRequest(http.MethodPost, "/spots", nil)
	_, err = DecodeRequestAddSpot(r)
	assert.Contains(t, err.Message, "unable to decode addSpotRequest")
	assert.Equal(t, http.StatusBadRequest, err.HTTPCode)
}

func TestDecodeRequestGetSpot(t *testing.T) {
	// nominal case
	r := httptest.NewRequest(http.MethodGet, "/spots", nil)

	params := make(map[string]string, 1)
	params[endpoints.ParamSpotName] = "LesSurplombs"
	r = mux.SetURLVars(r, params)

	request, _ := DecodeRequestGetSpot(r)
	assert.Equal(t, request, getspot.GetSpotRequest{SpotName: "LesSurplombs"})
}

func TestDecodeRequestUpdateSpot(t *testing.T) {
	// nominal case
	body := `{"name":"LesSurplombs","routes":[{"name":"AlineLaMaline","level":"5a","points":5,"information":"Not so easy"}],"metadata":{"info":"Park at the canoe spot"}}`
	r := httptest.NewRequest(http.MethodPut, "/spots/spotName", strings.NewReader(body))

	params := make(map[string]string, 1)
	params[endpoints.ParamSpotName] = "LesSurplombs"
	r = mux.SetURLVars(r, params)

	request, _ := DecodeRequestUpdateSpot(r)
	assert.Equal(t, updatespot.UpdateSpotRequest{
		Name:     "LesSurplombs",
		Routes:   []spot.Route{{Name: "AlineLaMaline", Level: "5a", Points: 5, Information: "Not so easy"}},
		Metadata: map[string]interface{}{"info": "Park at the canoe spot"}}, request)

	// invalid case : a space is present in the name
	body = `{"name":"Les Surplombs","routes":[{"name":"AlineLaMaline","level":"5a","points":5,"information":"Not so easy"}],"metadata":{"info":"Park at the canoe spot"}}`
	r = httptest.NewRequest(http.MethodPut, "/spots", strings.NewReader(body))
	_, err := DecodeRequestUpdateSpot(r)
	assert.Contains(t, err.Message, "spot name field is empty")
	assert.Equal(t, http.StatusBadRequest, err.HTTPCode)

	// invalid case : no body
	r = httptest.NewRequest(http.MethodPost, "/spots/", nil)
	_, err = DecodeRequestUpdateSpot(r)
	assert.Contains(t, err.Message, "unable to decode updateSpotRequest")
	assert.Equal(t, http.StatusBadRequest, err.HTTPCode)
}

func TestDecodeRequestDeleteSpot(t *testing.T) {
	// nominal case
	r := httptest.NewRequest(http.MethodDelete, "/spots/spotName", nil)

	params := make(map[string]string, 1)
	params[endpoints.ParamSpotName] = "LesSurplombs"
	r = mux.SetURLVars(r, params)

	request, _ := DecodeRequestDeleteSpot(r)
	assert.Equal(t, deletespot.DeleteSpotRequest{SpotName: "LesSurplombs"}, request)
}
