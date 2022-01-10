package resolver

import (
	"github.com/skeswa/chorro/apps/server/cache"
	"github.com/skeswa/chorro/apps/server/config"
	"github.com/skeswa/chorro/apps/server/db"
)

// Core type available to every GraphQL resolver.
//
// This struct is how we facilitate "dependency injection" in our GraphQL
// resolvers.
//
// For information about this, and other parts of server/graph, check out
// https://gqlgen.com
type Resolver struct {
	// Interface between the server and the cache service it uses for hot/live
	// data storage.
	Cache *cache.Cache
	// Settings and other values that affect how the server should function.
	Config *config.Config
	// Interface between the server and its primary data store.
	DB *db.DB
}
