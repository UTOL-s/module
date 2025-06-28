# Unified Transport Operations League API (Optimized)

This is an optimized, production-ready integrated example that demonstrates how to use all the modules from the UTOL module collection together in a single REST API application with enhanced performance, security, and reliability.

## 🚀 Key Optimizations

### Performance Improvements
- **Context-aware operations** with proper timeout handling
- **Optimized database queries** with connection pooling and batch operations
- **Response compression** with configurable levels
- **Request/response caching** with TTL controls
- **Concurrent request handling** with worker pools
- **Bulk operations** for efficient data processing

### Security Enhancements
- **Rate limiting** with configurable thresholds
- **Security headers** (XSS protection, CSRF, etc.)
- **Request size limits** to prevent DoS attacks
- **Input validation** with comprehensive error handling
- **CORS configuration** with security best practices
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
│   └── config.yaml          # Comprehensive configuration with performance settings
├── handlers/
│   ├── auth.go              # Authentication handlers with session management
│   ├── health.go            # Health check handlers with database monitoring
│   └── user.go              # User management handlers with validation & search
├── middleware/
│   └── middleware.go        # Optimized middleware (security, rate limiting, compression)
├── models/
│   └── user.go              # User model with validation, hooks, and bulk operations
├── providers/
│   └── providers.go         # Enhanced dependency injection providers
├── routes/
│   └── routes.go            # Route definitions with authentication groups
├── go.mod                   # Module dependencies
├── main.go                  # Optimized main application with lifecycle management
├── example_test.go          # Integration tests
├── Makefile                 # Build and development commands
└── README.md               # This file
```

## ⚡ Performance Features

### Database Optimization
- **Connection pooling** with optimized settings
- **Query caching** with TTL controls
- **Batch operations** for bulk data processing
- **Indexed queries** for faster searches
- **Transaction management** with rollback support

### API Performance
- **Response compression** (GZIP) with configurable levels
- **Request caching** with intelligent invalidation
- **Concurrent processing** with worker pools
- **Pagination** with efficient offset/limit handling
- **Search optimization** with full-text search support

### Monitoring & Observability
- **Request timing** with detailed metrics
- **Performance logging** with slow query detection
- **Health monitoring** with dependency checks
- **Error tracking** with structured logging
- **Request tracing** with correlation IDs

## 🔒 Security Features

### Authentication & Authorization
- **SuperTokens integration** with session management
- **Role-based access control** (RBAC)
- **Session validation** with timeout handling
- **Protected routes** with middleware integration

### Request Security
- **Rate limiting** (100 requests/minute default)
- **Request size limits** (10MB default)
- **Input validation** with comprehensive rules
- **SQL injection protection** with parameterized queries
- **XSS protection** with security headers

### API Security
- **CORS configuration** with security headers
- **Content Security Policy** (CSP)
- **HTTPS enforcement** (configurable)
- **Request timeout** protection (30s default)

## 📋 Enhanced API Endpoints

### Health & Status
- `GET /` - Welcome message with API information
- `GET /health` - Health check with database status
- `GET /health/ready` - Readiness check for load balancers
- `GET /health/live` - Liveness check for container orchestration

### Authentication
- `GET /api/auth/status` - Authentication status
- `GET /api/auth/protected` - Protected route (requires SuperTokens session)
- `GET /api/auth/verify` - Session verification

### User Management
- `POST /api/users` - Create user with validation
- `GET /api/users` - List users with pagination and filtering
- `GET /api/users/search?q=query` - Search users with full-text search
- `GET /api/users/:id` - Get user by ID
- `PUT /api/users/:id` - Update user (requires authentication)
- `DELETE /api/users/:id` - Delete user (requires authentication)
- `GET /api/users/me` - Get current user (requires authentication)
- `PATCH /api/users/:id/role` - Update user role (requires authentication)
- `PATCH /api/users/:id/activate` - Activate user (requires authentication)
- `PATCH /api/users/:id/deactivate` - Deactivate user (requires authentication)

## 🛠 Getting Started

### Prerequisites

1. Go 1.24.2 or later
2. PostgreSQL, MySQL, SQLite, or SQL Server database
3. SuperTokens instance (optional, for authentication features)

### Installation

1. **Navigate to the example:**
   ```bash
   cd api_rest
   ```

2. **Update the configuration:**
   ```bash
   # Edit configs/config.yaml with your settings
   vim configs/config.yaml
   ```

3. **Run the application:**
   ```bash
   # Using Makefile
   make run
   
   # Or directly
   go run main.go
   ```

### Environment Variables

- `CONFIG_FILE`: Path to the configuration file (default: `configs/config.yaml`)
- `ENVIRONMENT`: Application environment (development, staging, production)
- `LOG_LEVEL`: Logging level (debug, info, warn, error)

### Database Setup

The application will automatically create the necessary database tables on startup. Make sure your database is running and accessible with the credentials specified in the configuration.

### SuperTokens Setup

To use authentication features:

1. Set up a SuperTokens instance
2. Update the SuperTokens configuration in `configs/config.yaml`
3. The application will automatically initialize SuperTokens on startup

## 🧪 Testing the API

### Health Check
```bash
curl http://localhost:8080/health
```

### Create a User
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "username": "testuser",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

### List Users with Pagination
```bash
curl "http://localhost:8080/api/users?offset=0&limit=10"
```

### Search Users
```bash
curl "http://localhost:8080/api/users/search?q=john&offset=0&limit=10"
```

### Filter Users
```bash
curl "http://localhost:8080/api/users?role=admin&active=true"
```

### Update User Role
```bash
curl -X PATCH http://localhost:8080/api/users/1/role \
  -H "Content-Type: application/json" \
  -d '{"role": "admin"}'
```

### Authentication Status
```bash
curl http://localhost:8080/api/auth/status
```

## 📊 Performance Monitoring

### Built-in Metrics
- Request timing and latency
- Database query performance
- Error rates and types
- Rate limiting statistics
- Memory and CPU usage

### Health Checks
- Database connectivity
- External service status
- Application readiness
- Resource availability

### Logging
- Structured JSON logging
- Request correlation IDs
- Performance metrics
- Error tracking with stack traces

## 🔧 Configuration

### Performance Settings
```yaml
performance:
  db:
    enable_query_cache: true
    cache_ttl: 300
    batch_size: 100
  response:
    enable_compression: true
    compression_level: 5
  concurrency:
    max_goroutines: 1000
    worker_pool_size: 10
```

### Security Settings
```yaml
security:
  rate_limit:
    enabled: true
    requests_per_minute: 100
    burst: 20
  validation:
    max_request_size: "10MB"
  cors:
    allowed_origins: ["*"]
    allowed_methods: ["GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"]
```

### Monitoring Settings
```yaml
monitoring:
  health:
    enabled: true
    endpoint: "/health"
  metrics:
    enabled: false
    endpoint: "/metrics"
  logging:
    level: "info"
    format: "json"
```

## 🚀 Production Deployment

### Docker Support
```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api_rest main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api_rest .
COPY --from=builder /app/configs ./configs
EXPOSE 8080
CMD ["./api_rest"]
```

### Environment Configuration
```bash
# Production environment
export ENVIRONMENT=production
export LOG_LEVEL=info
export CONFIG_FILE=/app/configs/config.yaml
```

### Health Checks
```bash
# Kubernetes liveness probe
curl -f http://localhost:8080/health/live

# Kubernetes readiness probe
curl -f http://localhost:8080/health/ready
```

## 🔍 Troubleshooting

### Common Issues

**Database Connection Issues:**
- Verify database credentials in `configs/config.yaml`
- Check database server connectivity
- Review connection pool settings

**Performance Issues:**
- Monitor request timing logs
- Check database query performance
- Review rate limiting settings
- Analyze memory usage

**Authentication Issues:**
- Verify SuperTokens configuration
- Check session timeout settings
- Review CORS configuration

### Debug Mode
```bash
# Enable debug logging
export LOG_LEVEL=debug
go run main.go
```

### Performance Profiling
```bash
# CPU profiling
go run -cpuprofile=cpu.prof main.go

# Memory profiling
go run -memprofile=mem.prof main.go
```

## 📈 Performance Benchmarks

### Expected Performance
- **Request latency**: < 100ms for simple operations
- **Database queries**: < 50ms for indexed queries
- **Concurrent requests**: 1000+ requests/second
- **Memory usage**: < 100MB for typical workloads
- **CPU usage**: < 20% under normal load

### Optimization Tips
1. **Database indexing**: Ensure proper indexes on frequently queried fields
2. **Connection pooling**: Tune pool settings based on load
3. **Caching**: Enable query caching for read-heavy workloads
4. **Compression**: Use GZIP compression for large responses
5. **Rate limiting**: Adjust limits based on your use case

## 🤝 Contributing

This optimized example serves as a reference implementation for integrating all UTOL modules with production-ready features. Feel free to extend it with additional optimizations or use it as a starting point for your own applications.

### Adding New Features
1. Follow the existing patterns for handlers, models, and routes
2. Include proper validation and error handling
3. Add comprehensive tests
4. Update documentation
5. Consider performance implications

### Performance Guidelines
1. Use context with timeouts for all operations
2. Implement proper error handling and logging
3. Use database transactions for multi-step operations
4. Add appropriate indexes for query optimization
5. Monitor and profile performance regularly

## Local Development Database (PostgreSQL)

You can run a local PostgreSQL database using Docker Compose:

```sh
cd api_rest
# Start the database in the background
docker compose up -d
```

This will start a PostgreSQL 15 container with:
- user: `postgres`
- password: `password`
- database: `utol_api`

The app will connect to this database by default (see `configs/config.yaml`).

To stop the database:
```sh
docker compose down
``` 