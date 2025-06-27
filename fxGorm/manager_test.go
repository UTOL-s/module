package fxgorm

import (
	"testing"
)

// TestDatabaseManager tests the DatabaseManager creation and basic functionality
func TestDatabaseManager(t *testing.T) {
	config := &GormConfig{
		Database: DatabaseConfig{
			Type:     PostgreSQL,
			Host:     "localhost",
			Port:     5432,
			User:     "test",
			Password: "test",
			DBName:   "testdb",
			SSLMode:  "disable",
		},
	}

	manager := NewDatabaseManager(config)
	if manager == nil {
		t.Error("DatabaseManager should not be nil")
		return
	}

	if manager.config != config {
		t.Error("DatabaseManager should have the correct config")
	}
}

// TestDatabaseManagerWithConfig tests the dependency injection version
func TestDatabaseManagerWithConfig(t *testing.T) {
	config := &GormConfig{
		Database: DatabaseConfig{
			Type:     PostgreSQL,
			Host:     "localhost",
			Port:     5432,
			User:     "test",
			Password: "test",
			DBName:   "testdb",
			SSLMode:  "disable",
		},
	}

	manager := NewDatabaseManagerWithConfig(config)
	if manager == nil {
		t.Error("DatabaseManager should not be nil")
	}

	if manager.config != config {
		t.Error("DatabaseManager should have the correct config")
	}
}

// TestConfigValidation tests the configuration validation
func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name    string
		config  *GormConfig
		wantErr bool
	}{
		{
			name: "Valid PostgreSQL config",
			config: &GormConfig{
				Database: DatabaseConfig{
					Type:   PostgreSQL,
					Host:   "localhost",
					User:   "test",
					DBName: "testdb",
				},
			},
			wantErr: false,
		},
		{
			name: "Valid SQLite config",
			config: &GormConfig{
				Database: DatabaseConfig{
					Type: SQLite,
					File: "./test.db",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid PostgreSQL config - missing host",
			config: &GormConfig{
				Database: DatabaseConfig{
					Type:   PostgreSQL,
					User:   "test",
					DBName: "testdb",
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid SQLite config - missing file",
			config: &GormConfig{
				Database: DatabaseConfig{
					Type: SQLite,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestSetDefaults tests the default value setting
func TestSetDefaults(t *testing.T) {
	config := &GormConfig{}
	config.SetDefaults()

	// Check that defaults are set
	if config.Log.Level == 0 {
		t.Error("Log level should be set to default")
	}
	if config.Log.SlowThreshold == 0 {
		t.Error("Slow threshold should be set to default")
	}
	if config.Pool.MaxIdleConns == 0 {
		t.Error("Max idle conns should be set to default")
	}
	if config.Pool.MaxOpenConns == 0 {
		t.Error("Max open conns should be set to default")
	}
}

// TestPoolConfig tests the pool configuration functionality
func TestPoolConfig(t *testing.T) {
	config := &GormConfig{
		Database: DatabaseConfig{
			Type: SQLite,
			File: "./test.db",
		},
	}

	manager := NewDatabaseManager(config)

	// Test GetPoolConfig
	poolConfig := manager.GetPoolConfig()
	if poolConfig.MaxIdleConns != 0 {
		t.Error("Pool config should be empty initially")
	}

	// Test setting pool config
	newPoolConfig := PoolConfig{
		MaxIdleConns: 5,
		MaxOpenConns: 20,
	}

	// Note: SetPoolConfig will fail because db is not connected
	// This is expected behavior
	err := manager.SetPoolConfig(newPoolConfig)
	if err == nil {
		t.Error("SetPoolConfig should fail when db is not connected")
	}
}
