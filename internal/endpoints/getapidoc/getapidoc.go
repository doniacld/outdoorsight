package getapidoc

import (
	"io/ioutil"
	"net/http"

	"github.com/doniacld/outdoorsight/internal/endpointdef"

	"github.com/pkg/errors"
)

const (
	docAPIFile = "./doc/api/index.html"
)

// GetAPIDoc returns all the details on a given spot
func GetAPIDoc(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile(docAPIFile)
	if err != nil {
		panic(errors.Wrap(err, "unable to read doc api file"))
	}

	// set the response parameters
	w.Header().Set(endpointdef.ContentType, endpointdef.MimeTypeHTML)
	w.WriteHeader(GetAPIDocMeta.SuccessCode())
	w.Write(file)
}
