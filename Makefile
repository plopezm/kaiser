GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

BUILD_DIR=dist
BINARY_NAME=kaiser


all: clean test build
build:
	mkdir $(BUILD_DIR)
	cp kaiser.config.json $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v
test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)
	rm -fr $(BUILD_DIR)
devrun:
	$(GORUN) main.go

