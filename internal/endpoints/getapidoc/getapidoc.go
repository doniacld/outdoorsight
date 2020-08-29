package getapidoc

import (
	"github.com/doniacld/outdoorsight/internal/errors"
	"io/ioutil"
	"net/http"
)

const (
	// docAPIFile is copied during the build of the docker image
	docAPIFile = "./doc/api/index.html"
)

// GetAPIDoc returns the documentation on the API
func GetAPIDoc(_ GetAPIDocRequest) (GetAPIDocResponse, *errors.OsError) {
	file, err := ioutil.ReadFile(docAPIFile)
	if err != nil {
		return GetAPIDocResponse{}, errors.NewFromError(http.StatusInternalServerError, err, "unable to read doc api file")
	}
	return file, nil
}
