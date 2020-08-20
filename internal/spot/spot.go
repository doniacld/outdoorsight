package spot

import (
	"fmt"
	"unicode"

	"github.com/pkg/errors"
)

// Spot holds all the method to handle a spot
type Spot interface {
	GetRoutes() []Route
	Validate() error
}

// Details holds the information about a spot
type Details struct {
	Name     string                 `json:"name"`
	Routes   []Route                `json:"routes"`
	Metadata map[string]interface{} `json:"metadata"`
}

// Route holds the information about a route
type Route struct {
	Name        string `json:"name"`
	Level       string `json:"level"`
	Points      int    `json:"points"`
	Information string `json:"information"`
}

// NewRoute returns a new structure filled with given parameters
func NewRoute(name string, level string, points int, information string) Route {
	return Route{
		Name:        name,
		Level:       level,
		Points:      points,
		Information: information,
	}
}

// Validate checks that a spot name does not have a space
func (d Details) Validate() error {
	runeName := []rune(d.Name)
	for _, r := range runeName {
		if unicode.IsSpace(r) {
			return errors.New(fmt.Sprintf("spot name '%s' contains at least one space", d.Name))
		}
	}
	return nil
}
