# fxEcho Example Application

This example demonstrates a complete usage of the fxEcho module with Uber FX dependency injection, showcasing various features and best practices.

## Features Demonstrated

- **Dependency Injection**: Complete Uber FX setup with proper dependency management
- **Route Registration**: Individual routes and route groups using the builder pattern
- **Custom Middleware**: Request timing and request ID middleware
- **Service Layer**: Business logic separation with UserService
- **Handler Pattern**: Clean HTTP handlers with dependency injection
- **Configuration**: Integration with fxConfig module
- **Logging**: Structured logging with zap
- **Error Handling**: Proper HTTP error responses
- **Health Checks**: Built-in health endpoint

## Project Structure

```
example/
├── main.go              # Main application entry point
├── configs/
│   └── config.yaml      # Application configuration
└── README.md           # This file
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Welcome page with API information |
| GET | `/health` | Health check endpoint |
| GET | `/api/users` | List all users |
| GET | `/api/users/:id` | Get user by ID |
| POST | `/api/users` | Create a new user |

## Running the Example

### Prerequisites

- Go 1.24 or later
- The fxEcho module and its dependencies

### Quick Start

1. Navigate to the example directory:
   ```bash
   cd fxEcho/example
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. The server will start on `http://localhost:8080`

### Testing the API

#### Welcome Page
```bash
curl http://localhost:8080/
```

#### Health Check
```bash
curl http://localhost:8080/health
```

#### List Users
```bash
curl http://localhost:8080/api/users
```

#### Get User by ID
```bash
curl http://localhost:8080/api/users/1
```

#### Create New User
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

## Key Concepts Demonstrated

### 1. Dependency Injection with Uber FX

The application uses Uber FX for dependency injection, making it easy to:
- Manage dependencies
- Test components in isolation
- Configure the application
- Handle application lifecycle

### 2. Service Layer Pattern

The `UserService` demonstrates:
- Business logic separation
- Data management
- Logging integration
- Clean interfaces

### 3. Handler Pattern

HTTP handlers are:
- Injected with dependencies
- Focused on HTTP concerns
- Easy to test
- Well-structured

### 4. Route Registration

Routes are registered using the builder pattern:
- Type-safe route creation
- Group organization
- Clean separation of concerns

### 5. Custom Middleware

Custom middleware demonstrates:
- Request timing
- Request ID generation
- Logging integration
- Header manipulation

### 6. Configuration Management

Configuration is managed through:
- YAML files
- Environment variables
- Type-safe access
- Default values

## Configuration

The application uses the `fxConfig` module for configuration management. The `configs/config.yaml` file contains:

- Server settings (host, port, timeouts)
- Database configuration
- Middleware settings
- Application metadata

## Logging

The application uses structured logging with zap:
- Development-friendly output
- Request/response logging
- Error tracking
- Performance monitoring

## Error Handling

The application demonstrates proper error handling:
- HTTP status codes
- JSON error responses
- Logging of errors
- Graceful degradation

## Testing

To test the application:

1. Start the server
2. Use the provided curl commands
3. Check the logs for request processing
4. Verify responses match expected format

## Extending the Example

This example can be extended with:
- Database integration (using fxGorm)
- Authentication middleware
- Rate limiting
- API documentation
- Metrics collection
- More complex business logic

## Troubleshooting

### Common Issues

1. **Port already in use**: Change the port in `configs/config.yaml`
2. **Import errors**: Ensure all dependencies are properly installed
3. **Configuration not found**: Check that `configs/config.yaml` exists

### Debug Mode

To run in debug mode with more verbose logging, modify the logger configuration in `main.go`.

## Next Steps

After understanding this example, you can:
1. Add database integration
2. Implement authentication
3. Add more complex business logic
4. Create comprehensive tests
5. Deploy to production

This example serves as a foundation for building production-ready applications with the fxEcho module. 