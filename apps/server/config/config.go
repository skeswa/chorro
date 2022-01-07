package config

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Settings and other values that affect how the server should perform user
// authentication.
type AuthConfig struct {
	// Client ID for Google OAuth.
	GoogleOAuth2ClientID string
	// Client secret for Google OAuth.
	GoogleOAuth2ClientSecret string
}

// Settings and other values that affect how the server should function.
type Config struct {
	// Settings and other values that affect how the server should perform user
	// authentication.
	Auth AuthConfig
	// Broadly describes the circumstances under which the server is running.
	Environment Environment
	// Port over which HTTP requests should be serviced.
	HttpPort int
	// URL of the home page.
	HomeUrl string
	// Settings and other values that affect how the server should interface with
	// Postgres.
	Postgres PostgresConfig
	// Settings and other values that affect how the server should interface with
	// Redis.
	Redis RedisConfig
}

// Settings and other values that affect how the server should interface with
// Postgres.
type PostgresConfig struct {
	// Name of the server's Postgres database.
	DatabaseName string
	// Network address of Postgres.
	Host string
	// Password of the server's Postgres user.
	Password string
	// TCP port of Postgres.
	Port int
	// Name of the server's Postgres user.
	User string
}

// Settings and other values that affect how the server should interface with
// Redis.
type RedisConfig struct {
	// Network address of Redis.
	Host string
	// Password of the server's Redis user.
	Password string
	// TCP port of Redis.
	Port int
}

// Reads Config from the environment and command line arguments.
func New() (Config, error) {
	config := Config{}

	if err := (&cli.App{
		Flags: []cli.Flag{
			// Auth flags.
			&cli.StringFlag{
				EnvVars: []string{"GOOGLE_CLIENT_ID"},
				Name:    googleOAuth2ClientIDFlagName,
				Usage:   "Google OAuth 2.0 client id",
			},
			&cli.StringFlag{
				EnvVars: []string{"GOOGLE_SECRET"},
				Name:    googleOAuth2ClientSecretFlagName,
				Usage:   "Google OAuth 2.0 client secret",
			},

			// General flags.
			&cli.StringFlag{
				EnvVars: []string{"HOME_URL"},
				Name:    homeUrlFlagName,
				Usage:   "URL of the home page",
				Value:   "/",
			},
			&cli.StringFlag{
				Aliases: []string{"e"},
				EnvVars: []string{"NODE_ENV"},
				Name:    environmentFlagName,
				Usage:   "Broadly describes the circumstances under which the server is running",
				Value:   "production",
			},
			&cli.IntFlag{
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Name:    portFlagName,
				Usage:   "Port over which HTTP requests should be serviced",
				Value:   8000,
			},

			// Postgres flags.
			&cli.StringFlag{
				EnvVars: []string{"POSTGRES_DATABASE_NAME"},
				Name:    postgresDatabaseNameFlagName,
				Usage:   "Name of the server's Postgres database",
				Value:   "chorro",
			},
			&cli.StringFlag{
				EnvVars: []string{"POSTGRES_HOST"},
				Name:    postgresHostFlagName,
				Usage:   "Network address of Postgres",
				Value:   "postgres",
			},
			&cli.StringFlag{
				EnvVars: []string{"POSTGRES_PASSWORD"},
				Name:    postgresPasswordFlagName,
				Usage:   "Password of the server's Postgres user",
			},
			&cli.IntFlag{
				EnvVars: []string{"POSTGRES_PORT"},
				Name:    postgresPortFlagName,
				Usage:   "TCP port of Postgres instance",
				Value:   5432,
			},
			&cli.StringFlag{
				EnvVars: []string{"POSTGRES_USER"},
				Name:    postgresUserFlagName,
				Usage:   "Name of the server's Postgres user",
				Value:   "chorro",
			},

			// Redis flags.
			&cli.StringFlag{
				EnvVars: []string{"REDIS_HOST"},
				Name:    redisHostFlagName,
				Usage:   "Network address of Redis instance",
				Value:   "redis",
			},
			&cli.StringFlag{
				EnvVars: []string{"REDIS_PASSWORD"},
				Name:    redisPasswordFlagName,
				Usage:   "Password of the server's Redis user",
			},
			&cli.IntFlag{
				EnvVars: []string{"REDIS_PORT"},
				Name:    redisPortFlagName,
				Usage:   "TCP port of Redis instance",
				Value:   5432,
			},
		},
		Action: func(c *cli.Context) error {
			return config.init(c)
		},
	}).Run(os.Args); err != nil {
		return config, err
	}

	return config, nil
}

const (
	// Auth flag names.
	googleOAuth2ClientIDFlagName     = "google-oauth2-client-id"
	googleOAuth2ClientSecretFlagName = "google-oauth2-client-secret"

	// General flag names.
	environmentFlagName = "environment"
	homeUrlFlagName     = "home-url"
	portFlagName        = "port"

	// Postgres flag names.
	postgresDatabaseNameFlagName = "postgres-database-name"
	postgresHostFlagName         = "postgres-host"
	postgresPasswordFlagName     = "postgres-password"
	postgresPortFlagName         = "postgres-port"
	postgresUserFlagName         = "postgres-user"

	// Redis flag names.
	redisHostFlagName     = "redis-host"
	redisPasswordFlagName = "redis-password"
	redisPortFlagName     = "redis-port"
)

// Initializes this Config by reading values from the provided cli.Context.
func (config *Config) init(c *cli.Context) error {
	if err := config.Auth.init(c); err != nil {
		return err
	}

	config.Environment = toEnvironment(c.String(environmentFlagName))

	if config.HomeUrl = c.String(homeUrlFlagName); len(config.HomeUrl) < 1 {
		return fmt.Errorf("%s is not a valid %s", config.HomeUrl, homeUrlFlagName)
	}

	if config.HttpPort = c.Int(portFlagName); !isValidPort(config.HttpPort) {
		return fmt.Errorf("%d is not a valid %s", config.HttpPort, portFlagName)
	}

	if err := config.Postgres.init(c); err != nil {
		return err
	}

	if err := config.Redis.init(c); err != nil {
		return err
	}

	return nil
}

// Initializes this AuthConfig by reading values from the provided cli.Context.
func (config *AuthConfig) init(c *cli.Context) error {
	if config.GoogleOAuth2ClientID = c.String(googleOAuth2ClientIDFlagName); len(config.GoogleOAuth2ClientID) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.GoogleOAuth2ClientID, googleOAuth2ClientIDFlagName)
	}

	if config.GoogleOAuth2ClientSecret = c.String(googleOAuth2ClientSecretFlagName); len(config.GoogleOAuth2ClientSecret) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.GoogleOAuth2ClientSecret, googleOAuth2ClientSecretFlagName)
	}

	return nil
}

// Initializes this PostgresConfig by reading values from the provided
// cli.Context.
func (config *PostgresConfig) init(c *cli.Context) error {
	if config.DatabaseName = c.String(postgresDatabaseNameFlagName); len(config.DatabaseName) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.DatabaseName, postgresDatabaseNameFlagName)
	}

	if config.Host = c.String(postgresHostFlagName); len(config.Host) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.Host, postgresHostFlagName)
	}

	config.Password = c.String(postgresPasswordFlagName)

	if config.Port = c.Int(postgresPortFlagName); !isValidPort(config.Port) {
		return fmt.Errorf("%d is not a valid %s", config.Port, postgresPortFlagName)
	}

	if config.User = c.String(postgresUserFlagName); len(config.User) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.User, postgresUserFlagName)
	}

	return nil
}

// Initializes this RedisConfig by reading values from the provided cli.Context.
func (config *RedisConfig) init(c *cli.Context) error {
	if config.Host = c.String(redisHostFlagName); len(config.Host) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.Host, redisHostFlagName)
	}

	config.Password = c.String(redisPasswordFlagName)

	if config.Port = c.Int(redisPortFlagName); !isValidPort(config.Port) {
		return fmt.Errorf("%d is not a valid %s", config.Port, redisPortFlagName)
	}

	return nil
}

// Returns true if port is valid.
func isValidPort(port int) bool {
	return port > 1024 && port < 99999
}
