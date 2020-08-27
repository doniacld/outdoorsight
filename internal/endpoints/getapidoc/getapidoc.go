package getapidoc

import (
	"github.com/pkg/errors"
	"io/ioutil"
)

const (
	docAPIFile = "./doc/api/index.html"
)

// GetAPIDoc returns the documentation on the API
func GetAPIDoc(_ GetAPIDocRequest) (GetAPIDocResponse, error) {
	file, err := ioutil.ReadFile(docAPIFile)
	if err != nil {
		panic(errors.Wrap(err, "unable to read doc api file"))
	}

	return file, nil
}
