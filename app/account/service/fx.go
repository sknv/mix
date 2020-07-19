package service

import (
	"go.uber.org/fx"

	"mix/app/account/repos"
)

var Module = append(
	[]fx.Option{
		fx.Provide(NewAccountService),
	},
	repos.Module...,
)
