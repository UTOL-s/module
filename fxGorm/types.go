package fxgorm

import (
	"time"

	fxconfig "github.com/UTOL-s/module/fxConfig"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseType represents supported database types
type DatabaseType string

const (
	PostgreSQL DatabaseType = "postgres"
	MySQL      DatabaseType = "mysql"
	SQLite     DatabaseType = "sqlite"
	SQLServer  DatabaseType = "sqlserver"
)

// DatabaseConfig holds database-specific configuration
type DatabaseConfig struct {
	Type      DatabaseType `mapstructure:"type"`
	Host      string       `mapstructure:"host"`
	Port      int          `mapstructure:"port"`
	User      string       `mapstructure:"user"`
	Password  string       `mapstructure:"password"`
	DBName    string       `mapstructure:"dbname"`
	SSLMode   string       `mapstructure:"sslmode"`
	Charset   string       `mapstructure:"charset"`
	ParseTime bool         `mapstructure:"parse_time"`
	Loc       string       `mapstructure:"loc"`
	File      string       `mapstructure:"file"` // For SQLite
}

// PoolConfig holds connection pool configuration
type PoolConfig struct {
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level                     logger.LogLevel `mapstructure:"level"`
	SlowThreshold             time.Duration   `mapstructure:"slow_threshold"`
	Colorful                  bool            `mapstructure:"colorful"`
	IgnoreRecordNotFoundError bool            `mapstructure:"ignore_record_not_found_error"`
}

// GormConfig holds the complete GORM configuration
type GormConfig struct {
	Database DatabaseConfig `mapstructure:"database"`
	Pool     PoolConfig     `mapstructure:"pool"`
	Log      LogConfig      `mapstructure:"log"`
	Debug    bool           `mapstructure:"debug"`
}

// Params holds the dependency injection parameters
type Params struct {
	Config *fxconfig.Config
}

// DatabaseManager handles database operations
type DatabaseManager struct {
	config *GormConfig
	db     *gorm.DB
}

// NewDatabaseManager creates a new database manager
func NewDatabaseManager(config *GormConfig) *DatabaseManager {
	return &DatabaseManager{
		config: config,
	}
}
