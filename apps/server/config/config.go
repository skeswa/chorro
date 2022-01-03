package config

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Settings and other values that affect how the server should function.
type Config struct {
	// Settings and other values that affect how the server should interface with
	// the database.
	Database DatabaseConfig
	// Broadly describes the circumstances under which the server is running.
	Environment Environment
	// Port over which HTTP requests should be serviced.
	HttpPort int
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
	dbHostFlagName      = "db-host"
	dbNameFlagName      = "db-name"
	dbPassFlagName      = "db-pass"
	dbPortFlagName      = "db-port"
	dbUserFlagName      = "db-user"
	environmentFlagName = "environment"
	portFlagName        = "port"
)

// Initializes this Config by reading values from the provided cli.Context.
func (config *Config) init(c *cli.Context) error {
	if err := config.Database.init(c); err != nil {
		return err
	}

	config.Environment = toEnvironment(c.String(environmentFlagName))

	if config.HttpPort = c.Int(portFlagName); !isValidPort(config.HttpPort) {
		return fmt.Errorf("%d is not a valid %s", config.HttpPort, portFlagName)
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
