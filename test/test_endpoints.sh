#!/bin/bash

# Add spot
echo "Add spot to database"
curl -v -X POST http://localhost:8080/spots -d@test/addspot_request.json

echo "------------------------------------------"

# Add a spot that already exists
echo "Add spot to already exist, an error is expected"
curl -v -X POST http://localhost:8080/spots -d@test/addspot_request.json

echo "------------------------------------------"

# Get spot
echo "Get a created spot"
echo "curl -v http://localhost:8080/spots/LesSurplombs"
curl -v http://localhost:8080/spots/LesSurplombs

echo "------------------------------------------"

# Get a dummy spot
echo "Get a spot that does not exist"
echo "curl -v http://localhost:8080/spots/dummySpot"
curl -v http://localhost:8080/spots/dummySpot

# Update a spot
echo "Update existing spot details"
echo "curl -v -X PUT http://localhost:8080/spots/LesSurplombs -d@test/updatespot_request.json"
curl -v -X PUT http://localhost:8080/spots/LesSurplombs -d@test/updatespot_request.json

echo "------------------------------------------"

# Update a spot that does not exist
echo "Update a non existing spot"
echo "curl -v -X PUT http://localhost:8080/spots/DummySpot -d@test/updatespot_request.json"
curl -v -X PUT http://localhost:8080/spots/DummySpot -d@test/updatespot_request.json

echo "------------------------------------------"

# Delete a spot
echo "Delete a spot"
echo "curl -v -X DELETE http://localhost:8080/spots/LesSurplombs"
curl -v -X DELETE http://localhost:8080/spots/LesSurplombs

echo "------------------------------------------"

# Delete a spot that does not exist
echo "Delete a spot that does not exist"
echo "curl -v -X DELETE http://localhost:8080/spots/LesSurplombs"
curl -v -X DELETE http://localhost:8080/spots/LesSurplombs
