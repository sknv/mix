package resolvers

import (
	"context"

	"mix/app/api/models"
)

type QueryResolver struct {
	*Resolver
}

func NewQueryResolver(resolver *Resolver) *QueryResolver {
	return &QueryResolver{Resolver: resolver}
}

func (r *QueryResolver) Accounts(ctx context.Context) ([]*models.Account, error) {
	panic("not implemented")
}
