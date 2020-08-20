package spot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewRoute tests NewRoute method
func TestNewRoute(t *testing.T) {
	tt := map[string]struct {
		name        string
		level       string
		points      int
		information string
		expectedRes Route
	}{
		"nominal case": {name: "Aline la maline", level: "5a", points: 5, information: "maillon rapide",
			expectedRes: Route{
				Name:        "Aline la maline",
				Level:       "5a",
				Points:      5,
				Information: "maillon rapide",
			}},
	}

	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			route := NewRoute(tc.name, tc.level, tc.points, tc.information)
			assert.Equal(t, tc.expectedRes, route)
		})
	}
}

// TestValidate tests Validate method
func TestValidate(t *testing.T) {
	tt := map[string]struct {
		name        string
		expectedErr string
	}{
		"nominal case: no space":      {name: "laBagarre", expectedErr: ""},
		"error case: one space":       {name: "la Bagarre", expectedErr: "spot name 'la Bagarre' contains at least one space"},
		"error case: several spaces ": {name: "la Bagarre with spaces", expectedErr: "spot name 'la Bagarre with spaces' contains at least one space"},
	}

	for n, tc := range tt {
		spotDetails := Details{
			Name: tc.name,
		}
		t.Run(n, func(t *testing.T) {
			err := spotDetails.Validate()
			if tc.expectedErr != "" {
				assert.Equal(t, tc.expectedErr, err.Error())
			} else {
				assert.Equal(t, tc.name, spotDetails.Name)
			}
		})
	}
}
