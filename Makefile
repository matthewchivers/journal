# Project name
PROJECT_NAME := $(shell basename $(PWD))

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary name
BINARY_NAME=$(PROJECT_NAME)_bin

# Architectures
LINUX_AMD64=$(BINARY_NAME)_linux_amd64
WINDOWS_AMD64=$(BINARY_NAME)_windows_amd64.exe
DARWIN_AMD64=$(BINARY_NAME)_darwin_amd64
DARWIN_ARM64=$(BINARY_NAME)_darwin_arm64

all: build

build: test
	$(GOBUILD) -o $(BINARY_NAME)

buildall: build-linux-amd64 build-windows-amd64 build-darwin-amd64 build-darwin-arm64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_AMD64)

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(WINDOWS_AMD64)

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(DARWIN_AMD64)

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(DARWIN_ARM64)

lint:
	golangci-lint run

test: vet lint
	$(GOTEST) -v ./...

vet:
	$(GOCMD) vet ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(LINUX_AMD64)
	rm -f $(WINDOWS_AMD64)
	rm -f $(DARWIN_AMD64)
	rm -f $(DARWIN_ARM64)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
