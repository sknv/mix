package rpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"mix/pkg/log"
)

func Start(server *grpc.Server, listener net.Listener) {
	logger := log.Logger()
	logger.Infow("starting an rpc server", "address", listener.Addr())
	if err := server.Serve(listener); err != nil {
		// Cannot panic, because this probably is an intentional close
		logger.Infow("shutting down the grpc server", "reason", err)
	}
}

func Stop(ctx context.Context, server *grpc.Server) {
	// Try to stop the server gracefully
	serverStoppedGracefully := make(chan struct{})
	go func() {
		server.GracefulStop()
		serverStoppedGracefully <- struct{}{}
	}()

	// Wait for a graceful shutdown and then stop the server forcibly
	logger := log.Logger()
	select {
	case <-serverStoppedGracefully:
		logger.Info("rpc server stopped")
	case <-ctx.Done():
		server.Stop()
		logger.Info("rpc server forcibly stopped")
	}
}
