package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Build builds the global logger instance.
func Build(level string) error {
	// Load the log level from the argument
	var lvl zapcore.Level
	if err := lvl.Set(level); err != nil {
		return errors.Wrap(err, "failed to set a logger level")
	}

	// Configure a logger instance
	cfg := zap.NewDevelopmentConfig()
	cfg.Development = false              // disable the development mode
	cfg.EncoderConfig.StacktraceKey = "" // disable the default stacktrace formatting
	cfg.Level = zap.NewAtomicLevelAt(lvl)

	logger, err := cfg.Build()
	if err != nil {
		return errors.Wrap(err, "failed to build a logger instance")
	}

	zap.ReplaceGlobals(logger)
	return nil
}

// Logger returns the default logger.
func Logger() *zap.SugaredLogger {
	return zap.S()
}
