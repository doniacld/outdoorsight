[![Build Status](https://travis-ci.com/doniacld/outdoorsight.svg?token=izzKU5X6FkS6FPGKshop&branch=init-structure)](https://travis-ci.com/doniacld/outdoorsighty)
[![codecov](https://codecov.io/gh/doniacld/outdoorsight/branch/master/graph/badge.svg?token=G8F353D8BW)](https://codecov.io/gh/doniacld/outdoorsight)


# Outdoorsight

Outdoorsight is a web service dedicated to rock climbers.  
You can add your favorite climbing spots, and their routes that you achieved or the next ones! 

It is a CRUD and RESTful API, communicating through http.

## Setup

Prerequisites

- Install docker

## Run the app

    make run_outdoorsight 

Please note if you already ran this command, launch:
    
    make stop_outdoorsight

## Install

Build the app docker image

    make docker_build

Run the app docker image

    make docker_run

Run mongoDB image

    make docker_run_mongo

### Mongo

Connect to mongo docker

    18:48 $ docker exec -it mongoDB  bash

Connect to mongo

    root@2673d1c1092f:/# mongo

See databases

    > show databases

Use `outdoorsight` database

    > use outdoorsight
    switched to db outdoorsight

Show collections 

    > show collections
    spots

See all elements in `spots` collection

    > db.spots.find()

If you want to check more infos, please refer to [mongoDB documentation](https://docs.mongodb.com/manual).

## Available Functionalities

| Endpoint    | Description                                       |
|-------------|---------------------------------------------------|
| AddSpot     | Add a spot to your list of spot                   |
| GetSpot     | Retrieve the given spot with its details          |
| UpdateSpot  | Update the given spot with the furnished details  |
| DeleteSpot  | Delete a spot from your list of spots             |
| GetAPIDoc   | Get the API documention in Redoc format           |

## Source code organization

* cmd : contains the main
* doc : contains the swagger API documentation in YAML
* internal/db : contains all db methods related
* internal/endpointdef : contains the meta to define an endpoint
* internal/endpoints : contains all the endpoints
* internal/routers : holds the mux router with all the routes
* internal/spot : holds the definition of a spot
* misc : holds docker images aside of the app (ex: mongo)

```bash
.
├── bin
├── CHANGELOG.md
├── cmd
│   └── main.go
├── doc
├── Dockerfile
├── internal
│   ├── db
│   ├── endpointdef
│   ├── endpoints
│   ├── routers
│   └── spot
├── Makefile
├── misc
│   └── mongo
└── README.md
```

## External resources that were useful

- [Mongo driver documentation](https://godoc.org/go.mongodb.org/mongo-driver/mongo)

