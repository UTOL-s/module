package fxconfig

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	App struct {
		Name string `mapstructure:"name"`
		Port string `mapstructure:"port"`
	} `mapstructure:"app"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`
	Accessor *Accessor
}

var configAccessor = &Accessor{}

// Accessor provides Yokai-style config access
// e.g., fxConfig.Config().String("app.name")
type Accessor struct{}

func (a *Accessor) String(key string) string {
	return viper.GetString(key)
}
func (a *Accessor) Int(key string) int {
	return viper.GetInt(key)
}
func (a *Accessor) Bool(key string) bool {
	return viper.GetBool(key)
}
func (a *Accessor) Float64(key string) float64 {
	return viper.GetFloat64(key)
}
func (a *Accessor) StringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
func (a *Accessor) AllSettings() map[string]any {
	return viper.AllSettings()
}

// ConfigAccessor returns the global config accessor, Yokai-style
func ConfigAccessor() *Accessor {
	return configAccessor
}

func NewConfig() (*Config, error) {
	// Reset viper to ensure clean state for each config load
	viper.Reset()

	// Load .env file if it exists
	envFiles := []string{".env", ".env.local", ".env.development", ".env.production"}
	for _, envFile := range envFiles {
		if _, err := os.Stat(envFile); err == nil {
			if err := gotenv.Load(envFile); err != nil {
				return nil, fmt.Errorf("error loading %s: %w", envFile, err)
			}
			break // Load only the first .env file found
		}
	}

	// Configure viper settings
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Load relevant OS environment variables into viper
	// Focus on common config prefixes to avoid loading all system env vars
	relevantPrefixes := []string{"APP_", "SERVER_", "DB_", "DATABASE_", "SUPERTOKENS_", "CONFIG_"}
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			envKey := pair[0]
			// Only load environment variables with relevant prefixes
			for _, prefix := range relevantPrefixes {
				if strings.HasPrefix(envKey, prefix) {
					key := strings.ToLower(strings.ReplaceAll(envKey, "_", "."))
					viper.Set(key, pair[1])
					break
				}
			}
		}
	}

	// Read and expand env variables in config.yaml
	configFile := "./configs/config.yaml"
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	expanded := expandEnvWithDefaults(string(data))
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(expanded)); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	config.Accessor = configAccessor
	return &config, nil
}

// GetEnv is still available for direct env access
func GetEnv(key string) string {
	return os.Getenv(key)
}

// expandEnvWithDefaults expands environment variables with support for default values
// Supports syntax like ${VAR:-default} where default is used if VAR is not set
func expandEnvWithDefaults(s string) string {
	// Regular expression to match ${VAR:-default} pattern
	re := regexp.MustCompile(`\$\{([^:}]+)(?::-([^}]+))?\}`)

	return re.ReplaceAllStringFunc(s, func(match string) string {
		// Extract variable name and default value
		parts := re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}

		varName := parts[1]
		defaultValue := ""
		if len(parts) > 2 {
			defaultValue = parts[2]
		}

		// Get environment variable value
		envValue := os.Getenv(varName)
		if envValue != "" {
			return envValue
		}

		// Return default value if environment variable is not set
		return defaultValue
	})
}

// NewConfigAccessor returns the config accessor for DI
func NewConfigAccessor() *Accessor {
	return configAccessor
}

func (c *Config) PostgresDSN() string {
	return "host=" + c.Database.Host +
		" user=" + c.Database.User +
		" password=" + c.Database.Password +
		" dbname=" + c.Database.DBName +
		" port=" + fmt.Sprintf("%d", c.Database.Port) +
		" sslmode=" + c.Database.SSLMode
}
