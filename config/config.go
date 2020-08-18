package config

// Config defines the configuration structure.
type Config struct {
	General struct {
		LogLevel       int `mapstructure:"log_level"`
		MaxConnections int `mapstructure:"max_connections"`
	} `mapstructure:"general"`
}

// C holds the global configuration.
var C Config
