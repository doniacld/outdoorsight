#!/bin/bash

export ODS_ADDRESS=$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' outdoorsight)


# Add spot
echo "Adding spot to database"
curl -v -X POST http://$(ODS_ADDRESS):8080/spots -d@test/addspot_request.json

echo "------------------------------------------"

# Add a spot that already exists
echo "Adding a spot that already exists, an error is expected"
curl -v -X POST http://localhost:8080/spots -d@test/addspot_request.json

echo "------------------------------------------"

# Get spot
echo "Getting a created spot"
echo "curl -v http://localhost:8080/spots/LesSurplombs"
curl -v http://localhost:8080/spots/LesSurplombs

echo "------------------------------------------"

# Get a dummy spot
echo "Getting a spot that does not exist"
echo "curl -v http://localhost:8080/spots/dummySpot"
curl -v http://localhost:8080/spots/dummySpot

# Update a spot
echo "Updating existing spot details"
echo "curl -v -X PUT http://localhost:8080/spots/LesSurplombs -d@test/updatespot_request.json"
curl -v -X PUT http://localhost:8080/spots/LesSurplombs -d@test/updatespot_request.json

echo "------------------------------------------"

# Update a spot that does not exist
echo "Updating a non existing spot"
echo "curl -v -X PUT http://localhost:8080/spots/DummySpot -d@test/updatespot_request.json"
curl -v -X PUT http://localhost:8080/spots/DummySpot -d@test/updatespot_request.json

echo "------------------------------------------"

# Delete a spot
echo "Deleting a spot"
echo "curl -v -X DELETE http://localhost:8080/spots/LesSurplombs"
curl -v -X DELETE http://localhost:8080/spots/LesSurplombs

echo "------------------------------------------"

# Delete a spot that does not exist
echo "Deleting a spot that does not exist"
echo "curl -v -X DELETE http://localhost:8080/spots/dummySpot"
curl -v -X DELETE http://localhost:8080/spots/dummySpot
