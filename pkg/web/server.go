package web

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	"mix/pkg/log"
)

// Start starts a web server.
func Start(server *http.Server) {
	logger := log.Logger()
	logger.Infow("starting an http server", "address", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		// Cannot panic, because this probably is an intentional close
		logger.Infow("shutting down the http server", "reason", err)
	}
}

// Shutdown stops a web server.
func Shutdown(ctx context.Context, server *http.Server) error {
	logger := log.Logger()
	if err := server.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to shutdown the http server")
	}
	logger.Info("http server stopped")
	return nil
}
