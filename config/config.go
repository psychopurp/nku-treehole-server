package config

import (
	"nku-treehole-server/pkg/logger"

	"github.com/spf13/viper"
)

var (
	Config *Configuration
)

// Struct of Configuration instance.
// It include Database and Server configuration
type Configuration struct {
	Server   ServerConnection
	Database DatabaseConfiguration
}

// Struct of Database Configuration instance.
type DatabaseConfiguration struct {
	Driver       string `mapstructure:"DATABASE_DRIVER"`
	Dbname       string `mapstructure:"DATABASE_NAME"`
	Username     string `mapstructure:"DATABASE_USERNAME"`
	Password     string `mapstructure:"DATABASE_PASSWORD"`
	Host         string `mapstructure:"DATABASE_HOST"`
	Port         string `mapstructure:"DATABASE_PORT"`
	MaxLifetime  int    `mapstructure:"DATABASE_MAX_LIFETIME"`
	MaxOpenConns int    `mapstructure:"DATABASE_MAX_OPEN_CONNS"`
	MaxIdleConns int    `mapstructure:"DATABASE_MAX_IDLE_CONNS"`
}

// Struct of Server Configuration instance.
type ServerConnection struct {
	Port        string `mapstructure:"SERVER_PORT"`
	Secret      string `mapstructure:"SERVER_SECRET"`
	Mode        string `mapstructure:"SERVER_MODE"`
	Name        string `mapstructure:"SERVER_NAME"`
	ExpiresHour int64  `mapstructure:"SERVER_EXPIRES_HOUR"`
}

// Setup the configuration
func Setup(configPath string) {
	var (
		databaseConfiguration DatabaseConfiguration
		serverConfiguration   ServerConnection
	)

	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("Error reading config file, %s", err)
	}

	unmarshalConfiguration(&databaseConfiguration)
	unmarshalConfiguration(&serverConfiguration)

	configuration := Configuration{
		Database: databaseConfiguration,
		Server:   serverConfiguration,
	}

	Config = &configuration
}

// Helper to unmarshal
func unmarshalConfiguration(configuration interface{}) {
	err := viper.Unmarshal(configuration)
	if err != nil {
		logger.Fatalf("Unable to decode into struct, %v", err)
	}
}

// GetConfig return the configuration instance
func GetConfig() *Configuration {
	return Config
}
