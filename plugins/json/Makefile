GOCMD=go
GOBUILD=$(GOCMD) build --buildmode=plugin
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

BUILD_DIR=.
BINARY_NAME=kaiser.json.plugin.so


all: clean test build
build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v
test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)
devrun:
	$(GORUN) main.go