# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=toi

# arch and os


all: clean build
build:
		$(GOBUILD) -o dist/ ./cmd/toi
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run:
		$(GOBUILD) -o dist/ ./cmd/toi
		./$(BINARY_NAME)
deps:
		$(GOGET) ./...
