package fxgorm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected DatabaseConfig
		wantErr  bool
	}{
		{
			name: "PostgreSQL URL",
			url:  "postgres://user:pass@localhost:5432/testdb?sslmode=disable",
			expected: DatabaseConfig{
				URL:      "postgres://user:pass@localhost:5432/testdb?sslmode=disable",
				Type:     PostgreSQL,
				Host:     "localhost",
				Port:     5432,
				User:     "user",
				Password: "pass",
				DBName:   "testdb",
				SSLMode:  "disable",
			},
			wantErr: false,
		},
		{
			name: "MySQL URL",
			url:  "mysql://root:password@localhost:3306/mydb?charset=utf8mb4&parseTime=true&loc=Local",
			expected: DatabaseConfig{
				URL:       "mysql://root:password@localhost:3306/mydb?charset=utf8mb4&parseTime=true&loc=Local",
				Type:      MySQL,
				Host:      "localhost",
				Port:      3306,
				User:      "root",
				Password:  "password",
				DBName:    "mydb",
				Charset:   "utf8mb4",
				ParseTime: true,
				Loc:       "Local",
			},
			wantErr: false,
		},
		{
			name: "SQLite URL",
			url:  "sqlite:./test.db",
			expected: DatabaseConfig{
				URL:  "sqlite:./test.db",
				Type: SQLite,
				File: "./test.db",
			},
			wantErr: false,
		},
		{
			name: "SQL Server URL",
			url:  "sqlserver://sa:password@localhost:1433?database=testdb",
			expected: DatabaseConfig{
				URL:      "sqlserver://sa:password@localhost:1433?database=testdb",
				Type:     SQLServer,
				Host:     "localhost",
				Port:     1433,
				User:     "sa",
				Password: "password",
				DBName:   "testdb",
			},
			wantErr: false,
		},
		{
			name:    "Invalid URL",
			url:     "invalid-url",
			wantErr: true,
		},
		{
			name:    "Unsupported scheme",
			url:     "redis://localhost:6379",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &GormConfig{
				Database: DatabaseConfig{
					URL: tt.url,
				},
				Pool: PoolConfig{
					MaxIdleConns:    10,
					MaxOpenConns:    100,
					ConnMaxLifetime: time.Hour,
					ConnMaxIdleTime: time.Minute * 10,
				},
				Log: LogConfig{
					Level:         logger.Info,
					SlowThreshold: time.Second * 5,
					Colorful:      true,
				},
			}

			err := config.parseURL()

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Type, config.Database.Type)
			assert.Equal(t, tt.expected.Host, config.Database.Host)
			assert.Equal(t, tt.expected.Port, config.Database.Port)
			assert.Equal(t, tt.expected.User, config.Database.User)
			assert.Equal(t, tt.expected.Password, config.Database.Password)
			assert.Equal(t, tt.expected.DBName, config.Database.DBName)
			assert.Equal(t, tt.expected.SSLMode, config.Database.SSLMode)
			assert.Equal(t, tt.expected.Charset, config.Database.Charset)
			assert.Equal(t, tt.expected.ParseTime, config.Database.ParseTime)
			assert.Equal(t, tt.expected.Loc, config.Database.Loc)
			assert.Equal(t, tt.expected.File, config.Database.File)
		})
	}
}

func TestBuildDSN_WithURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		expectedDSN string
		wantErr     bool
	}{
		{
			name:        "PostgreSQL URL",
			url:         "postgres://user:pass@localhost:5432/testdb?sslmode=disable",
			expectedDSN: "postgres://user:pass@localhost:5432/testdb?sslmode=disable",
			wantErr:     false,
		},
		{
			name:        "MySQL URL",
			url:         "mysql://root:password@localhost:3306/mydb?charset=utf8mb4",
			expectedDSN: "mysql://root:password@localhost:3306/mydb?charset=utf8mb4",
			wantErr:     false,
		},
		{
			name:        "SQLite URL",
			url:         "sqlite:./test.db",
			expectedDSN: "./test.db",
			wantErr:     false,
		},
		{
			name:        "SQL Server URL",
			url:         "sqlserver://sa:password@localhost:1433?database=testdb",
			expectedDSN: "sqlserver://sa:password@localhost:1433?database=testdb",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &GormConfig{
				Database: DatabaseConfig{
					URL: tt.url,
				},
			}

			dsn, err := config.buildDSN()

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedDSN, dsn)
		})
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "Valid PostgreSQL URL",
			url:     "postgres://user:pass@localhost:5432/testdb",
			wantErr: false,
		},
		{
			name:    "Valid MySQL URL",
			url:     "mysql://root:password@localhost:3306/mydb",
			wantErr: false,
		},
		{
			name:    "Valid SQLite URL",
			url:     "sqlite:./test.db",
			wantErr: false,
		},
		{
			name:    "Empty URL",
			url:     "",
			wantErr: true,
		},
		{
			name:    "Invalid URL format",
			url:     "not-a-url",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &GormConfig{
				Database: DatabaseConfig{
					URL: tt.url,
				},
			}

			err := config.validateURL()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
