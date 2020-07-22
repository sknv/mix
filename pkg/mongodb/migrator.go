package mongodb

import (
	"context"
	"os/exec"

	"github.com/pkg/errors"
)

// Migrate migrates the database.
func Migrate(ctx context.Context, config Config) error {
	cmd := exec.CommandContext(
		ctx,
		"migrate",
		"-path", config.MigrationsDir,
		"-database", config.MigrationConnectionString(),
		"up",
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrapf(err, "failed to run a command: %s", out)
	}
	return nil
}
