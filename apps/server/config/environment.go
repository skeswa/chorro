package config

import "strings"

// Broadly describes the circumstances under which the server is running.
type Environment int

const (
	// `Environment` used during local development.
	DevEnvironment Environment = iota
	// `Environment` used when deployed into a production setting.
	ProdEnvironment
)

// Interprets `environmentString` as an `Environment`.
func toEnvironment(environmentString string) Environment {
	santizedEnvironmentString := strings.ToLower(strings.TrimSpace(environmentString))

	if santizedEnvironmentString == "prod" || santizedEnvironmentString == "production" {
		return ProdEnvironment
	}

	return DevEnvironment
}
