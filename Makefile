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
DOKRCMD=docker
DOKRBBUILD=$(DOKRCMD) build
DOKRRUN=$(DOKRCMD) run

# Targets
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
		$(DOKRBBUILD) -t outdoorsight .
docker_run:
		$(DOKRRUN) -p 8080:8080 --net=host outdoorsight
docker_run_mongo:
		cd
		$(DOKRRUN) -p 8080:8080 --net=host outdoorsight
doc:
		redoc-cli bundle -o doc/api/index.html doc/api/src/paths.yml
help:
	@echo "Compilation, image build, documentation build of Outdoorsight app"
	@echo "build               : Build Outdoorsight app"
	@echo "run                 : Run Outdoorsight app"
	@echo "test                : Launch all units tests"
	@echo "test_cover          : Launch all units tests with the cover of each package"
	@echo "docker_build        : Build the Outdoorsight docker image"
	@echo "docker_run          : Run the Outdoorsight docker image"
	@echo "docker_run_mongo    : Run the Mongo docker image"
	@echo "doc                 : Render the yaml api documentation into a html"
	@echo "clean               : Remove temporary files"
	@echo "clean_cache         : Remove the cache"


