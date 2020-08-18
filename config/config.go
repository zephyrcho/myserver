package config

// Config defines the configuration structure.
type Config struct {
	General struct {
		LogLevel         int `mapstructure:"log_level"`
		max_idle_connect int `mapstructure:"max_idle_connect"`
	} `mapstructure:"general"`
}

// C holds the global configuration.
var C Config
