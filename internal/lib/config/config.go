package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

const (
	envPrefix = "mix"

	defaultEnv        = "production"
	defaultLogLevel   = "info"
	defaultBucketPath = "public/media"
)

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// Config is a global config.
type Config struct {
	Application Application
}

func (c *Config) Validate() error {
	return multierr.Combine(
		c.Application.Validate(),
	)
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// Application config.
type Application struct {
	Env        string
	Addr       string
	LogLevel   string
	Secret     string
	BucketPath string
}

func (a *Application) Validate() error {
	if a.Addr == "" {
		return errors.New("empty address provided for an http server to start on")
	}
	if a.Secret == "" {
		return errors.New("empty secret provided")
	}
	return nil
}

func (a *Application) IsProduction() bool {
	return a.Env == "production"
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// ParseConfig will parse the configuration from the environment variables and a file with the specified path.
// Environment variables have priority over ones specified in the file.
func ParseConfig(filePath string) (*Config, error) {
	setDefaults()

	// Parse the file
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read the config file")
	}

	bindEnvVars() // remember to parse the environment variables

	// Unmarshal the config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal the configuration")
	}

	// Validate the provided configuration
	if err := config.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate the config")
	}
	return &config, nil
}

func setDefaults() {
	viper.SetDefault("Application.Env", defaultEnv)
	viper.SetDefault("Application.Addr", "")
	viper.SetDefault("Application.LogLevel", defaultLogLevel)
	viper.SetDefault("Application.Secret", "")
	viper.SetDefault("Application.BucketPath", defaultBucketPath)
}

func bindEnvVars() {
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
