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
	// Settings and other values that affect how the server should interface with
	// the database.
	Database DatabaseConfig
	// Broadly describes the circumstances under which the server is running.
	Environment Environment
	// Port over which HTTP requests should be serviced.
	HttpPort int
}

// Returns the base URL that external clients will use to reach this server over
// HTTP(s).
func (c *Config) PublicBaseUrl() string {
	if c.Environment == DevEnvironment {
		return fmt.Sprintf("http://localhost:%d", c.HttpPort)
	}

	return "https://chorro.app"
}

// Settings and other values that affect how the server should interface with
// the database.
type DatabaseConfig struct {
	// Network address of the database.
	Host string
	// Name of the server's database.
	Name string
	// Password of the server's database user.
	Password string
	// TCP port of the database.
	Port int
	// Name of the server's database user.
	User string
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

			// Database flags.
			&cli.StringFlag{
				EnvVars: []string{"DB_HOST"},
				Name:    dbHostFlagName,
				Usage:   "Network address of the database",
				Value:   "postgres",
			},
			&cli.StringFlag{
				EnvVars: []string{"DB_NAME"},
				Name:    dbNameFlagName,
				Usage:   "Name of the server's database",
				Value:   "chorro",
			},
			&cli.StringFlag{
				EnvVars: []string{"DB_PASS"},
				Name:    dbPassFlagName,
				Usage:   "Password of the server's database user",
			},
			&cli.IntFlag{
				EnvVars: []string{"DB_PORT"},
				Name:    dbPortFlagName,
				Usage:   "TCP port of the database",
				Value:   5432,
			},
			&cli.StringFlag{
				EnvVars: []string{"DB_USER"},
				Name:    dbUserFlagName,
				Usage:   "Name of the server's database user",
				Value:   "chorro",
			},

			// General flags.
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

	// Database flag names.
	dbHostFlagName = "db-host"
	dbNameFlagName = "db-name"
	dbPassFlagName = "db-pass"
	dbPortFlagName = "db-port"
	dbUserFlagName = "db-user"

	// General flag names.
	environmentFlagName = "environment"
	portFlagName        = "port"
)

// Initializes this Config by reading values from the provided cli.Context.
func (config *Config) init(c *cli.Context) error {
	if err := config.Auth.init(c); err != nil {
		return err
	}

	if err := config.Database.init(c); err != nil {
		return err
	}

	config.Environment = toEnvironment(c.String(environmentFlagName))

	if config.HttpPort = c.Int(portFlagName); !isValidPort(config.HttpPort) {
		return fmt.Errorf("%d is not a valid %s", config.HttpPort, portFlagName)
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

// Initializes this DatabaseConfig by reading values from the provided
// cli.Context.
func (config *DatabaseConfig) init(c *cli.Context) error {
	if config.Host = c.String(dbHostFlagName); len(config.Host) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.Host, dbHostFlagName)
	}

	if config.Name = c.String(dbNameFlagName); len(config.Name) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.Name, dbNameFlagName)
	}

	config.Password = c.String(dbNameFlagName)

	if config.Name = c.String(dbNameFlagName); len(config.Name) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.Name, dbNameFlagName)
	}

	if config.Port = c.Int(dbPortFlagName); !isValidPort(config.Port) {
		return fmt.Errorf("%d is not a valid %s", config.Port, dbPortFlagName)
	}

	if config.User = c.String(dbUserFlagName); len(config.User) < 1 {
		return fmt.Errorf("\"%s\" is not a valid %s", config.User, dbUserFlagName)
	}

	return nil
}

// Returns true if port is valid.
func isValidPort(port int) bool {
	return port > 1024 && port < 99999
}
