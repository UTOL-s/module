# Makefile for Unified Transport Operations League API (Optimized)

.PHONY: build run test clean deps help profile benchmark optimize

# Default target
all: build

# Build the application
build:
	@echo "Building optimized UTOL API..."
	go build -ldflags="-s -w" -o api_rest_optimized main.go
	@echo "Build complete!"

# Run the application
run:
	@echo "Starting optimized UTOL API..."
	go run main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmarks
benchmark:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f api_rest_optimized
	rm -f *.prof
	rm -f coverage.out
	rm -f coverage.html
	go clean

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Run with specific config file
run-config:
	@echo "Starting UTOL API with custom config..."
	CONFIG_FILE=configs/config.yaml go run main.go

# Database setup (example)
db-setup:
	@echo "Setting up database..."
	@echo "Please ensure your database is running and configured in configs/config.yaml"

# Performance profiling
profile-cpu:
	@echo "Running with CPU profiling..."
	go run -cpuprofile=cpu.prof main.go

profile-memory:
	@echo "Running with memory profiling..."
	go run -memprofile=mem.prof main.go

# Analyze profiles
analyze-cpu:
	@echo "Analyzing CPU profile..."
	go tool pprof cpu.prof

analyze-memory:
	@echo "Analyzing memory profile..."
	go tool pprof mem.prof

# Performance optimization
optimize:
	@echo "Running performance optimizations..."
	@echo "1. Building with optimizations..."
	go build -ldflags="-s -w" -o api_rest_optimized main.go
	@echo "2. Running tests..."
	go test ./...
	@echo "3. Running benchmarks..."
	go test -bench=. -benchmem ./...
	@echo "Optimization complete!"

# Security scan
security-scan:
	@echo "Running security scan..."
	gosec ./...

# Dependency vulnerability check
vuln-check:
	@echo "Checking for vulnerabilities..."
	go list -json -deps ./... | nancy sleuth

# Load testing (requires hey tool)
load-test:
	@echo "Running load test..."
	@if command -v hey >/dev/null 2>&1; then \
		hey -n 1000 -c 10 http://localhost:8080/health; \
	else \
		echo "Install hey tool: go install github.com/rakyll/hey@latest"; \
	fi

# Stress testing
stress-test:
	@echo "Running stress test..."
	@if command -v hey >/dev/null 2>&1; then \
		hey -n 10000 -c 100 http://localhost:8080/health; \
	else \
		echo "Install hey tool: go install github.com/rakyll/hey@latest"; \
	fi

# Development mode
dev:
	@echo "Starting in development mode..."
	LOG_LEVEL=debug go run main.go

# Production build
prod-build:
	@echo "Building for production..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o api_rest_prod main.go

# Docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t utol-api:latest .

# Docker run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 utol-api:latest

# Docker Compose up
db-up:
	docker compose up -d

# Docker Compose down
db-down:
	docker compose down

# Show help
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  benchmark      - Run benchmarks"
	@echo "  clean          - Clean build artifacts"
	@echo "  deps           - Install dependencies"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  run-config     - Run with specific config file"
	@echo "  db-setup       - Database setup instructions"
	@echo "  profile-cpu    - Run with CPU profiling"
	@echo "  profile-memory - Run with memory profiling"
	@echo "  analyze-cpu    - Analyze CPU profile"
	@echo "  analyze-memory - Analyze memory profile"
	@echo "  optimize       - Run performance optimizations"
	@echo "  security-scan  - Run security scan"
	@echo "  vuln-check     - Check for vulnerabilities"
	@echo "  load-test      - Run load test"
	@echo "  stress-test    - Run stress test"
	@echo "  dev            - Start in development mode"
	@echo "  prod-build     - Build for production"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  db-up          - Docker Compose up"
	@echo "  db-down        - Docker Compose down"
	@echo "  help           - Show this help message" 