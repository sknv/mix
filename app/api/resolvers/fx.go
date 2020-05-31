package resolvers

import (
	"go.uber.org/fx"
)

var Module = []fx.Option{
	fx.Provide(
		NewResolver,
		NewMutationResolver,
		NewQueryResolver,
	),
}
