package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"mix/app/api/graphql"
	"mix/app/api/resolvers"
)

type Application struct {
	Resolver *resolvers.Resolver
}

func NewApplication(
	resolver *resolvers.Resolver,
) *Application {
	return &Application{
		Resolver: resolver,
	}
}

func (a *Application) Route(router chi.Router) {
	router.Get("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graphql.NewExecutableSchema(graphql.Config{
			Resolvers: a.Resolver,
		}),
	))
}
