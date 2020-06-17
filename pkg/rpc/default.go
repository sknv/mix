package rpc

import (
	grpcware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"mix/pkg/rpc/interceptors"
)

func NewServer(options ...grpc.ServerOption) *grpc.Server {
	opts := append(
		options,
		grpc.UnaryInterceptor(
			grpcware.ChainUnaryServer(
				interceptors.RequestId,
				interceptors.Logger,
			),
		),
	)
	return grpc.NewServer(opts...)
}
