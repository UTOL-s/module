package fxgorm

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// parseURL parses a database URL and populates the legacy fields for compatibility
func (gc *GormConfig) parseURL() error {
	if gc.Database.URL == "" {
		return nil // No URL to parse
	}

	// Handle SQLite special case (sqlite:path format)
	if strings.HasPrefix(gc.Database.URL, "sqlite:") {
		gc.Database.Type = SQLite
		gc.Database.File = strings.TrimPrefix(gc.Database.URL, "sqlite:")
		return nil
	}

	u, err := url.Parse(gc.Database.URL)
	if err != nil {
		return fmt.Errorf("invalid database URL: %w", err)
	}

	// Set database type based on scheme
	switch u.Scheme {
	case "postgres", "postgresql":
		gc.Database.Type = PostgreSQL
	case "mysql":
		gc.Database.Type = MySQL
	case "sqlite":
		gc.Database.Type = SQLite
		// For standard sqlite:// URLs
		gc.Database.File = strings.TrimPrefix(u.Path, "/")
		return nil
	case "sqlserver":
		gc.Database.Type = SQLServer
	default:
		return fmt.Errorf("unsupported database scheme: %s", u.Scheme)
	}

	// Parse connection details for other databases
	gc.Database.Host = u.Hostname()
	if u.Port() != "" {
		if port, err := strconv.Atoi(u.Port()); err == nil {
			gc.Database.Port = port
		}
	}

	if u.User != nil {
		gc.Database.User = u.User.Username()
		if password, ok := u.User.Password(); ok {
			gc.Database.Password = password
		}
	}

	// Database name is the path without leading slash
	if u.Path != "" && u.Path != "/" {
		gc.Database.DBName = strings.TrimPrefix(u.Path, "/")
	}

	// Parse query parameters
	query := u.Query()
	for key, values := range query {
		if len(values) > 0 {
			switch key {
			case "sslmode":
				gc.Database.SSLMode = values[0]
			case "charset":
				gc.Database.Charset = values[0]
			case "parseTime":
				gc.Database.ParseTime = values[0] == "true"
			case "loc":
				gc.Database.Loc = values[0]
			case "database": // SQL Server uses 'database' query param
				if gc.Database.Type == SQLServer {
					gc.Database.DBName = values[0]
				}
			}
		}
	}

	return nil
}

// buildDSN constructs the database connection string based on the database type
func (gc *GormConfig) buildDSN() (string, error) {
	// If URL is provided, use it directly (after parsing for type detection)
	if gc.Database.URL != "" {
		if err := gc.parseURL(); err != nil {
			return "", err
		}

		// For most databases, we can use the URL directly
		// But we need to handle some special cases
		switch gc.Database.Type {
		case SQLite:
			// SQLite URLs need special handling
			return gc.Database.File, nil
		default:
			// For other databases, we can use the URL as-is
			return gc.Database.URL, nil
		}
	}

	// Fall back to legacy DSN building
	switch gc.Database.Type {
	case PostgreSQL:
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			gc.Database.Host, gc.Database.User, gc.Database.Password,
			gc.Database.DBName, gc.Database.Port, gc.Database.SSLMode), nil
	case MySQL:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", gc.Database.User, gc.Database.Password,
			gc.Database.Host, gc.Database.Port, gc.Database.DBName)
		if gc.Database.Charset != "" {
			dsn += "?charset=" + gc.Database.Charset
		}
		if gc.Database.ParseTime {
			if gc.Database.Charset != "" {
				dsn += "&parseTime=true"
			} else {
				dsn += "?parseTime=true"
			}
		}
		if gc.Database.Loc != "" {
			if strings.Contains(dsn, "?") {
				dsn += "&loc=" + gc.Database.Loc
			} else {
				dsn += "?loc=" + gc.Database.Loc
			}
		}
		return dsn, nil
	case SQLite:
		if gc.Database.File == "" {
			return "", fmt.Errorf("SQLite database file path is required")
		}
		return gc.Database.File, nil
	case SQLServer:
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			gc.Database.User, gc.Database.Password,
			gc.Database.Host, gc.Database.Port, gc.Database.DBName), nil
	default:
		return "", fmt.Errorf("unsupported database type: %s", gc.Database.Type)
	}
}

// openDatabase opens a database connection based on the configuration
func (gc *GormConfig) openDatabase() (*gorm.DB, error) {
	dsn, err := gc.buildDSN()
	if err != nil {
		return nil, err
	}

	var dialector gorm.Dialector

	// Determine database type (either from URL parsing or legacy config)
	dbType := gc.Database.Type
	if dbType == "" && gc.Database.URL != "" {
		// Parse URL to determine type if not already set
		if err := gc.parseURL(); err != nil {
			return nil, err
		}
		dbType = gc.Database.Type
	}

	switch dbType {
	case PostgreSQL:
		dialector = postgres.Open(dsn)
	case MySQL:
		dialector = mysql.Open(dsn)
	case SQLite:
		dialector = sqlite.Open(dsn)
	case SQLServer:
		dialector = sqlserver.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	return gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			log.New(log.Writer(), "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             gc.Log.SlowThreshold,
				LogLevel:                  gc.Log.Level,
				IgnoreRecordNotFoundError: gc.Log.IgnoreRecordNotFoundError,
				Colorful:                  gc.Log.Colorful,
			},
		),
		DryRun: gc.Debug,
	})
}

// Connect establishes a database connection with validation
func (dm *DatabaseManager) Connect() error {
	// Validate configuration
	if err := dm.config.Validate(); err != nil {
		return fmt.Errorf("configuration validation failed: %w", err)
	}

	// Set defaults
	dm.config.SetDefaults()

	// Open database connection
	db, err := dm.config.openDatabase()
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	if err := dm.config.configureConnectionPool(db); err != nil {
		return fmt.Errorf("failed to configure connection pool: %w", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	dm.db = db
	return nil
}

// GetDB returns the database instance
func (dm *DatabaseManager) GetDB() *gorm.DB {
	return dm.db
}

// Close closes the database connection
func (dm *DatabaseManager) Close() error {
	if dm.db != nil {
		sqlDB, err := dm.db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
