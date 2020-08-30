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
	Name     string                 `json:"name,omitempty"`
	Routes   []Route                `json:"routes,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Route holds the information about a route
type Route struct {
	Name        string `json:"name,omitempty"`
	Level       string `json:"level,omitempty"`
	Points      int    `json:"points,omitempty"`
	Information string `json:"information,omitempty"`
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
	// spot name should not contain spaces
	if containSpace(d.Name) {
		return errors.New(fmt.Sprintf("spot name '%s' contains at least one space", d.Name))
	}

	// route names should not contain spaces
	for _, route := range d.Routes {
		if containSpace(route.Name) {
			return errors.New(fmt.Sprintf("route name '%s' contains at least one space", route.Name))
		}
	}
	return nil
}

// containSpace returns true if the given string contains a space
func containSpace(name string) bool {
	runeName := []rune(name)
	for _, r := range runeName {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}
