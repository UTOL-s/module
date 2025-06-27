package fxgorm

import (
	"fmt"
	"time"

	"go.uber.org/fx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// configureConnectionPool configures the connection pool settings
func (gc *GormConfig) configureConnectionPool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Set default values if not configured
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

	sqlDB.SetMaxIdleConns(gc.Pool.MaxIdleConns)
	sqlDB.SetMaxOpenConns(gc.Pool.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(gc.Pool.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(gc.Pool.ConnMaxIdleTime)

	return nil
}

// NewGormDB creates a new GORM database instance with dynamic configuration
func NewGormDB(p Params) (*gorm.DB, error) {
	gormConfig := NewGormConfig(p.Config)

	// Set default values for unconfigured settings
	if gormConfig.Log.Level == 0 {
		gormConfig.Log.Level = logger.Info
	}
	if gormConfig.Log.SlowThreshold == 0 {
		gormConfig.Log.SlowThreshold = time.Second * 5
	}

	db, err := gormConfig.openDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := gormConfig.configureConnectionPool(db); err != nil {
		return nil, fmt.Errorf("failed to configure connection pool: %w", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// NewDatabaseManagerWithConfig creates a new database manager with the given configuration
func NewDatabaseManagerWithConfig(config *GormConfig) *DatabaseManager {
	return NewDatabaseManager(config)
}

// FxGorm provides the GORM module for dependency injection
var FxGorm = fx.Module(
	"fxgorm",
	fx.Provide(NewGormConfig),
	fx.Provide(NewGormDB),
	fx.Provide(NewDatabaseManagerWithConfig),
)
