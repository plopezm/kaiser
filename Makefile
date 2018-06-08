GOCMD=go
GOINSTALL=$(GOCDM)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

BUILD_DIR=dist
BINARY_NAME=kaiser


all: clean build
build:
	mkdir $(BUILD_DIR)
	cp kaiser.config.json ./dist
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/kaiserd 
test:
	$(GOTEST) -v ./**/*_test.go
clean:
	$(GOCLEAN)
	rm -fr $(BUILD_DIR)
devrun:
	$(GORUN) cmd/kaiserd/main.go

