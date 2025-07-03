package fxconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigAccessor(t *testing.T) {
	accessor := NewConfigAccessor()

	if accessor == nil {
		t.Error("ConfigAccessor should not be nil")
	}
}

func TestGetEnv(t *testing.T) {
	// Set a test environment variable
	testKey := "TEST_CONFIG_KEY"
	testValue := "test_value"
	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey)

	result := GetEnv(testKey)
	if result != testValue {
		t.Errorf("Expected %s, got %s", testValue, result)
	}
}

func TestAccessorMethods(t *testing.T) {
	accessor := &Accessor{}

	// Test that methods don't panic
	_ = accessor.String("test")
	_ = accessor.Int("test")
	_ = accessor.Bool("test")
	_ = accessor.Float64("test")
	_ = accessor.AllSettings()
}

func TestConfigAccessorSingleton(t *testing.T) {
	accessor1 := ConfigAccessor()
	accessor2 := ConfigAccessor()

	if accessor1 != accessor2 {
		t.Error("ConfigAccessor should return the same instance")
	}
}

func TestNewConfigWithEnvFile(t *testing.T) {
	// Create a temporary .env file
	envContent := `
DB_HOST=test-host
DB_PORT=5433
DB_USER=test-user
DB_PASSWORD=test-password
DB_NAME=test-db
DB_SSLMODE=require
APP_NAME=Test App
SERVER_PORT=9090
SUPERTOKENS_CONNECTION_URI=test-uri
SUPERTOKENS_CONNECTION_API_KEY=test-key
`

	// Write .env file
	err := os.WriteFile(".env", []byte(envContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(".env")

	// Create a temporary config.yaml file
	configContent := `
app:
  name: ${APP_NAME}
  version: "1.0.0"
  environment: "development"
  debug: false

server:
  host: "0.0.0.0"
  port: ${SERVER_PORT}
  read_timeout: 30
  write_timeout: 30
  idle_timeout: 60
  max_header_bytes: 1048576

database:
  type: "postgres"
  host: ${DB_HOST}
  port: ${DB_PORT}
  user: ${DB_USER}
  password: ${DB_PASSWORD}
  dbname: ${DB_NAME}
  sslmode: ${DB_SSLMODE}

supertokens:
  connection_uri: ${SUPERTOKENS_CONNECTION_URI}
  connection_api_key: ${SUPERTOKENS_CONNECTION_API_KEY}
  app_name: "UTOL API"
  api_domain: "http://localhost:8080"
  website_domain: "http://localhost:3000"
  api_base_path: "/api/auth"
  web_base_path: "/api/auth"
  email:
    host: "smtp.gmail.com"
    password: "your_email_password"
    email: "your_email@gmail.com"
`

	// Create configs directory if it doesn't exist
	if err := os.MkdirAll("configs", 0755); err != nil {
		t.Fatal(err)
	}

	// Write config.yaml file
	err = os.WriteFile("configs/config.yaml", []byte(configContent), 0644)
	assert.NoError(t, err)
	defer os.RemoveAll("configs")

	// Load configuration
	config, err := NewConfig()
	assert.NoError(t, err)
	assert.NotNil(t, config)

	// Verify that environment variables were loaded correctly
	assert.Equal(t, "Test App", config.App.Name)
	assert.Equal(t, "9090", config.Accessor.String("server.port"))
	assert.Equal(t, "test-host", config.Database.Host)
	assert.Equal(t, 5433, config.Database.Port)
	assert.Equal(t, "test-user", config.Database.User)
	assert.Equal(t, "test-password", config.Database.Password)
	assert.Equal(t, "test-db", config.Database.DBName)
	assert.Equal(t, "require", config.Database.SSLMode)
	assert.Equal(t, "test-uri", config.Accessor.String("supertokens.connection_uri"))
	assert.Equal(t, "test-key", config.Accessor.String("supertokens.connection_api_key"))
}

func TestNewConfigWithDefaultValues(t *testing.T) {
	// Clean up any existing environment variables that might interfere
	envVars := []string{
		"APP_NAME", "APP_VERSION", "APP_ENVIRONMENT",
		"SERVER_HOST", "SERVER_PORT",
		"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSLMODE",
		"SUPERTOKENS_CONNECTION_URI", "SUPERTOKENS_CONNECTION_API_KEY",
		"SUPERTOKENS_APP_NAME", "SUPERTOKENS_API_DOMAIN", "SUPERTOKENS_WEBSITE_DOMAIN",
		"SUPERTOKENS_API_BASE_PATH", "SUPERTOKENS_WEB_BASE_PATH",
		"SUPERTOKENS_EMAIL_HOST", "SUPERTOKENS_EMAIL_PASSWORD", "SUPERTOKENS_EMAIL",
	}

	for _, envVar := range envVars {
		os.Unsetenv(envVar)
	}

	// Create a temporary config.yaml file with default values
	configContent := `
app:
  name: ${APP_NAME:-"Default App Name"}
  version: ${APP_VERSION:-"1.0.0"}
  environment: ${APP_ENVIRONMENT:-"development"}
  debug: false

server:
  host: ${SERVER_HOST:-"0.0.0.0"}
  port: ${SERVER_PORT:-"8080"}
  read_timeout: 30
  write_timeout: 30
  idle_timeout: 60
  max_header_bytes: 1048576

database:
  type: "postgres"
  host: ${DB_HOST:-"localhost"}
  port: ${DB_PORT:-5432}
  user: ${DB_USER:-"postgres"}
  password: ${DB_PASSWORD:-"password"}
  dbname: ${DB_NAME:-"utol_api"}
  sslmode: ${DB_SSLMODE:-"disable"}

supertokens:
  connection_uri: ${SUPERTOKENS_CONNECTION_URI}
  connection_api_key: ${SUPERTOKENS_CONNECTION_API_KEY}
  app_name: ${SUPERTOKENS_APP_NAME:-"UTOL API"}
  api_domain: ${SUPERTOKENS_API_DOMAIN:-"http://localhost:8080"}
  website_domain: ${SUPERTOKENS_WEBSITE_DOMAIN:-"http://localhost:3000"}
  api_base_path: ${SUPERTOKENS_API_BASE_PATH:-"/api/auth"}
  web_base_path: ${SUPERTOKENS_WEB_BASE_PATH:-"/api/auth"}
  email:
    host: ${SUPERTOKENS_EMAIL_HOST:-"smtp.gmail.com"}
    password: ${SUPERTOKENS_EMAIL_PASSWORD}
    email: ${SUPERTOKENS_EMAIL}
`

	// Create configs directory if it doesn't exist
	if err := os.MkdirAll("configs", 0755); err != nil {
		t.Fatal(err)
	}

	// Write config.yaml file
	err := os.WriteFile("configs/config.yaml", []byte(configContent), 0644)
	assert.NoError(t, err)
	defer os.RemoveAll("configs")

	// Load configuration without .env file
	config, err := NewConfig()
	assert.NoError(t, err)
	assert.NotNil(t, config)

	// Verify that default values are used when environment variables are not set
	assert.Equal(t, "Default App Name", config.App.Name)
	assert.Equal(t, "1.0.0", config.Accessor.String("app.version"))
	assert.Equal(t, "development", config.Accessor.String("app.environment"))
	assert.Equal(t, "0.0.0.0", config.Accessor.String("server.host"))
	assert.Equal(t, "8080", config.Accessor.String("server.port"))
	assert.Equal(t, "localhost", config.Database.Host)
	assert.Equal(t, 5432, config.Database.Port)
	assert.Equal(t, "postgres", config.Database.User)
	assert.Equal(t, "password", config.Database.Password)
	assert.Equal(t, "utol_api", config.Database.DBName)
	assert.Equal(t, "disable", config.Database.SSLMode)
	assert.Equal(t, "UTOL API", config.Accessor.String("supertokens.app_name"))
	assert.Equal(t, "http://localhost:8080", config.Accessor.String("supertokens.api_domain"))
	assert.Equal(t, "http://localhost:3000", config.Accessor.String("supertokens.website_domain"))
	assert.Equal(t, "/api/auth", config.Accessor.String("supertokens.api_base_path"))
	assert.Equal(t, "/api/auth", config.Accessor.String("supertokens.web_base_path"))
	assert.Equal(t, "smtp.gmail.com", config.Accessor.String("supertokens.email.host"))
}
