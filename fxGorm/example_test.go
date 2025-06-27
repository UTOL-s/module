package fxgorm

import (
	"testing"

	fxconfig "github.com/UTOL-s/module/fxConfig"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// ExampleGormConfig demonstrates how to create a GORM configuration
func ExampleNewGormConfig() {
	// This would typically come from your main config
	config := &fxconfig.Config{}

	// Create GORM configuration
	gormConfig := NewGormConfig(config)
	_ = gormConfig
}

// ExampleFxGorm demonstrates how to use the FX module
func ExampleFxGorm() {
	app := fx.New(
		FxGorm,
		fx.Invoke(func(db *gorm.DB) {
			// Your application logic here
			_ = db
		}),
	)

	app.Run()
}

// TestGormConfig tests the configuration creation
func TestGormConfig(t *testing.T) {
	config := &fxconfig.Config{}
	gormConfig := NewGormConfig(config)

	if gormConfig == nil {
		t.Error("GormConfig should not be nil")
	}
}

// TestDatabaseTypes tests the database type constants
func TestDatabaseTypes(t *testing.T) {
	types := []DatabaseType{PostgreSQL, MySQL, SQLite, SQLServer}

	for _, dbType := range types {
		if dbType == "" {
			t.Errorf("Database type should not be empty")
		}
	}
}

// TestBuildDSN tests DSN building for different database types
func TestBuildDSN(t *testing.T) {
	tests := []struct {
		name    string
		config  *GormConfig
		wantErr bool
	}{
		{
			name: "PostgreSQL DSN",
			config: &GormConfig{
				Database: DatabaseConfig{
					Type:     PostgreSQL,
					Host:     "localhost",
					Port:     5432,
					User:     "postgres",
					Password: "password",
					DBName:   "testdb",
					SSLMode:  "disable",
				},
			},
			wantErr: false,
		},
		{
			name: "SQLite DSN",
			config: &GormConfig{
				Database: DatabaseConfig{
					Type: SQLite,
					File: "./test.db",
				},
			},
			wantErr: false,
		},
		{
			name: "SQLite DSN without file",
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
			dsn, err := tt.config.buildDSN()
			if (err != nil) != tt.wantErr {
				t.Errorf("buildDSN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && dsn == "" {
				t.Error("buildDSN() returned empty DSN")
			}
		})
	}
}
