# Unified Transport Operations League API (Optimized)

This is an optimized, production-ready integrated example that demonstrates how to use all the modules from the UTOL module collection together in a single REST API application with enhanced performance, security, and reliability.

## 🚀 Key Features

### Simplified CORS Implementation
- **Default CORS**: Automatically applied through the FX middleware system
- **Development-friendly**: Allows all origins (`*`) for easy development
- **Production-ready**: Includes `ProductionCORSHandler` for specific domain restrictions
- **Configurable**: Can be easily customized for different environments

### Performance Improvements
- **Context-aware operations** with proper timeout handling
- **Optimized database queries** with connection pooling and batch operations
- **Response compression** with configurable levels
- **Request/response caching** with TTL controls
- **Concurrent request handling** with worker pools
- **Simplified middleware** for better performance

### Security Enhancements
- **Simplified CORS** with sensible defaults
- **Rate limiting** with configurable thresholds
- **Security headers** (XSS protection, CSRF, etc.)
- **Request size limits** to prevent DoS attacks
- **Input validation** with comprehensive error handling
- **Request timeout** protection

### Reliability Features
- **Graceful shutdown** with signal handling
- **Health checks** with database connectivity testing
- **Error recovery** with detailed logging
- **Request tracing** with unique request IDs
- **Structured logging** with performance metrics
- **Database migrations** with automatic rollback

## 📁 Project Structure

```
api_rest/
├── configs/
│   └── config.yaml          # Simplified configuration (CORS now automatic)
├── internal/
│   ├── handler/
│   │   ├── health.go        # Health check handlers
│   │   └── user.go          # User management handlers
│   ├── router.go            # Clean router (CORS handled by FX)
│   ├── register.go          # Dependencies with CORS middleware
│   └── bootstrap.go         # Application bootstrap
├── main.go                  # Application entry point
├── go.mod                   # Go module definition
└── README.md               # This file
```

## 🔧 CORS Configuration

### Default CORS (Development)
The API automatically applies a default CORS configuration that:
- Allows all origins (`*`) for development flexibility
- Includes standard HTTP methods (GET, POST, PUT, PATCH, DELETE, OPTIONS)
- Supports common headers (Authorization, Content-Type, etc.)
- Sets a 24-hour cache for preflight requests

### Production CORS
For production environments, you can use the `ProductionCORSHandler`:

```go
// In register.go, replace the default CORS with:
fx.Annotate(
    func() echo.MiddlewareFunc {
        return fxSupertoken.ProductionCORSHandler([]string{
            "https://yourdomain.com",
            "https://www.yourdomain.com",
        })
    },
    fx.ResultTags(`group:"middlewares"`),
),
```

## 🚀 Quick Start

### Prerequisites
- Go 1.24 or later
- PostgreSQL database
- Environment variables configured (see `env.example`)

### Installation
```bash
# Clone the repository
git clone <repository-url>
cd api_rest

# Copy environment variables
cp env.example .env
# Edit .env with your database credentials

# Install dependencies
go mod download

# Run the application
go run main.go serve
```

### Testing CORS
```bash
# Test CORS preflight request
curl -i -H "Origin: http://localhost:3000" \
     -H "Access-Control-Request-Method: POST" \
     -H "Access-Control-Request-Headers: Content-Type" \
     -X OPTIONS http://localhost:8080/api/v1/users

# Test regular API request
curl -i http://localhost:8080/api/health
```

## 🔍 API Endpoints

### Health Checks
- `GET /api/health` - Basic health check with database status
- `GET /api/health/ready` - Readiness check for Kubernetes

### User Management
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users` - List all users
- `GET /api/v1/users/search` - Search users
- `GET /api/v1/users/:id` - Get user by ID (requires authentication)
- `PUT /api/v1/users/:id` - Update user (requires authentication)
- `DELETE /api/v1/users/:id` - Delete user (requires authentication)

### Authentication
- Authentication endpoints are handled by SuperTokens
- Default path: `/api/auth/*`

## 📋 Configuration

The application uses a simplified configuration approach:

```yaml
# config.yaml
app:
  name: "Unified Transport Operations League API"
  version: "1.0.0"
  environment: "development"

server:
  host: "0.0.0.0"
  port: "8080"
  read_timeout: 30
  write_timeout: 30

database:
  type: "postgres"
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "utol_api"

# CORS is now handled automatically by the FX middleware system
# No manual CORS configuration needed in config.yaml
```

## 🏗️ Architecture

The application follows a clean architecture pattern with dependency injection:

- **FX Framework**: Manages dependency injection and lifecycle
- **Echo Framework**: HTTP server and routing
- **GORM**: Database ORM with PostgreSQL
- **SuperTokens**: Authentication and session management
- **Simplified CORS**: Automatic CORS handling through middleware

## 🔧 Development

### Running in Development Mode
```bash
# Set environment
export APP_ENVIRONMENT=development

# Run with hot reload (if you have air installed)
air

# Or run directly
go run main.go serve
```

### Testing
```bash
# Run tests
go test ./...

# Run with coverage
go test -cover ./...
```

## 🚀 Production Deployment

### Environment Variables
```bash
# Production environment
export APP_ENVIRONMENT=production
export SERVER_HOST=0.0.0.0
export SERVER_PORT=8080

# Database
export DB_HOST=your-db-host
export DB_PORT=5432
export DB_USER=your-db-user
export DB_PASSWORD=your-db-password
export DB_NAME=your-db-name

# SuperTokens
export SUPERTOKENS_CONNECTION_URI=your-supertokens-uri
export SUPERTOKENS_CONNECTION_API_KEY=your-api-key
```

### Docker Support
```bash
# Build Docker image
docker build -t utol-api .

# Run with Docker Compose
docker-compose up -d
```

## 🔍 Troubleshooting

### Common Issues

**CORS Issues:**
- The API now uses automatic CORS handling
- For development: All origins are allowed by default
- For production: Update the middleware registration in `register.go`

**Database Connection Issues:**
- Verify database credentials in `.env`
- Check database server connectivity
- Review connection pool settings in `config.yaml`

**Performance Issues:**
- Monitor request timing logs
- Check database query performance
- Review rate limiting settings

### Debug Mode
```bash
# Enable debug logging
export LOG_LEVEL=debug
go run main.go serve
```

## 📈 Performance Benchmarks

### Expected Performance
- **Request latency**: < 100ms for simple operations
- **Throughput**: > 1000 requests/second
- **Memory usage**: < 100MB baseline
- **Database connections**: Pooled with 25 idle, 100 max

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🔗 Related Modules

- **fxConfig**: Configuration management
- **fxEcho**: Echo framework integration
- **fxGorm**: Database ORM integration
- **fxSupertoken**: Authentication and CORS middleware 