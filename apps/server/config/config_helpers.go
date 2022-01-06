package config

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

// Returns a configuration struct for Fiber's Redis storage adapter.
func (c *Config) ForFiberRedis() redis.Config {
	return redis.Config{
		Host:     c.Redis.Host,
		Port:     c.Redis.Port,
		Username: "",
		Password: c.Redis.Password,
		Database: 0,
		Reset:    false,
	}
}

// Returns a configuration struct for Fiber's session middleware that leans on
// the specified storage to keep track of cookies internally.
func (c *Config) ForFiberSession(storage fiber.Storage) session.Config {
	// Sessions should be 30 days long (unless we're in dev where the session
	// should be a year because who cares).
	sessionLength := 30 * (24 * time.Hour)
	if c.Environment == DevEnvironment {
		sessionLength = 365 * (24 * time.Hour)
	}

	// Limit session cookie domain to just chorro.app and all its TLDs (unless
	// we're in dev where all domains can see the cookie).
	sessionCookieDomain := fmt.Sprintf(".%s", productionDomain)
	if c.Environment == DevEnvironment {
		sessionCookieDomain = ""
	}

	return session.Config{
		CookieDomain: sessionCookieDomain,
		Expiration:   sessionLength,
		KeyLookup:    "cookie:sid",
		Storage:      storage,
	}
}

// Returns the base URL that external clients will use to reach this server over
// HTTP(s).
func (c *Config) PublicBaseUrl() string {
	if c.Environment == DevEnvironment {
		return fmt.Sprintf("http://localhost:%d", c.HttpPort)
	}

	return fmt.Sprintf("https://%s", productionDomain)
}
