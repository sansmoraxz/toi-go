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

build.linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o dist/toi-linux ./cmd/toi
build.win:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o dist/toi-windows.exe ./cmd/toi

clean:
		$(GOCLEAN)
		rm -f dist/$(BINARY_NAME)
run:
		$(GOBUILD) -o dist/ ./cmd/toi
		./dist/$(BINARY_NAME)
deps:
		$(GOGET) ./...
