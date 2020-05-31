package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"mix/app/api/graphql"
)

type Resolver struct{}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) Mutation() graphql.MutationResolver {
	return &MutationResolver{Resolver: r}
}

func (r *Resolver) Query() graphql.QueryResolver {
	return &QueryResolver{Resolver: r}
}
