.PHONY: help build run test clean lint fmt

# Variables
BINARY_NAME=ade
GO=go
GOFLAGS=-v

help:
	@echo "Available commands:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Run the application"
	@echo "  make test        - Run tests"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make lint        - Run linter"
	@echo "  make fmt         - Format code"
	@echo "  make deps        - Download dependencies"

build:
	$(GO) build $(GOFLAGS) -o bin/$(BINARY_NAME) ./cmd/ade

run: build
	./bin/$(BINARY_NAME)

test:
	$(GO) test $(GOFLAGS) -race -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

clean:
	rm -rf bin/ coverage.out coverage.html

lint:
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

fmt:
	$(GO) fmt ./...
	goimports -w .

deps:
	$(GO) mod download
	$(GO) mod tidy

install-tools:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install golang.org/x/tools/cmd/goimports@latest

.DEFAULT_GOAL := help
