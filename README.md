[![Build Status](https://travis-ci.com/doniacld/outdoorsight.svg?token=izzKU5X6FkS6FPGKshop&branch=init-structure)](https://travis-ci.com/doniacld/outdoorsighty)
[![codecov](https://codecov.io/gh/doniacld/outdoorsight/branch/master/graph/badge.svg?token=G8F353D8BW)](https://codecov.io/gh/doniacld/outdoorsight)


# Outdoorsight

Outdoorsight is a **web service dedicated to rock climbers.**  
You can add your favourite climbing spots and their routes, whether you achieved them or plan on doing so!  
It is a CRUD and RESTful API, communicating through http.

## Available Functionalities

| Endpoint    | Description                                       |
|-------------|---------------------------------------------------|
| AddSpot     | Add a spot to your list of spot                   |
| GetSpot     | Retrieve the given spot with its details          |
| UpdateSpot  | Update the given spot with the furnished details  |
| DeleteSpot  | Delete a spot from your list of spots             |
| GetAPIDoc   | Get the API documentation in Redoc format         |

## Setup

Prerequisites

- [go 1.14+](https://golang.org/doc/install)
- Install [Docker](https://docs.docker.com/get-docker/)

In case you want to rework the documentation:
* [Redoc](https://github.com/Redocly/redoc)

## Users

**Run the app**

    make run_outdoorsight 

It will execute the following actions :
* build the app
* build the docker image of the app `outdoorsight`
* create a custom network into docker names `ods-network`
* run mongo docker image with the container name `mongoDB`
* run the docker image `outdoorsight` with the external mongo IP dynamically created

Please note if you already ran this command, launch:

    make stop_outdoorsight

> It stops and deletes the docker images `outdoorsight` and `mongoDB` but also remove the created network `ods-network`.

**Test endpoints**  

To test the endpoints, you can launch the following command:

    make test_endpoints
    
It will execute a list of curls on the existing endpoints.
* Add spot
* Get spot
* Get spot on a non existing spot
* Update an existing spot
* Update a non existing spot
* Delete spot
* Delete spot on a non existing spot

**Documentation**  

You can retrieve the API documentation in Redoc format through this URL:

    http://$ODS_ADDRESS:8080/doc/api

It can be modified and regenerated thanks to:

    make render_doc

> It is regenerated in the principal make run of the app `make run_outdoorsight`.

## Developers

Create a custom network

    make create_network
    
It will create a custom network `ods-network`, you can check it by executing: `docker network ls`

```bash
$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
c138929b4188        bridge              bridge              local
f9338c39767d        host                host                local
5d605d641126        local               bridge              local
9a1685a5e7f5        none                null                local
7fc7df208412        ods-network         bridge              local
```

Run mongoDB image

    make docker_run_mongo

Build the app docker image

    make docker_build

Run the app docker image
> It takes the external IP of mongo docker image as env variable.
> Please note that you can retrieve this IP through `make getaddress_mongo`
    
    make docker_run

To confirm that your services can be launched together, please run `docker network inspect ods-network`.  
You should obtain something like below with `mongoDB` and `outdoorsight` containers.


```bash
$ docker network inspect ods-network
[
    {
        "Name": "ods-network",
        "Id": "7fc7df208412cf3b8f729ea0fcc729cc1fe6c12d4ed0d8355ac77daaec93f09d",
        "Created": "2020-09-01T11:41:21.093803458+02:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "192.168.64.0/20",
                    "Gateway": "192.168.64.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "1357b14a60d9d6926ac72e4da688ab2866b297aa77708287a9aba5e75765a3e4": {
                "Name": "outdoorsight",
                "EndpointID": "8dfec01d03f09b88195dde38920a3ecc973eee91af3afe793b63938c1e5a9478",
                "MacAddress": "02:42:c0:a8:40:03",
                "IPv4Address": "192.168.64.3/20",
                "IPv6Address": ""
            },
            "4a90ad7ecc41b3f22402e8c2f0b0e93cb810314b4b516174615a556492163330": {
                "Name": "mongoDB",
                "EndpointID": "71a9288b2fed837e49dc6c941aec42191e4daa94e5557ff24cf3d9e56611e94c",
                "MacAddress": "02:42:c0:a8:40:02",
                "IPv4Address": "192.168.64.2/20",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {}
    }
]
```

### Mongo

Connect to mongo docker

    $ docker exec -it mongoDB bash

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

* `cmd` : contains the main
* `doc` : contains the swagger API documentation in YAML and the generated html
* `internal/db` : contains all db methods related
* `internal/endpointdef` : contains the meta to define an endpoint
* `internal/endpoints` : contains all the endpoints
* `internal/errors` : holds the internal error library
* `internal/routers` : holds the mux router with all the routes
* `internal/spot` : holds the definition of a spot
* `misc` : holds docker images aside of the app (here and for the moment: mongo)

## What's next ?

**Mongo**  
* Add a copy of session each time an endpoint is called instead of creating a new connection
* Add indexes on spots to have a more efficient on data access
* Create a new collection `routes` to store all the routes and develop CRUD endpoints on `route` resource

**Endpoints**  
* Implement GetSpots and AddSpots endpoints to retrieve export and imports a list of spots
* Implement CRUD API on `routes` resource

**Packaging**  
* Rework Dockerfile to build a multistage file:
    1. Copy all the source code and build the main
    2. Copy the yaml doc and render into a HTML file
    3. Take the bin and HTML and run the image

**UI**
* A reflexion on UI using VueJS is in progress (currently learning VueJS with [this tutorial](https://www.grafikart.fr/formations/vuejs))

**General**
* Have a persistent database 
* I would like to have a user system and create a website on which everybody could register, save their spots and rock climb achievements!

## Some stuff that I learned from working on this project

* How to create a network bridge with Docker and make two containers talk to each other
* Mongo driver methods are less extended than the shell Mongo methods!
* Variables in Makefile are set only in the Makefile shell...
* How to separate business from database
* Set and manipulate HTTP responses

## Some external resources

- [Go documentation](https://godoc.org)
- [Mongo driver documentation](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
- [Networking with standalone containers](https://docs.docker.com/network/network-tutorial-standalone)
- [Medium article about clean architecture on Golang](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)
