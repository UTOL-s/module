# Technology Stack & Build System

## Core Technologies

### Language & Runtime
- **Go 1.24.2**: Primary programming language
- **Uber FX**: Dependency injection framework for application lifecycle management
- **Echo v4**: High-performance HTTP web framework

### Database & ORM
- **GORM**: Object-relational mapping library
- **PostgreSQL**: Primary database (production)
- **MySQL, SQLite, SQL Server**: Additional supported databases
- Connection pooling and migration support

### Authentication & Security
- **SuperTokens**: Authentication and session management
- **JWT**: Token-based authentication
- **CORS**: Cross-origin resource sharing with configurable policies

### Configuration & Logging
- **Viper**: Configuration management with YAML support
- **Zap**: Structured logging library
- Environment variable support with expansion

### Development Tools
- **Air**: Hot reloading for development
- **Make**: Build automation and task management
- **Docker**: Containerization support
- **golangci-lint**: Code linting (optional)

## Build System

### Primary Commands

#### Development (Hot Reload)
```bash
# Quick start with hot reload (recommended)
./run-api.sh
make dev
air
```

#### Production Build
```bash
# Build optimized binary
./run-api.sh build
make api-rest-build
cd api_rest && go build -o api_rest_optimized main.go
```

#### Testing
```bash
# Run all tests
./run-api.sh test
make api-rest-test
go test ./...
```

#### Utilities
```bash
# Install development tools
make install-air
./run-api.sh install

# Code formatting and linting
make fmt
make lint

# Clean build artifacts
./run-api.sh clean
make clean
```

### Docker Support
```bash
# Build and run with Docker
make docker-build
make docker-run
docker-compose up -d
```

### Performance & Monitoring
```bash
# Performance profiling
make profile
make benchmark

# Load testing
make load-test

# Security scanning
make security
```

## Module Dependencies

### Core Dependencies
- `go.uber.org/fx`: Dependency injection
- `github.com/labstack/echo/v4`: HTTP framework
- `github.com/spf13/viper`: Configuration
- `go.uber.org/zap`: Logging
- `gorm.io/gorm`: ORM
- `github.com/supertokens/supertokens-golang`: Authentication

### Database Drivers
- `gorm.io/driver/postgres`: PostgreSQL
- `gorm.io/driver/mysql`: MySQL
- `gorm.io/driver/sqlite`: SQLite
- `gorm.io/driver/sqlserver`: SQL Server

### Development Dependencies
- `github.com/stretchr/testify`: Testing framework
- `github.com/air-verse/air`: Hot reloading
- `github.com/spf13/cobra`: CLI framework

## Environment Configuration

### Required Environment Variables
```bash
# Application
export APP_ENVIRONMENT=development
export CONFIG_FILE=api_rest/configs/config.yaml

# Database
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_USER=postgres
export DATABASE_PASSWORD=password
export DATABASE_DBNAME=utol_api

# SuperTokens
export SUPERTOKENS_CONNECTION_URI=your-supertokens-uri
export SUPERTOKENS_CONNECTION_API_KEY=your-api-key
```

## Performance Characteristics

### Expected Benchmarks
- **Request latency**: < 100ms for simple operations
- **Throughput**: > 1000 requests/second
- **Memory usage**: < 100MB baseline
- **Database connections**: Pooled (25 idle, 100 max)

### Optimization Features
- HTTP timeouts and connection pooling
- Response compression with configurable levels
- Request/response caching with TTL controls
- Concurrent request handling with worker pools
- Graceful shutdown with 30-second timeout