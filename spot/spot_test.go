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
