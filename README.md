# Unified Transport Operations League (UTOL) Module Collection

A comprehensive collection of Go modules for building production-ready applications with dependency injection, HTTP servers, database operations, and authentication.

## 🚀 Quick Start

### Running the API REST Example

The easiest way to run the integrated API REST example is using the provided tools:

#### Option 1: Using the Runner Script (Recommended)
```bash
# Run with hot reload (default)
./run-api.sh

# Run without hot reload
./run-api.sh run

# Build the application
./run-api.sh build

# Run tests
./run-api.sh test

# Show help
./run-api.sh help
```

#### Option 2: Using Make
```bash
# Run with hot reload
make dev

# Run without hot reload
make api-rest

# Build the application
make api-rest-build

# Run tests
make api-rest-test

# Show all available commands
make help
```

#### Option 3: Direct Commands
```bash
# Install Air for hot reloading
go install github.com/air-verse/air@latest

# Run with Air (hot reload)
air

# Run without Air
cd api_rest && go run main.go
```

## 📁 Module Structure

```
module/
├── fxConfig/           # Configuration management with YAML support
├── fxEcho/             # HTTP server with Echo framework and dependency injection
├── fxGorm/             # Database operations with GORM ORM
├── fxSupertoken/       # Authentication and session management with SuperTokens
├── api_rest/           # Integrated example combining all modules
├── .air.toml           # Air configuration for hot reloading
├── Makefile            # Root level Makefile for development tasks
├── run-api.sh          # Easy-to-use runner script
└── README.md           # This file
```

## 🔧 Development Tools

### Air - Hot Reloading
Air provides automatic reloading when files change, making development much faster.

**Installation:**
```bash
go install github.com/air-verse/air@latest
```

**Configuration:**
- `.air.toml` - Air configuration file (already configured for api_rest)
- Watches Go files, YAML configs, and templates
- Automatically rebuilds and restarts on changes

**Usage:**
```bash
# Run with Air
air

# Or use the runner script
./run-api.sh
```

### Make Commands
The root Makefile provides convenient commands for development:

```bash
# Development
make dev              # Run api_rest with Air
make api-rest         # Run api_rest without Air
make api-rest-dev     # Run api_rest with Air
make api-rest-build   # Build api_rest
make api-rest-test    # Test api_rest
make api-rest-clean   # Clean api_rest

# Utilities
make install-air      # Install Air
make fmt              # Format all Go code
make lint             # Lint all Go code
make deps             # Install dependencies

# Performance
make profile          # Run performance profiling
make benchmark        # Run benchmarks
make load-test        # Run load tests

# Security
make security         # Run security scan

# Docker
make docker-build     # Build Docker image
make docker-run       # Run Docker container
```

### Runner Script
The `run-api.sh` script provides a simple interface for common tasks:

```bash
# Quick commands
./run-api.sh          # Run with hot reload (default)
./run-api.sh dev      # Run with hot reload
./run-api.sh run      # Run without hot reload
./run-api.sh build    # Build application
./run-api.sh test     # Run tests
./run-api.sh clean    # Clean build artifacts
./run-api.sh install  # Install Air
./run-api.sh help     # Show help
```

## 📋 Available Modules

### fxConfig
Configuration management with YAML support, environment variables, and hot reloading.

**Features:**
- YAML configuration files
- Environment variable support
- Hot reloading
- Type-safe configuration access

### fxEcho
HTTP server with Echo framework and Uber FX dependency injection.

**Features:**
- Echo web framework integration
- Dependency injection with Uber FX
- Middleware support
- Route management
- Graceful shutdown

### fxGorm
Database operations with GORM ORM and connection pooling.

**Features:**
- GORM ORM integration
- Connection pooling
- Multiple database drivers (PostgreSQL, MySQL, SQLite, SQL Server)
- Migration support
- Query optimization

### fxSupertoken
Authentication and session management with SuperTokens.

**Features:**
- SuperTokens integration
- Session management
- Authentication middleware
- Role-based access control

### api_rest (Integrated Example)
A complete REST API example that demonstrates how to use all modules together.

**Features:**
- Complete user management API
- Authentication with SuperTokens
- Database operations with GORM
- Configuration management
- Health checks
- Performance optimizations
- Security features

## 🛠 Development Workflow

### 1. Start Development
```bash
# Quick start with hot reload
./run-api.sh

# Or using make
make dev
```

### 2. Make Changes
Edit any Go files, YAML configs, or templates. Air will automatically detect changes and reload the application.

### 3. Test Your Changes
```bash
# Run tests
./run-api.sh test

# Or
make api-rest-test
```

### 4. Build for Production
```bash
# Build optimized binary
./run-api.sh build

# Or
make api-rest-build
```

## 🔍 Configuration

### Air Configuration (.air.toml)
The Air configuration is set up to:
- Watch the `api_rest` directory
- Monitor Go files, YAML configs, and templates
- Exclude temporary and build directories
- Provide colored output for different operations

### Environment Variables
```bash
# Configuration file location
export CONFIG_FILE=api_rest/configs/config.yaml

# Environment
export ENVIRONMENT=development

# Log level
export LOG_LEVEL=debug
```

## 🚀 Production Deployment

### Docker
```bash
# Build Docker image
make docker-build

# Run Docker container
make docker-run
```

### Manual Deployment
```bash
# Build for production
make api-rest-build

# Run the binary
./api_rest/api_rest_optimized
```

## 📊 Monitoring and Debugging

### Health Checks
```bash
# Health check
curl http://localhost:8080/health

# Readiness check
curl http://localhost:8080/health/ready

# Liveness check
curl http://localhost:8080/health/live
```

### Performance Profiling
```bash
# CPU profiling
make profile

# Memory profiling
cd api_rest && make profile-memory

# Benchmarks
make benchmark
```

### Load Testing
```bash
# Load test
make load-test

# Stress test
cd api_rest && make stress-test
```

## 🤝 Contributing

1. **Fork the repository**
2. **Create a feature branch**
3. **Make your changes**
4. **Test your changes**
   ```bash
   make api-rest-test
   ```
5. **Submit a pull request**

### Development Guidelines
- Follow Go best practices
- Add tests for new features
- Update documentation
- Use the provided development tools
- Run linting and formatting before committing

## 📚 Documentation

- [API REST Example README](api_rest/README.md) - Detailed documentation for the integrated example
- [Individual Module READMEs](fxConfig/README.md) - Documentation for each module
- [Configuration Guide](api_rest/configs/config.yaml) - Configuration options and examples

## 🆘 Troubleshooting

### Common Issues

**Air not found:**
```bash
# Install Air
./run-api.sh install

# Or manually
go install github.com/air-verse/air@latest
```

**Build errors:**
```bash
# Clean and rebuild
./run-api.sh clean
./run-api.sh build
```

**Port already in use:**
```bash
# Check what's using the port
lsof -i :8080

# Kill the process or change the port in config
```

**Database connection issues:**
- Check database credentials in `api_rest/configs/config.yaml`
- Ensure database server is running
- Verify network connectivity

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 


