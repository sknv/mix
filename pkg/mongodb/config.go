package mongodb

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

const (
	DefaultHost                 = "127.0.0.1"
	DefaultPort                 = 27017
	DefaultMigrationsCollection = "schema_migrations"
)

// Config is a db config.
type Config struct {
	Host                 string
	Port                 int
	User                 string
	Password             string
	Database             string
	MigrationsDir        string
	MigrationsCollection string
}

// Validate validates the db config.
func (c *Config) Validate() error {
	var err error
	if c.Host == "" {
		err = multierr.Append(err, errors.New("empty database host provided"))
	}
	if c.Port == 0 {
		err = multierr.Append(err, errors.New("empty database port provided"))
	}
	if c.User == "" {
		err = multierr.Append(err, errors.New("empty database user provided"))
	}
	if c.Password == "" {
		err = multierr.Append(err, errors.New("empty database password provided"))
	}
	if c.Database == "" {
		err = multierr.Append(err, errors.New("empty database name provided"))
	}
	if c.MigrationsDir == "" {
		err = multierr.Append(err, errors.New("empty migrations directory provided"))
	}
	if c.MigrationsCollection == "" {
		err = multierr.Append(err, errors.New("empty migrations collection name provided"))
	}
	return err
}

// URI returns a connection string.
func (c *Config) URI() string {
	return fmt.Sprintf("mongodb://%s:%d", c.Host, c.Port)
}

// MigrationConnectionString returns a connection string for migrator.
func (c *Config) MigrationConnectionString() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s?authSource=admin&x-migrations-collection=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.MigrationsCollection,
	)
}
