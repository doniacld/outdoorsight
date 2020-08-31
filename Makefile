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
# Targets
tidy:
	$(GOMOD) tidy
build:
		$(GOBUILD) -o $(BINARY_PATH) -v $(SOURCE_ENTRYPOINT)
run:
		$(GORUN) $(SOURCE_ENTRYPOINT)
test:
		$(GOTEST) -v ./... -cover
test_cover:
		$(GOTEST) -v ./... -coverprofile=coverage.txt -covermode=atomic
clean:
		$(GOCLEAN) $(SOURCE_ENTRYPOINT)
		rm -f $(BINARY_PATH)
clean_cache:
		$(GOCLEAN) --cache --testcache $(SOURCE_ENTRYPOINT)
docker_build:
		$(DOCKERBBUILD) -t outdoorsight .
docker_run:
		$(DOCKERRUN) --net=host -p 8080:8080 outdoorsight
docker_run_link:
		$(DOCKERRUN) -p 8080:8080 --name outdoorsight --link=mongoDB:database outdoorsight
docker_run_mongo:
		cd misc/mongo
		$(DOCKERRUN) -d --name mongoDB mongo
render_doc:
		redoc-cli bundle -o doc/api/index.html doc/api/src/paths.yml
run_outdoorsight:
		$(MAKE) build
		$(MAKE) render_doc
		$(MAKE) docker_build
		$(MAKE) docker_run_mongo
		$(MAKE) docker_run_link
stop_outdoorsight:
		$(DOCKERSTOP) outdoorsight mongoDB
		$(DOCKERRM) outdoorsight mongoDB
help:
	@echo "Compilation, image build, documentation build of Outdoorsight app"
	@echo "tidy                : Update dependencies (go mod tidy)"
	@echo "build               : Build Outdoorsight app"
	@echo "run                 : Run Outdoorsight app"
	@echo "test                : Launch all units tests"
	@echo "test_cover          : Launch all units tests with the cover of each package"
	@echo "docker_build        : Build the Outdoorsight docker image"
	@echo "docker_run          : Run the Outdoorsight docker image"
	@echo "docker_run_mongo    : Run the Mongo docker image"
	@echo "run_outdoorsight    : Build the app, build and run the app docker image and mongo! Just launch your curl."
	@echo "stop_outdoorsight   : Spot the docker containers and delete them"
	@echo "render_doc          : Render the yaml api documentation into a html"
	@echo "clean               : Remove temporary files"
	@echo "clean_cache         : Remove the cache"


