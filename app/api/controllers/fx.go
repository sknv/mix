package controllers

import (
	"go.uber.org/fx"
)

var Module = []fx.Option{
	fx.Provide(NewAccountController),
}
