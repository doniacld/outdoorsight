package spot

// Spot holds all the method to handle a spot
type Spot interface {
	GetRoutes() []Route
}

// Details holds the information about a spot
type Details struct {
	Name     string                 `json:"name" bson:"name"`
	Routes   []Route                `json:"routes" bson:"routes"`
	Metadata map[string]interface{} `json:"metadata" bson:"metadata"`
}

// Route holds the information about a route
type Route struct {
	Name        string `json:"name" bson:"name"`
	Level       string `json:"level" bson:"level"`
	Points      int    `json:"points" bson:"points"`
	Information string `json:"information" bson:"information"`
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
