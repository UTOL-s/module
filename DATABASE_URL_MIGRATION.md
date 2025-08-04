# Database Configuration Migration to URL Format

## Overview

The UTOL module collection has been updated to support a simplified database configuration using a single URL format instead of separate connection parameters. This change makes configuration more portable and follows industry standards.

## Changes Made

### 1. Configuration Files Updated

#### `configs/config.yaml.example`
- **Before**: Separate database connection parameters
- **After**: Single `url` field with examples for all database types

#### `api_rest/configs/config.yaml`
- **Before**: Individual database connection fields
- **After**: Single `DATABASE_URL` environment variable

#### `api_rest/env.example`
- **Before**: Multiple database environment variables
- **After**: Single `DATABASE_URL` with legacy variables commented out

### 2. Code Changes

#### `fxGorm/types.go`
- Added `URL` field to `DatabaseConfig` struct
- Maintained all legacy fields for backward compatibility

#### `fxGorm/config.go`
- Updated `NewGormConfig()` to read URL configuration
- Added `validateURL()` function for URL-based validation
- Enhanced `Validate()` to handle both URL and legacy formats

#### `fxGorm/database.go`
- Added `parseURL()` function to parse database URLs
- Updated `buildDSN()` to handle URL-based configuration
- Enhanced `openDatabase()` for URL support
- Special handling for SQLite URLs (`sqlite:path` format)

#### `fxGorm/README.md`
- Updated documentation with URL-based examples
- Added migration guide
- Maintained legacy documentation for backward compatibility

### 3. Testing

#### `fxGorm/url_test.go` (New)
- Comprehensive tests for URL parsing
- Tests for all supported database types
- Validation tests for URL formats
- DSN building tests with URLs

## Supported URL Formats

### PostgreSQL
```
postgres://user:password@localhost:5432/dbname?sslmode=disable
```

### MySQL
```
mysql://user:password@localhost:3306/dbname?charset=utf8mb4&parseTime=true&loc=Local
```

### SQLite
```
sqlite:./data/database.db
```

### SQL Server
```
sqlserver://user:password@localhost:1433?database=dbname
```

## Migration Guide

### For New Projects
Use the URL-based configuration:

```yaml
database:
  url: "postgres://postgres:password@localhost:5432/mydb?sslmode=disable"
```

### For Existing Projects
No changes required! The legacy configuration format is still fully supported:

```yaml
database:
  type: "postgres"
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "mydb"
  sslmode: "disable"
```

### Environment Variables

#### New (Recommended)
```bash
export DATABASE_URL=postgres://postgres:password@localhost:5432/mydb?sslmode=disable
```

#### Legacy (Still Supported)
```bash
export DATABASE_TYPE=postgres
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_USER=postgres
export DATABASE_PASSWORD=password
export DATABASE_DBNAME=mydb
export DATABASE_SSLMODE=disable
```

## Backward Compatibility

- All existing configurations continue to work without changes
- Legacy environment variables are still supported
- URL-based configuration takes precedence when both formats are present
- No breaking changes to the API

## Benefits

1. **Simplified Configuration**: Single URL instead of multiple parameters
2. **Industry Standard**: Follows common database URL conventions
3. **Portability**: Easy to copy/paste connection strings
4. **Environment Friendly**: Single environment variable for database connection
5. **Tool Compatibility**: Works with standard database tools and libraries

## Testing

All changes have been thoroughly tested:
- ✅ URL parsing for all database types
- ✅ Backward compatibility with legacy configuration
- ✅ DSN building with both URL and legacy formats
- ✅ Validation for both configuration types
- ✅ Integration tests with the API REST example

## Files Modified

- `configs/config.yaml.example`
- `api_rest/configs/config.yaml`
- `api_rest/env.example`
- `fxGorm/types.go`
- `fxGorm/config.go`
- `fxGorm/database.go`
- `fxGorm/README.md`
- `fxGorm/example_test.go`

## Files Added

- `fxGorm/url_test.go`
- `DATABASE_URL_MIGRATION.md` (this file)

## Next Steps

1. Update your projects to use the new URL-based configuration when convenient
2. Update documentation and examples in your applications
3. Consider using the URL format for new deployments and configurations
4. The legacy format will continue to be supported for the foreseeable future


<!--  -->