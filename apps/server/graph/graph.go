package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/skeswa/chorro/apps/server/cache"
	"github.com/skeswa/chorro/apps/server/config"
	"github.com/skeswa/chorro/apps/server/db"
	"github.com/skeswa/chorro/apps/server/graph/resolver"
	"github.com/skeswa/chorro/apps/server/graph/server"
	"github.com/skeswa/chorro/apps/server/session"
)

// Initialize everything GraphQL related for the server.
func Setup(
	cache *cache.Cache,
	config *config.Config,
	db *db.DB,
	mux *http.ServeMux,
) {
	resolver := &resolver.Resolver{
		Cache:  cache,
		Config: config,
		DB:     db,
	}

	respondWithGraphQLPlayground := playground.Handler(
		"GraphQL playground",
		"/api",
	)

	// Serve the GraphQL Playground to authenticated users for easy exploration
	// of the API.
	mux.HandleFunc("/api/playground", func(responseWriter http.ResponseWriter, request *http.Request) {
		// GET only.
		if request.Method != http.MethodGet {
			responseWriter.WriteHeader(http.StatusNotFound)

			return
		}

		session := session.ExtractFrom(request.Context())
		if !session.IsUserLoggedIn {
			responseWriter.WriteHeader(http.StatusUnauthorized)

			return
		}

		respondWithGraphQLPlayground(responseWriter, request)
	})

	executableSchema := server.NewExecutableSchema(server.Config{
		Resolvers: resolver,
	})
	graphQLServer := handler.NewDefaultServer(executableSchema)

	// Serve the GraphQL API itself from /api.
	mux.Handle("/api", graphQLServer)
}
