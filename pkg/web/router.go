package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	webware "mix/pkg/web/middleware"
)

// Option configures a Router.
type Option func(chi.Router)

// NewRouter returns a new router.
func NewRouter(options ...Option) chi.Router {
	router := chi.NewRouter()

	for _, opt := range options {
		opt(router)
	}

	// Add default middleware
	router.Use(
		middleware.RealIP,
		webware.RequestId,
		webware.Logger,
	)
	return router
}

// Throttle option.
func Throttle(limit int) Option {
	return func(r chi.Router) {
		r.Use(middleware.Throttle(limit))
	}
}
