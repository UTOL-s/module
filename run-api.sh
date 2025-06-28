#!/bin/bash

# UTOL API REST Runner Script
# This script makes it easy to run the api_rest example with different options

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_header() {
    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}================================${NC}"
}

# Function to check if Air is installed
check_air() {
    if command -v air >/dev/null 2>&1; then
        return 0
    else
        return 1
    fi
}

# Function to install Air
install_air() {
    print_status "Installing Air for hot reloading..."
    go install github.com/air-verse/air@latest
    print_status "Air installed successfully!"
}

# Function to show help
show_help() {
    print_header "UTOL API REST Runner"
    echo ""
    echo "Usage: $0 [OPTION]"
    echo ""
    echo "Options:"
    echo "  dev, --dev, -d     Run with Air (hot reload) - default"
    echo "  run, --run, -r     Run without hot reload"
    echo "  build, --build, -b Build the application"
    echo "  test, --test, -t   Run tests"
    echo "  clean, --clean, -c Clean build artifacts"
    echo "  install, --install Install Air"
    echo "  help, --help, -h   Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0                 # Run with hot reload (default)"
    echo "  $0 dev             # Run with hot reload"
    echo "  $0 run             # Run without hot reload"
    echo "  $0 build           # Build the application"
    echo "  $0 test            # Run tests"
    echo ""
}

# Function to run with Air (hot reload)
run_with_air() {
    if ! check_air; then
        print_warning "Air not found. Installing..."
        install_air
    fi
    
    print_status "Starting api_rest with Air (hot reload)..."
    print_status "The application will automatically reload when you make changes"
    print_status "Press Ctrl+C to stop"
    echo ""
    
    # Check if .air.toml exists
    if [ ! -f ".air.toml" ]; then
        print_error ".air.toml configuration file not found!"
        exit 1
    fi
    
    air
}

# Function to run without hot reload
run_simple() {
    print_status "Starting api_rest without hot reload..."
    print_status "Press Ctrl+C to stop"
    echo ""
    
    cd api_rest
    go run main.go serve
}

# Function to build the application
build_app() {
    print_status "Building api_rest application..."
    cd api_rest
    make build
    print_status "Build completed successfully!"
}

# Function to run tests
run_tests() {
    print_status "Running api_rest tests..."
    cd api_rest
    make test
    print_status "Tests completed!"
}

# Function to clean build artifacts
clean_build() {
    print_status "Cleaning build artifacts..."
    cd api_rest
    make clean
    rm -rf ../tmp/
    rm -f ../air.log
    print_status "Clean completed!"
}

# Function to install Air
install_air_only() {
    install_air
}

# Main script logic
main() {
    # Default to dev mode if no arguments provided
    if [ $# -eq 0 ]; then
        run_with_air
        exit 0
    fi
    
    case "$1" in
        dev|--dev|-d)
            run_with_air
            ;;
        run|--run|-r)
            run_simple
            ;;
        build|--build|-b)
            build_app
            ;;
        test|--test|-t)
            run_tests
            ;;
        clean|--clean|-c)
            clean_build
            ;;
        install|--install)
            install_air_only
            ;;
        help|--help|-h)
            show_help
            ;;
        *)
            print_error "Unknown option: $1"
            echo ""
            show_help
            exit 1
            ;;
    esac
}

# Run main function with all arguments
main "$@" 