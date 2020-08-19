[![Build Status](https://travis-ci.com/doniacld/outdoorsight.svg?token=izzKU5X6FkS6FPGKshop&branch=init-structure)](https://travis-ci.com/doniacld/outdoorsighty)
[![codecov](https://codecov.io/gh/doniacld/outdoorsight/branch/master/graph/badge.svg?token=G8F353D8BW)](https://codecov.io/gh/doniacld/outdoorsight)


# Outdoorsight

Add your favorite climbing spots and the routes you achieved !! 

## Packaging

Build the docker image

    make docker_build

Run the docker image

    make docker_run

## API requests examples

### Add Spot

    curl -v -X PUT http://127.0.0.1:8080/spots -d '{
      "name": "luceram",
      "routes": [
        {
          "name": "Aline la maline",
          "level": "5a",
          "points": 5,
          "information": "ishouldbeanesayone"
        }
      ],
      "metadata": {
        "info": "go to the river"
      }
    }'

### Get Spot

    curl -v -X GET http://127.0.0.1:8080/spots/luceram

### Delete Spot

    curl -v -X DELETE http://127.0.0.1:8080/spots/luceram

### Update Spot

    curl -v -X POST http://127.0.0.1:8080/spots/luceram -d '{
      "routes": [
        {
          "name": "Bibi et Fricotin",
          "level": "5c",
          "points": 6,
        }
      ],
          "metadata": {
            "exposition": "Sun appears at 12pm during the summerF"
          }
        }'
    
## External resources

### Mongo related
https://godoc.org/go.mongodb.org/mongo-driver/mongo
