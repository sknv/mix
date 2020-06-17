package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"mix/app/api/graphql"
	"mix/app/proto"
)

type Resolver struct {
	AccountClient proto.AccountClient
}

func NewResolver(
	accountClient proto.AccountClient,
) *Resolver {
	return &Resolver{
		AccountClient: accountClient,
	}
}

func (r *Resolver) Mutation() graphql.MutationResolver {
	return &MutationResolver{Resolver: r}
}

func (r *Resolver) Query() graphql.QueryResolver {
	return &QueryResolver{Resolver: r}
}
