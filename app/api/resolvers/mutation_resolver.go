package resolvers

import (
	"context"

	"mix/app/api/models"
)

type MutationResolver struct {
	*Resolver
}

func NewMutationResolver(resolver *Resolver) *MutationResolver {
	return &MutationResolver{Resolver: resolver}
}

func (r *MutationResolver) CreateAccount(ctx context.Context, input models.NewAccount) (*models.Account, error) {
	panic("not implemented")
}
