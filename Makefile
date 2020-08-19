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
		docker build -t outdoorsight .
docker_run:
		docker run -p 8080:8080 --net=host outdoorsight

doc:
		redoc-cli bundle -o doc/api/index.html doc/api/src/paths.yml
