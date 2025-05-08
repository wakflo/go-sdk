package sdk

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/wakflo/go-sdk/validator"
)

func Register(integration Integration) Integration {
	return integration
}

func registerIntegration(flow string, readme string) IntegrationMetadata {
	reg, err := registerInternal(flow, readme)
	if err != nil {
		panic(err)
	}
	return *reg
}

func registerInternal(flow string, readme string) (*IntegrationMetadata, error) {
	readme, err := ReadREADME(readme)
	if err != nil {
		return nil, fmt.Errorf("failed to read README.md: %w", err)
	}

	info, err := ReadFloFile(flow)
	if err != nil {
		return nil, err
	}

	// name := slug.Make(info.Name)
	// info.Name = name
	info.Documentation = readme

	return info, nil
}

type IntegrationSchemaModel struct {
	Name        string   `json:"name" toml:"name" yaml:"name" validate:"required"`
	Description string   `json:"description" toml:"description"  yaml:"description" validate:"required"`
	Version     string   `json:"version" toml:"version" yaml:"version" validate:"required"`
	Authors     []string `json:"authors" toml:"authors" yaml:"authors" validate:"required"`
	Website     *string  `json:"website" toml:"website" yaml:"website"`
	Categories  []string `json:"categories" toml:"categories" yaml:"categories" validate:"required"`
	Icon        string   `json:"icon" toml:"icon" yaml:"icon" validate:"required"`
}

type SchemaConfig struct {
	Integration IntegrationMetadata `json:"integration" toml:"integration" yaml:"integration" validate:"required"`
}

func ReadFloFile(content string) (*IntegrationMetadata, error) {
	// Deserialize the file content into SchemaConfig
	var config SchemaConfig
	err := toml.Unmarshal([]byte(content), &config)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize TOML: %w", err)
	}

	v := validator.NewDefaultValidator()

	err = v.Validate(&config)
	if err != nil {
		return nil, err
	}

	return &config.Integration, nil
}

// ReadREADME extracts the content of README.md from the current directory.
func ReadREADME(content string) (string, error) {
	// fileName := "README.md"
	// // Resolve the full path relative to the caller's directory
	// fullPath, err := GetRelativePathWithDepth(fileName, 4)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to resolve full path: %w", err)
	// }
	//
	// // Define the path to README.md
	// readmePath := filepath.Join(fullPath)
	//
	// // Check if the file exists
	// if _, err := os.Stat(readmePath); os.IsNotExist(err) {
	// 	return "", fmt.Errorf("README.md not found in the directory: %s", fullPath)
	// }
	//
	// // Read the file content
	// content, err := os.ReadFile(readmePath)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to read README.md: %w", err)
	// }

	// Trim any leading or trailing whitespace
	return strings.TrimSpace(string(content)), nil
}

// GetRelativePathWithDepth resolves a file path relative to the location of the caller, with adjustable depth.
func GetRelativePathWithDepth(relativePath string, depth int) (string, error) {
	// Get the caller's frame based on the specified depth
	_, callerFile, _, ok := runtime.Caller(depth)
	if !ok {
		return "", fmt.Errorf("unable to retrieve caller information")
	}

	// Get the directory of the caller
	callerDir := filepath.Dir(callerFile)

	// Resolve the full path relative to the caller's directory
	fullPath := filepath.Join(callerDir, relativePath)
	return fullPath, nil
}
