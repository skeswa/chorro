package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/skeswa/chorro/apps/server/graph/resolver"
	"github.com/skeswa/chorro/apps/server/graph/server"
)

// Initialize everything GraphQL related for the server.
func Setup(mux *http.ServeMux) {
	resolver := &resolver.Resolver{}

	executableSchema := server.NewExecutableSchema(server.Config{
		Resolvers: resolver,
	})
	graphQLServer := handler.NewDefaultServer(executableSchema)

	mux.Handle("/api/playground", playground.Handler(
		"GraphQL playground",
		"/api",
	))
	mux.Handle("/api", graphQLServer)
}
