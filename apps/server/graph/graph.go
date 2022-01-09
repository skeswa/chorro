package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/skeswa/chorro/apps/server/cache"
	"github.com/skeswa/chorro/apps/server/graph/resolver"
	"github.com/skeswa/chorro/apps/server/graph/server"
	"github.com/skeswa/chorro/apps/server/session"
)

// Initialize everything GraphQL related for the server.
func Setup(cache *cache.Cache, mux *http.ServeMux) {
	resolver := &resolver.Resolver{}

	executableSchema := server.NewExecutableSchema(server.Config{
		Resolvers: resolver,
	})
	graphQLServer := handler.NewDefaultServer(executableSchema)

	respondWithGraphQLPlayground := playground.Handler(
		"GraphQL playground",
		"/api",
	)

	mux.HandleFunc("/api/playground", func(responseWriter http.ResponseWriter, request *http.Request) {
		// GET only.
		if request.Method != http.MethodGet {
			responseWriter.WriteHeader(http.StatusNotFound)

			return
		}

		session := session.Read(cache, request, responseWriter)
		if !session.IsUserLoggedIn {
			responseWriter.WriteHeader(http.StatusUnauthorized)

			return
		}

		respondWithGraphQLPlayground(responseWriter, request)
	})
	mux.Handle("/api", graphQLServer)
}
