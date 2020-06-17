package main

import (
	"context"
	"flag"
	stdlog "log"
	"net"

	"github.com/pkg/errors"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"mix/app/account"
	"mix/app/account/config"
	"mix/app/account/service"
	"mix/app/proto"
	"mix/pkg/log"
	"mix/pkg/rpc"
)

const (
	defaultConfigPath = "configs/api/config.toml"
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
			fx.Provide(NewServer),

			fx.Invoke(BuildLogger),
			fx.Invoke(Route),
		},
		account.Module...,
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

func NewServer(lc fx.Lifecycle, config *config.Config) *grpc.Server {
	// Create an rpc server
	server := rpc.NewServer()

	// Start and stop the rpc server
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			lis, err := net.Listen("tcp", config.Application.Addr) // listen on the address
			if err != nil {
				return errors.Wrap(err, "failed to listen")
			}

			go rpc.Start(server, lis)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			rpc.Stop(ctx, server)
			return nil
		},
	})

	return server
}

// Route routes and handles rpc requests
func Route(server *grpc.Server, account *service.AccountService) {
	proto.RegisterAccountServer(server, account)
}
