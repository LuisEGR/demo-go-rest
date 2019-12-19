# App Data
BINARY_NAME=prueba

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
GOGET=$(GOCMD) get
APP_DIR=src/main.go
GO_DIR=/go/
BINARY_PATH=bin
BINARY_UNIX=$(BINARY_NAME)_unix


# export GOPATH=/Users/luisenriquegonzalezrodriguez/Documents/Gitlab/cp/services/go-template/lib/
# export GOPATH=$(CURDIR)$(GO_DIR)

all: test build

build: 
	$(GOBUILD) -v -o $(BINARY_PATH)/$(BINARY_NAME) $(APP_DIR)

init: 
	export GOPATH=$(CURDIR)
	dep init

test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/markbates/goth
		$(GOGET) github.com/markbates/pop


# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v    