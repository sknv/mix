package interceptors

import (
	"context"

	"google.golang.org/grpc"

	"mix/pkg/log"
	"mix/pkg/rand"
)

// RequestId is an interceptor that injects a request id into the context of each request.
func RequestId(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	var requestId string // TODO: fetch request id from incoming meta
	if requestId == "" {
		requestId = rand.String(16)
	}
	newCtx := log.PutRequestId(ctx, requestId)

	return handler(newCtx, req)
}
