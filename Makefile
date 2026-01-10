.PHONY: help test test-integration test-all lint fmt build clean install-tools coverage

help:
	@echo "Available targets:"
	@echo "  test              - Run unit tests"
	@echo "  test-integration  - Run integration tests"
	@echo "  test-all          - Run all tests"
	@echo "  lint              - Run linters"
	@echo "  fmt               - Format code"
	@echo "  build             - Build the project"
	@echo "  clean             - Clean build artifacts"
	@echo "  install-tools     - Install development tools"
	@echo "  coverage          - Generate test coverage report"

test:
	go test -v -race -cover ./...

test-integration:
	go test -v -race -tags=integration ./tests/integration/...

test-all:
	go test -v -race -cover -tags=integration ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -s -w .
	goimports -w .

build:
	go build ./...

clean:
	go clean ./...
	rm -f coverage.out coverage.html

install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "Tools installed successfully"

coverage:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
