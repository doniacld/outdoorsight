[![Build Status](https://travis-ci.com/doniacld/outdoorsight.svg?token=izzKU5X6FkS6FPGKshop&branch=init-structure)](https://travis-ci.com/doniacld/outdoorsighty)
[![codecov](https://codecov.io/gh/doniacld/outdoorsight/branch/master/graph/badge.svg?token=G8F353D8BW)](https://codecov.io/gh/doniacld/outdoorsight)


# Outdoorsight

Outdoorsight is a web service dedicated to rock climbers.  
You can add your favourite climbing spots and their routes, whether you achieved them or plan on doing so!
It is a CRUD and RESTful API, communicating through http.

## Setup

Prerequisites

- [go 1.14+](https://golang.org/doc/install)
- Install [Docker](https://docs.docker.com/get-docker/)

## Users

**Run the app**

    make run_outdoorsight 

Please note if you already ran this command, launch:
    
    make stop_outdoorsight


## Admins

Build the app docker image

    make docker_build

Run the app docker image

    make docker_run

Run mongoDB image

    make docker_run_mongo

### Mongo

Connect to mongo docker

    $ docker exec -it mongoDB  bash

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
| GetAPIDoc   | Get the API documentation in Redoc format           |


So you can retrieve the API documentation in Redoc format through this URL:

    http://localhost:8080/doc/api

It can be modified and regenerated thanks to:

    make render_doc

> It is regenerated in the principal make run of the app `make run_outdoorsight`.

## Source code organization

The Outdoorsight project tree

```bash
.
├── CHANGELOG.md
├── cmd
│   └── main.go
├── doc
├── Dockerfile
├── internal
│   ├── db
│   ├── endpointdef
│   ├── endpoints
│   ├── errors
│   ├── routers
│   └── spot
├── Makefile
├── misc
│   └── mongo
└── README.md
```

* cmd : contains the main
* doc : contains the swagger API documentation in YAML
* internal/db : contains all db methods related
* internal/endpointdef : contains the meta to define an endpoint
* internal/endpoints : contains all the endpoints
* internal/errors : holds the internal error library
* internal/routers : holds the mux router with all the routes
* internal/spot : holds the definition of a spot
* misc : holds docker images aside of the app (ex: mongo)

## What's next ?

**Mongo**
* Add mocks on DB to mock it
* Add a copy of session each time an endpoint is called instead of creating a new connection
* Add indexes on spots to have a more efficient on data access
* Create a new collection `routes` to store all the routes and develop CRUD endpoints on `route` resource

**Endpoints**
* Implement getSpots endpoint to retrieve all the spots
* Implement CRUD API on `routes` resource

**UI**
* A reflexion on UI using VueJS is in progress (currently learning VueJS)

**General**
* Have a persistent database 
* I would like to have a user system and create a website on which everybody could register, save their spots and rock climb achievements!

## External resources that were useful

- [Mongo driver documentation](https://godoc.org/go.mongodb.org/mongo-driver/mongo)

