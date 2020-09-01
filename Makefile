# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Source parameters
SOURCE_ENTRYPOINT=./cmd/main.go

# Binary parameters
BINARY_NAME=cmd/main
BINARY_DESTINATION=./bin
BINARY_PATH=$(BINARY_DESTINATION)/$(BINARY_NAME)

# Docker parameters
DOCKERCMD=docker
DOCKERBBUILD=$(DOCKERCMD) build
DOCKERRUN=$(DOCKERCMD) run
DOCKERSTOP=$(DOCKERCMD) stop
DOCKERRM=$(DOCKERCMD) rm
DOCKERINSPECT=$(DOCKERCMD) inspect
DOCKERNETWORK=$(DOCKERCMD) network

# Targets
help:
	@echo "Compilation, image build, documentation build of Outdoorsight app"
	@echo "--------------     User     ------------------"
	@echo "run_outdoorsight    : Build the app, build and run the app docker image and mongo! Just launch your curl."
	@echo "stop_outdoorsight   : Spot the docker containers and delete them"
	@echo "test_endpoints      : Launch curls on endpoints (test/test_endpoints.sh)"
	@echo "--------------  Developer   ------------------"
	@echo "tidy                : Update dependencies (go mod tidy)"
	@echo "build               : Build Outdoorsight app"
	@echo "run                 : Run Outdoorsight app"
	@echo "unit_test           : Launch all units tests with the cover of each package"
	@echo "docker_build        : Build the Outdoorsight docker image"
	@echo "docker_run          : Run the Outdoorsight docker image"
	@echo "docker_run_mongo    : Run the Mongo docker image"
	@echo "render_doc          : Render the yaml api documentation into a html"
	@echo "clean               : Remove temporary files"
	@echo "clean_cache         : Remove the cache"
	@echo "get_address_mongo   : Retrieve the external IP address of mongoDB docker"
	@echo "get_address_ods     : Retrieve the external IP address of outdoorsight docker"

# DEVELOPER
tidy:
		$(GOMOD) tidy
build:
		$(GOBUILD) -o $(BINARY_PATH) -v $(SOURCE_ENTRYPOINT)
run:
		$(GORUN) $(SOURCE_ENTRYPOINT)
unit_test:
		$(GOTEST) -v ./... -coverprofile=coverage.txt -covermode=atomic
clean:
		$(GOCLEAN) $(SOURCE_ENTRYPOINT)
		rm -f $(BINARY_PATH)
clean_cache:
		$(GOCLEAN) --cache --testcache $(SOURCE_ENTRYPOINT)
docker_build: build
		$(DOCKERBBUILD) -t outdoorsight .
docker_run: export_address_mongo
		$(DOCKERRUN) -p 8080:8080 -e mongo_address=$(MONGO_ADDRESS) --network ods-network --name outdoorsight outdoorsight
docker_run_mongo:
		cd misc/mongo
		$(DOCKERRUN) -dit -p 27017:27017 --name mongoDB --network ods-network mongo
render_doc:
		redoc-cli bundle -o doc/api/index.html doc/api/src/paths.yml
create_network:
		$(DOCKERNETWORK) create ods-network
get_address_mongo: docker_run_mongo
		$(eval export MONGO_ADDRESS=$(shell docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mongoDB))
		@echo MONGO_ADDRESS=$(MONGO_ADDRESS)
get_address_ods: docker_run
		$(eval export ODS_ADDRESS=$(shell docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' outdoorsight))
		@echo ODS_ADDRESS=$(ODS_ADDRESS)
# USER
stop_outdoorsight:
		$(DOCKERSTOP) outdoorsight mongoDB
		$(DOCKERNETWORK) rm ods-network
		$(DOCKERRM) outdoorsight mongoDB
run_outdoorsight:
		$(MAKE) build
		$(MAKE) docker_build
		$(MAKE) create_network
		$(MAKE) docker_run_mongo
		$(MAKE) get_address_mongo
		$(MAKE) docker_run
test_endpoints:
		test/test_endpoints.sh
