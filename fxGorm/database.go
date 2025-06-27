package fxgorm

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// buildDSN constructs the database connection string based on the database type
func (gc *GormConfig) buildDSN() (string, error) {
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

	switch gc.Database.Type {
	case PostgreSQL:
		dialector = postgres.Open(dsn)
	case MySQL:
		dialector = mysql.Open(dsn)
	case SQLite:
		dialector = sqlite.Open(dsn)
	case SQLServer:
		dialector = sqlserver.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", gc.Database.Type)
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
