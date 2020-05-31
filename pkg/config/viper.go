package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// ParseConfig will parse the configuration from the environment variables and a file with the specified path.
func ParseConfig(filePath string, config interface{}) error {
	// Parse the file
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read the config file")
	}

	// Unmarshal the config
	if err := viper.Unmarshal(config); err != nil {
		return errors.Wrap(err, "failed to unmarshal the configuration")
	}

	return nil
}
