package api

import (
	"go.uber.org/fx"

	"mix/app/api/resolvers"
)

var Module = append(
	[]fx.Option{
		fx.Provide(NewApplication),
	},
	resolvers.Module...,
)
