package interceptors

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"mix/pkg/log"
)

// Logger is an interceptor to log the rpc handlers.
func Logger(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	resp, err := handler(ctx, req)

	logger := log.ExtractLogger(ctx)
	logger.Infow(info.FullMethod, "latency", time.Since(start))
	return resp, err
}
