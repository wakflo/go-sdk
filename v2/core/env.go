package core

import (
	"fmt"
)

// Environment represents the execution environment
type Environment string

const (
	// EnvironmentTest represents the test environment
	EnvironmentTest Environment = "test"
	// EnvironmentDebug represents the debug environment
	EnvironmentDebug Environment = "debug"
	// EnvironmentProd represents the production environment
	EnvironmentProd Environment = "prod"
)

// String returns the string representation of the environment
func (e Environment) String() string {
	return string(e)
}

// IsValid checks if the environment is valid
func (e Environment) IsValid() bool {
	switch e {
	case EnvironmentTest, EnvironmentDebug, EnvironmentProd:
		return true
	default:
		return false
	}
}

func (e Environment) Values() (kinds []string) {
	for _, s := range []Environment{EnvironmentTest, EnvironmentDebug, EnvironmentProd} {
		kinds = append(kinds, string(s))
	}
	return
}

// ParseEnvironment parses a string into an Environment
func ParseEnvironment(s string) (Environment, error) {
	env := Environment(s)
	if !env.IsValid() {
		return "", fmt.Errorf("invalid environment: %s", s)
	}
	return env, nil
}

// IsProduction returns true if the environment is production
func (e Environment) IsProduction() bool {
	return e == EnvironmentProd
}

// IsTest returns true if the environment is test
func (e Environment) IsTest() bool {
	return e == EnvironmentTest
}

// IsDebug returns true if the environment is debug
func (e Environment) IsDebug() bool {
	return e == EnvironmentDebug
}
