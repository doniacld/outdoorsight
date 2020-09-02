#!/bin/bash

export ODS_ADDRESS=$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' outdoorsight)

# Add spot
echo -e "\e[32mAdding spot to database"
echo -e "\e[32mcurl -v -X POST http://$ODS_ADDRESS:8080/spots -d@test/addspot_request.json "
echo -e "\e[0m------------------------------------------"
curl -v -X POST http://$ODS_ADDRESS:8080/spots -d@test/addspot_request.json

echo -e "\e[0m------------------------------------------"

# Add a spot that already exists
echo -e "\e[32mAdding a spot that already exists, an error is expected"
echo -e "\e[32mcurl -v -X POST http://$ODS_ADDRESS:8080/spots -d@test/addspot_request.json"
echo -e "\e[0m------------------------------------------"
curl -v -X POST http://$ODS_ADDRESS:8080/spots -d@test/addspot_request.json

echo -e "\e[0m------------------------------------------"

# Get spot
echo -e "\e[32mGetting a created spot"
echo -e "\e[32mcurl -v http://$ODS_ADDRESS:8080/spots/LesSurplombs"
echo -e "\e[0m------------------------------------------"
curl -v http://$ODS_ADDRESS:8080/spots/LesSurplombs

echo -e "\e[0m------------------------------------------"

# Get a dummy spot
echo -e "\e[32mGetting a spot that does not exist"
echo -e "\e[32mcurl -v http://$ODS_ADDRESS:8080/spots/dummySpot"
echo -e "\e[0m------------------------------------------"
curl -v http://$ODS_ADDRESS:8080/spots/dummySpot

# Update a spot
echo -e "\e[32mUpdating existing spot details"
echo -e "\e[32mcurl -v -X PUT http://$ODS_ADDRESS:8080/spots/LesSurplombs -d@test/updatespot_request.json"
echo -e "\e[0m------------------------------------------"
curl -v -X PUT http://$ODS_ADDRESS:8080/spots/LesSurplombs -d@test/updatespot_request.json

echo -e "\e[0m------------------------------------------"

# Update a spot that does not exist
echo -e "\e[32mUpdating a non existing spot"
echo -e "\e[32mcurl -v -X PUT http://$ODS_ADDRESS:8080/spots/DummySpot -d@test/updatespot_request.json"
echo -e "\e[0m------------------------------------------"
curl -v -X PUT http://$ODS_ADDRESS:8080/spots/DummySpot -d@test/updatespot_request.json

echo -e "\e[0m------------------------------------------"

# Delete a spot
echo -e "\e[32mDeleting a spot"
echo -e "\e[32mcurl -v -X DELETE http://$ODS_ADDRESS:8080/spots/LesSurplombs"
echo -e "\e[0m------------------------------------------"
curl -v -X DELETE http://$ODS_ADDRESS:8080/spots/LesSurplombs

echo -e "\e[0m------------------------------------------"

# Delete a spot that does not exist
echo -e "\e[32mDeleting a spot that does not exist"
echo -e "\e[32mcurl -v -X DELETE http://$ODS_ADDRESS:8080/spots/dummySpot"
echo -e "\e[0m------------------------------------------"
curl -v -X DELETE http://$ODS_ADDRESS:8080/spots/dummySpot
