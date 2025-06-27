package fxgorm

import (
	"fmt"

	"database/sql"
)

// GetPoolStats returns current connection pool statistics
func (dm *DatabaseManager) GetPoolStats() (*sql.DBStats, error) {
	if dm.db == nil {
		return nil, fmt.Errorf("database not connected")
	}

	sqlDB, err := dm.db.DB()
	if err != nil {
		return nil, err
	}

	stats := sqlDB.Stats()
	return &stats, nil
}

// SetPoolConfig updates the connection pool configuration
func (dm *DatabaseManager) SetPoolConfig(config PoolConfig) error {
	if dm.db == nil {
		return fmt.Errorf("database not connected")
	}

	sqlDB, err := dm.db.DB()
	if err != nil {
		return err
	}

	// Update configuration
	dm.config.Pool = config

	// Apply new settings
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	return nil
}

// GetPoolConfig returns the current pool configuration
func (dm *DatabaseManager) GetPoolConfig() PoolConfig {
	return dm.config.Pool
}
