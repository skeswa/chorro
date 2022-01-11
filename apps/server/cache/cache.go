package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/skeswa/chorro/apps/server/config"
)

// Interface between the server and the cache service it uses for hot/live data
// storage.
type Cache struct {
	// Redis session store.
	SessionStore sessions.Store

	// Redis cache client.
	client *redis.Client
}

// Creates and initializes a new Cache based on the specified config.
func New(config *config.Config) (*Cache, error) {
	client := redis.NewClient(config.ForRedis())

	// As sessions are stored in the cache, we create an instance of the
	// gorilla/sessions here for Redis.
	sessionStore, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create Redis session store")
	}

	// Before the session store can be used, it must be configured.
	sessionStore.KeyPrefix("chorro_")
	sessionStore.Options(config.ForSessionStore())

	return &Cache{
		client:       client,
		SessionStore: sessionStore,
	}, nil
}
