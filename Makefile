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
	$(GOTEST) -cover ./...
clean:
	$(GOCLEAN)
	rm -fr $(BUILD_DIR)
devrun: run sparun
run:
	$(GORUN) --race cmd/kaiserd/main.go
sparun:
	cd ./cmd/webapp/kaiser-spa && npm start

