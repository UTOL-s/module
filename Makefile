# Makefile for Unified Transport Operations League Module
# Root level Makefile for managing the api_rest example and development tasks

.PHONY: help install-air dev build test clean api-rest api-rest-dev api-rest-build api-rest-test api-rest-clean

# Default target
help:
	@echo "UTOL Module Development Commands"
	@echo "================================="
	@echo ""
	@echo "API REST Example Commands:"
	@echo "  api-rest        - Run the api_rest example"
	@echo "  api-rest-dev    - Run api_rest with Air (hot reload)"
	@echo "  api-rest-build  - Build the api_rest example"
	@echo "  api-rest-test   - Run tests for api_rest"
	@echo "  api-rest-clean  - Clean api_rest build artifacts"
	@echo ""
	@echo "Development Tools:"
	@echo "  install-air     - Install Air for hot reloading"
	@echo "  dev             - Run api_rest in development mode with Air"
	@echo "  build           - Build all examples"
	@echo "  test            - Run all tests"
	@echo "  clean           - Clean all build artifacts"
	@echo ""
	@echo "Individual Module Commands:"
	@echo "  fxecho-dev      - Run fxEcho example with Air"
	@echo "  fxgorm-test     - Run fxGorm tests"
	@echo "  fxsupertoken-dev - Run fxSupertoken example with Air"
	@echo ""

# Install Air for hot reloading
install-air:
	@echo "Installing Air for hot reloading..."
	@if command -v air >/dev/null 2>&1; then \
		echo "Air is already installed"; \
	else \
		go install github.com/air-verse/air@latest; \
		echo "Air installed successfully"; \
	fi

# Development mode with Air
dev: install-air
	@echo "Starting api_rest in development mode with Air..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "Air not found. Installing..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

# Build all examples
build:
	@echo "Building all examples..."
	@cd api_rest && make build
	@echo "All examples built successfully"

# Run all tests
test:
	@echo "Running all tests..."
	@cd api_rest && make test
	@echo "All tests completed"

# Clean all build artifacts
clean:
	@echo "Cleaning all build artifacts..."
	@cd api_rest && make clean
	@rm -rf tmp/
	@rm -f air.log
	@echo "All build artifacts cleaned"

# API REST Example Commands
api-rest:
	@echo "Running api_rest example..."
	@cd api_rest && go run main.go

api-rest-dev: install-air
	@echo "Running api_rest with Air (hot reload)..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "Air not found. Installing..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

api-rest-build:
	@echo "Building api_rest example..."
	@cd api_rest && make build

api-rest-test:
	@echo "Running api_rest tests..."
	@cd api_rest && make test

api-rest-clean:
	@echo "Cleaning api_rest build artifacts..."
	@cd api_rest && make clean

# Individual module commands
fxecho-dev: install-air
	@echo "Running fxEcho example with Air..."
	@if [ -d "fxEcho/example" ]; then \
		cd fxEcho/example && air; \
	else \
		echo "fxEcho example directory not found"; \
	fi

fxgorm-test:
	@echo "Running fxGorm tests..."
	@cd fxGorm && go test -v ./...

fxsupertoken-dev: install-air
	@echo "Running fxSupertoken example with Air..."
	@if [ -d "fxSupertoken/example" ]; then \
		cd fxSupertoken/example && air; \
	else \
		echo "fxSupertoken example directory not found"; \
	fi

# Quick start commands
start: api-rest-dev
	@echo "Quick start: api_rest with hot reload"

start-simple: api-rest
	@echo "Quick start: api_rest without hot reload"

# Development utilities
fmt:
	@echo "Formatting all Go code..."
	@find . -name "*.go" -exec go fmt {} \;

lint:
	@echo "Linting all Go code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download
	@cd api_rest && go mod tidy && go mod download

# Docker commands
docker-build:
	@echo "Building Docker image for api_rest..."
	@cd api_rest && make docker-build

docker-run:
	@echo "Running Docker container for api_rest..."
	@cd api_rest && make docker-run

# Performance and profiling
profile:
	@echo "Running performance profiling..."
	@cd api_rest && make profile-cpu

benchmark:
	@echo "Running benchmarks..."
	@cd api_rest && make benchmark

# Security
security:
	@echo "Running security scan..."
	@cd api_rest && make security-scan

# Load testing
load-test:
	@echo "Running load test..."
	@cd api_rest && make load-test 