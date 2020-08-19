package db

// SpotDetails is the database structure that holds the information about a spot
// We do not want to expose the database format to the world.
type SpotDetails struct {
	Name     string                 `bson:"name"`
	Routes   []Route                `bson:"routes"`
	Metadata map[string]interface{} `bson:"metadata"`
}

// Route holds the information about a route
type Route struct {
	Name        string `bson:"name"`
	Level       string `bson:"level"`
	Points      int    `bson:"points"`
	Information string `bson:"information"`
}
