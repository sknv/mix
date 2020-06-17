package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"mix/pkg/config"
)

const (
	envPrefix = "mix.account"

	defaultEnv      = "production"
	defaultLogLevel = "info"
)

// Config is a global config.
type Config struct {
	Application Application
}

func (c *Config) Validate() error {
	return c.Application.Validate()
}

// Application config.
type Application struct {
	Env      string
	Addr     string
	LogLevel string
}

func (a *Application) Validate() error {
	if a.Addr == "" {
		return errors.New("empty address provided for an http server to start on")
	}
	return nil
}

func (a *Application) IsProduction() bool {
	return a.Env == "production"
}

// ParseConfig will parse the configuration from the environment variables and a file with the specified path.
// Environment variables have priority over ones specified in the file.
func ParseConfig(filePath string) (*Config, error) {
	bindEnvVars()
	setDefaults()

	// Parse the config
	var cfg Config
	if err := config.ParseConfig(filePath, &cfg); err != nil {
		return nil, err
	}

	// Validate the provided configuration
	if err := cfg.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate the config")
	}
	return &cfg, nil
}

func bindEnvVars() {
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func setDefaults() {
	viper.SetDefault("Application.Env", defaultEnv)
	viper.SetDefault("Application.Addr", "")
	viper.SetDefault("Application.LogLevel", defaultLogLevel)
}
