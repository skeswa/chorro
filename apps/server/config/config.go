package config

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Settings and other values that affect how the server should function.
type Config struct {
	// Broadly describes the circumstances under which the server is running.
	Environment Environment
	// Port over which HTTP requests should be serviced.
	HttpPort int
}

// Reads `Config` from the environment and command line arguments.
func ReadConfig() (Config, error) {
	config := Config{}

	if err := (&cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Name:    "port",
				Usage:   "Port over which HTTP requests should be serviced",
				Value:   8000,
			},
			&cli.StringFlag{
				Aliases: []string{"e"},
				EnvVars: []string{"NODE_ENV"},
				Name:    "environment",
				Usage:   "Broadly describes the circumstances under which the server is running",
				Value:   "dev",
			},
		},
		Action: func(c *cli.Context) error {
			config.HttpPort = c.Int("port")
			if config.HttpPort <= 1024 || config.HttpPort > 10000 {
				return fmt.Errorf("%d is not a valid port", config.HttpPort)
			}

			config.Environment = toEnvironment(c.String("environment"))

			return nil
		},
	}).Run(os.Args); err != nil {
		return config, err
	}

	return config, nil
}
