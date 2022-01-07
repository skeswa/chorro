package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/pkg/errors"
	"github.com/skeswa/chorro/apps/server/config"
	"github.com/skeswa/goth_fiber"
)

// Reads session data from the incoming request, and writes session data to the
// outgoing response.
type Store struct {
	// Fiber session store wrapped by this struct.
	store *session.Store
}

// Creates, initializes, and returns a new Store according to the specified
// configuration.
func New(configuration config.Config) *Store {
	// For session management, we talk to Redis.
	fiberRedisStore := redis.New(configuration.ForFiberRedis())
	store := session.New(configuration.ForFiberSession(fiberRedisStore))

	return &Store{store}
}

// Tell Goth to use this session Store to keep track of the session.
func (store *Store) BindToGoth() {
	// We use Goth to manage authentication against 3rd party providers. To add
	// compatibility between our web server, Fiber, and Goth, we use a cool little
	// library called goth_fiber. It needs to be told about our session store so
	// it can automate OAuth handshake.
	goth_fiber.SessionStore = store.store
}

// Reads the Session from the given context.
func (store *Store) Session(context *fiber.Ctx) (Session, error) {
	rawSession, err := store.store.Get(context)
	if err != nil {
		return Session{}, errors.Wrap(err, "Failed to wrap raw session")
	}

	return wrapRawSession(rawSession), nil
}
