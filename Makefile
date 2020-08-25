# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-moose

# Scripts
all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

compile:

	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/go-moose-freebsd-386 .
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/go-moose-darwin-386 .
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/go-moose-linux-386 .
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/go-moose-windows-386 .

	# 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/go-moose-freebsd-amd64 .
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/go-moose-darwin-amd64 .
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/go-moose-linux-amd64 .
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/go-moose-windows-amd64 .
