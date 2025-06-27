package fxconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
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
func (a *Accessor) AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

// ConfigAccessor returns the global config accessor, Yokai-style
func ConfigAccessor() *Accessor {
	return configAccessor
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read and expand env variables in config.yaml
	configFile := "./configs/config.yaml"
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	expanded := os.ExpandEnv(string(data))
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
	return viper.GetString(key)
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
