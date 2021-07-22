package db

// SpotDetails is the database structure that holds the information about a spot
// We do not want to expose the database format to the world, so we redefine it.
type SpotDetails struct {
	Name     string                 `bson:"name,omitempty"`
	Routes   []Route                `bson:"routes,omitempty"`
	Metadata map[string]interface{} `bson:"metadata,omitempty"`
}

// Route holds the information about a route
type Route struct {
	Name        string `bson:"name,omitempty"`
	Level       string `bson:"level,omitempty"`
	Points      int    `bson:"points,omitempty"`
	Information string `bson:"information,omitempty"`
}
