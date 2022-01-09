package config

import (
	"fmt"
	"math"
	"net/http"
	"net/url"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

// Returns CORS configuration for the HTTP server.
func (c *Config) ForCors() cors.Options {
	// Take all the default options in development.
	if c.Environment == DevEnvironment {
		return cors.Options{}
	}

	homeUrlWithWildcardDomain, err := url.Parse(c.HomeUrl)
	if err != nil {
		panic(fmt.Sprintf("\"%s\" is not a valid URL", c.HomeUrl))
	}

	homeUrlWithWildcardDomain.Path = ""
	homeUrlWithWildcardDomain.Host = "*." + homeUrlWithWildcardDomain.Hostname()

	return cors.Options{
		AllowedOrigins:   []string{homeUrlWithWildcardDomain.String()},
		AllowCredentials: true,
	}
}

// Returns a configuration struct for Fiber's session middleware that leans on
// the specified storage to keep track of cookies internally.
func (c *Config) ForSessionStore() sessions.Options {
	// Restrict cookies to secured connectoins only in non-development
	// environments.
	isHttpsOnly := c.Environment != DevEnvironment

	// Prefix for paths to which session cookies should apply.
	pathPrefix := "/"

	// Allow cross-site cookie transfer to allow for authenticated requests to
	// sub-domains.
	//
	// For more information, see
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite#lax
	sameSiteMode := http.SameSiteLaxMode

	// Sessions should be 30 days long (unless we're in dev where the session
	// should be as long as possible because who cares).
	sessionLengthInSeconds := 30 * 24 * 60 * 60
	if c.Environment == DevEnvironment {
		sessionLengthInSeconds = math.MaxInt32
	}

	// Limit session cookie domain to just chorro.app and all its TLDs (unless
	// we're in dev where all domains can see the cookie).
	sessionCookieDomain := fmt.Sprintf(".%s", productionDomain)
	if c.Environment == DevEnvironment {
		sessionCookieDomain = ""
	}

	return sessions.Options{
		Domain:   sessionCookieDomain,
		MaxAge:   sessionLengthInSeconds,
		Path:     pathPrefix,
		SameSite: sameSiteMode,
		Secure:   isHttpsOnly,
	}
}

// Returns a configuration struct for Gorm.
func (c *Config) ForGorm() *gorm.Config {
	return &gorm.Config{}
}

// Returns a DSN for connecting Gorm to Postgres.
func (c *Config) ForGormPostgresDriver() string {
	return fmt.Sprintf(
		"host=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"port=%d "+
			"sslmode=disable "+
			"TimeZone=America/New_York",
		c.Postgres.Host,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.DatabaseName,
		c.Postgres.Port,
	)
}

// Returns a configuration struct for Redis.
func (c *Config) ForRedis() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		DB:       0,
		Password: c.Redis.Password,
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
