package main

import (
	"context"
	"flag"
	stdlog "log"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/fx"

	"mix/app/api"
	"mix/app/api/config"
	"mix/pkg/log"
	"mix/pkg/web"
)

const (
	defaultConfigPath = "configs/api/config.toml"

	maxRequestsAllowed = 5000
)

func main() {
	// Parse flags
	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	// Parse the config
	cfg, err := config.ParseConfig(*configPath)
	if err != nil {
		stdlog.Fatalf("failed to parse the config file: %+v", err)
	}

	// Provide dependencies
	options := append(
		[]fx.Option{
			fx.Provide(func() *config.Config { return cfg }),
			fx.Provide(NewRouter),

			fx.Invoke(BuildLogger),
			fx.Invoke(Route),
		},
		api.Module...,
	)
	app := fx.New(options...)
	app.Run()
}

func BuildLogger(lc fx.Lifecycle, config *config.Config) {
	// Build a logger
	if err := log.Build(config.Application.LogLevel); err != nil {
		stdlog.Fatalf("failed to build a logger: %+v", err)
	}

	// Remember to flush the logger
	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			log.Logger().Sync()
			return nil
		},
	})
}

func NewRouter(lc fx.Lifecycle, config *config.Config) chi.Router {
	// Create a router
	router := web.NewRouter(
		web.Throttle(maxRequestsAllowed),
	)

	// Create an http server
	server := http.Server{
		Addr:    config.Application.Addr,
		Handler: router,
	}

	// Start and stop the http server
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go web.Start(&server)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return web.Shutdown(ctx, &server)
		},
	})

	return router
}

// Route routes and handles http requests
func Route(app *api.Application, router chi.Router) {
	app.Route(router)
}
