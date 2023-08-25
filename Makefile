# Go parameters
GOCMD=go
GODEPCMD=godep
GOBUILD=$(GOCMD) build
GODEPBUILD=$(GODEPCMD) $(GOBUILD)
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix
BIN_APP= groco-app

all: clean build run

clean:
	$(GOCLEAN)
	rm -f $(BIN_APP)

build: clean
	$(GODEPBUILD) -o $(BIN_APP) ./app

run: build
	cd app; ../$(BIN_APP)

test:
	$(GOTEST) ./app