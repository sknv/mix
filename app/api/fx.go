package api

import (
	"go.uber.org/fx"

	"mix/app/api/controllers"
)

var Module = append(
	[]fx.Option{
		fx.Provide(NewApplication),
	},
	controllers.Module...,
)
