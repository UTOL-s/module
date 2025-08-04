package fxgorm

import (
	"fmt"
	"strings"
	"time"

	fxconfig "github.com/UTOL-s/module/fxConfig"
	"gorm.io/gorm/logger"
)

// NewGormConfig creates a new GORM configuration from the main config
func NewGormConfig(config *fxconfig.Config) *GormConfig {
	return &GormConfig{
		Database: DatabaseConfig{
			// Primary URL-based configuration
			URL: config.Accessor.String("database.url"),

			// Legacy configuration for backward compatibility
			Type:      DatabaseType(config.Accessor.String("database.type")),
			Host:      config.Accessor.String("database.host"),
			Port:      config.Accessor.Int("database.port"),
			User:      config.Accessor.String("database.user"),
			Password:  config.Accessor.String("database.password"),
			DBName:    config.Accessor.String("database.dbname"),
			SSLMode:   config.Accessor.String("database.sslmode"),
			Charset:   config.Accessor.String("database.charset"),
			ParseTime: config.Accessor.Bool("database.parse_time"),
			Loc:       config.Accessor.String("database.loc"),
			File:      config.Accessor.String("database.file"),
		},
		Pool: PoolConfig{
			MaxIdleConns:    config.Accessor.Int("database.pool.max_idle_conns"),
			MaxOpenConns:    config.Accessor.Int("database.pool.max_open_conns"),
			ConnMaxLifetime: time.Duration(config.Accessor.Int("database.pool.conn_max_lifetime")) * time.Second,
			ConnMaxIdleTime: time.Duration(config.Accessor.Int("database.pool.conn_max_idle_time")) * time.Second,
		},
		Log: LogConfig{
			Level:                     logger.LogLevel(config.Accessor.Int("database.log.level")),
			SlowThreshold:             time.Duration(config.Accessor.Int("database.log.slow_threshold")) * time.Millisecond,
			Colorful:                  config.Accessor.Bool("database.log.colorful"),
			IgnoreRecordNotFoundError: config.Accessor.Bool("database.log.ignore_record_not_found_error"),
		},
		Debug: config.Accessor.Bool("database.debug"),
	}
}

// SetDefaults sets default values for unconfigured settings
func (gc *GormConfig) SetDefaults() {
	if gc.Log.Level == 0 {
		gc.Log.Level = logger.Info
	}
	if gc.Log.SlowThreshold == 0 {
		gc.Log.SlowThreshold = time.Second * 5
	}
	if gc.Pool.MaxIdleConns == 0 {
		gc.Pool.MaxIdleConns = 10
	}
	if gc.Pool.MaxOpenConns == 0 {
		gc.Pool.MaxOpenConns = 100
	}
	if gc.Pool.ConnMaxLifetime == 0 {
		gc.Pool.ConnMaxLifetime = time.Hour
	}
	if gc.Pool.ConnMaxIdleTime == 0 {
		gc.Pool.ConnMaxIdleTime = time.Minute * 10
	}
}

// Validate validates the configuration
func (gc *GormConfig) Validate() error {
	// If URL is provided, use URL-based validation
	if gc.Database.URL != "" {
		return gc.validateURL()
	}

	// Fall back to legacy validation
	return gc.validateLegacy()
}

// validateURL validates URL-based configuration
func (gc *GormConfig) validateURL() error {
	if gc.Database.URL == "" {
		return fmt.Errorf("database URL is required")
	}

	// Handle SQLite special case (sqlite:path format)
	if strings.HasPrefix(gc.Database.URL, "sqlite:") {
		return nil
	}

	// Basic URL format validation for other databases
	if !strings.Contains(gc.Database.URL, "://") {
		return fmt.Errorf("invalid database URL format: missing protocol")
	}

	return nil
}

// validateLegacy validates legacy configuration format
func (gc *GormConfig) validateLegacy() error {
	if gc.Database.Type == "" {
		gc.Database.Type = PostgreSQL // Default to PostgreSQL
	}

	switch gc.Database.Type {
	case PostgreSQL, MySQL, SQLServer:
		if gc.Database.Host == "" {
			return fmt.Errorf("host is required for %s database", gc.Database.Type)
		}
		if gc.Database.User == "" {
			return fmt.Errorf("user is required for %s database", gc.Database.Type)
		}
		if gc.Database.DBName == "" {
			return fmt.Errorf("dbname is required for %s database", gc.Database.Type)
		}
	case SQLite:
		if gc.Database.File == "" {
			return fmt.Errorf("file path is required for SQLite database")
		}
	default:
		return fmt.Errorf("unsupported database type: %s", gc.Database.Type)
	}

	return nil
}
