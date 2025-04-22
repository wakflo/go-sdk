// Copyright 2022-present Wakflo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdk

import (
	"context"

	"github.com/wakflo/go-sdk/v2/core"
)

// IntegrationType categorizes the integration functionality.
type IntegrationType string

const (
	// IntegrationTypeData indicates an integration focused on data operations.
	IntegrationTypeData IntegrationType = "data"

	// IntegrationTypeMessaging indicates an integration focused on messaging.
	IntegrationTypeMessaging IntegrationType = "messaging"

	// IntegrationTypeStorage indicates an integration focused on storage.
	IntegrationTypeStorage IntegrationType = "storage"

	// IntegrationTypeAnalytics indicates an integration focused on analytics.
	IntegrationTypeAnalytics IntegrationType = "analytics"

	// IntegrationTypeUtility indicates an integration providing utility functions.
	IntegrationTypeUtility IntegrationType = "utility"

	// IntegrationTypeFlow indicates an integration providing flow control.
	IntegrationTypeFlow IntegrationType = "flow"

	// IntegrationTypeAI indicates an integration providing AI/ML capabilities.
	IntegrationTypeAI IntegrationType = "ai"
)

// IntegrationDefinition represents the definition of an integration
type IntegrationDefinition struct {
	ID            string                        `json:"id"`
	Name          string                        `json:"name"`
	DisplayName   string                        `json:"displayName"`
	Description   string                        `json:"description"`
	Version       string                        `json:"version"`
	Icon          string                        `json:"icon"`
	Actions       map[string]*ActionDefinition  `json:"actions"`
	Triggers      map[string]*TriggerDefinition `json:"triggers"`
	Auth          core.AuthMetadata             `json:"auth"`
	Metadata      IntegrationMetadata           `json:"metadata"`
	BuildMetadata core.IntegrationBuildMetadata `json:"buildMetadata"`
	Documentation string                        `json:"documentation,omitempty"`
	ReleaseNotes  string                        `json:"releaseNotes,omitempty"`
	Type          IntegrationType               `json:"type"`
	Categories    []string                      `json:"categories"`
	Tags          []string                      `json:"tags"`
	Authors       []string                      `json:"authors"`
	Website       *string                       `json:"website"`
	License       *string                       `json:"license"`
	Copyright     *string                       `json:"copyright"`
	LicenseURL    *string                       `json:"licenseUrl"`
	CopyrightURL  *string                       `json:"copyrightUrl"`
	Source        *string                       `json:"source"`

	Implementation Integration `json:"-"`
}

// IntegrationMetadata contains metadata about an integration.
type IntegrationMetadata struct {
	// Name is the human-readable name of the integration
	Name string `json:"name" toml:"name" yaml:"name" validate:"required"`

	// Description provides details about the integration's purpose
	Description string `json:"description" toml:"description"  yaml:"description" validate:"required"`

	// Type categorizes the integration functionality
	Type IntegrationType `json:"type"`

	// Version is the semantic version of the integration
	Version string `json:"version" toml:"version" yaml:"version" validate:"required"`

	// Icon is a URL or base64-encoded image for the integration icon
	Icon string `json:"icon" toml:"icon" yaml:"icon" validate:"required"`

	// Publisher identifies who created the integration
	Authors []string `json:"authors" toml:"authors" yaml:"authors" validate:"required"`

	// Website is the URL to the integration's documentation or website
	Website string `json:"website" toml:"website" yaml:"website"`

	// Category provides a more specific categorization
	Categories []string `json:"categories" toml:"categories" yaml:"categories" validate:"required"`

	// Tags are searchable labels for the integration
	Tags []string `json:"tags,omitempty" toml:"tags,omitempty" yaml:"tags,omitempty"`

	// ReleaseNotes documents changes in this version
	ReleaseNotes string `json:"releaseNotes,omitempty"`

	// Documentation provides comprehensive usage instructions
	Documentation string `json:"documentation,omitempty"`
}

// Integration defines the interface for integration plugins.
type Integration interface {
	// Metadata returns metadata about the integration
	Metadata() IntegrationMetadata

	// Auth returns the authentication requirements for the integration
	Auth() *core.AuthMetadata

	// Triggers returns all triggers provided by this integration
	Triggers() []Trigger

	// Actions returns all actions provided by this integration
	Actions() []Action
}

// IntegrationRegistry manages available integrations.
type IntegrationRegistry interface {
	// RegisterIntegration adds an integration to the registry
	RegisterIntegration(integration Integration) error

	// RegisterIntegrationDefinition adds an integration to the registry
	RegisterIntegrationDefinition(integration IntegrationDefinition) error

	// UnregisterIntegration removes an integration from the registry
	UnregisterIntegration(name string, version string) error

	// GetIntegration retrieves a specific version of an integration by its name and version from the registry.
	GetIntegration(ctx context.Context, name string, version string) (Integration, error)

	// GetIntegrationDefinition retrieves a specific version of an integration by its name and version from the registry.
	GetIntegrationDefinition(ctx context.Context, name string, version string) (*IntegrationDefinition, error)

	// GetLatestIntegration retrieves the latest version of an integration
	GetLatestIntegration(ctx context.Context, name string) (*IntegrationDefinition, error)

	// ListIntegrations returns all registered integrations
	ListIntegrations() []*IntegrationDefinition

	// ListIntegrationsByType returns integrations of a specific type
	ListIntegrationsByType(integrationType IntegrationType) []*IntegrationDefinition

	// ListIntegrationVersions returns all versions of an integration
	ListIntegrationVersions(ctx context.Context, name string) ([]*IntegrationDefinition, error)

	// GetIntegrationMetadata retrieves metadata for a specific integration given its name and version. It returns the metadata or an error.
	GetIntegrationMetadata(ctx context.Context, name string, version string) (*IntegrationMetadata, error)

	RegisterActionDefinition(
		ctx context.Context,
		integrationName string,
		version string,
		actionID string,
		action ActionDefinition,
	) error

	RegisterTriggerDefinition(
		ctx context.Context,
		integrationName string,
		version string,
		triggerID string,
		trigger TriggerDefinition,
	) error

	RegisterAction(
		ctx context.Context,
		integrationName string,
		version string,
		actionID string,
		action Action,
	) error

	RegisterTrigger(
		ctx context.Context,
		integrationName string,
		version string,
		triggerID string,
		trigger Trigger,
	) error

	// GetAction retrieves an action by integration ID and action ID
	GetAction(integrationID, version, actionID string) (*ActionDefinition, error)

	// GetTrigger retrieves a trigger by integration ID and trigger ID
	GetTrigger(integrationID, version, actionID string) (*TriggerDefinition, error)

	// ListActions returns all actions for an integration version
	ListActions(
		ctx context.Context,
		integrationName string,
		version string,
	) (map[string]*ActionDefinition, error)

	// ListTriggers returns all triggers for an integration version
	ListTriggers(
		ctx context.Context,
		integrationName string,
		version string,
	) (map[string]*TriggerDefinition, error)

	// ListAllActions returns all registered actions across all integrations
	ListAllActions(ctx context.Context) map[string]map[string]map[string]*ActionDefinition

	// ListAllTriggers returns all registered triggers across all integrations
	ListAllTriggers(ctx context.Context) map[string]map[string]map[string]*TriggerDefinition

	// Initialize initializes all registered integrations
	Initialize(ctx context.Context) error

	// Shutdown performs cleanup for all registered integrations
	Shutdown(ctx context.Context) error
}

// DynamicIntegrationLoader loads integrations from external sources.
type DynamicIntegrationLoader interface {
	// LoadIntegration loads an integration from a file or URL
	LoadIntegration(ctx context.Context, source string) (*IntegrationDefinition, error)

	// ValidateIntegration checks if an integration package is valid
	ValidateIntegration(ctx context.Context, source string) (bool, error)

	// GetIntegrationDefinition extracts metadata from an integration package
	GetIntegrationDefinition(ctx context.Context, source string) (*IntegrationDefinition, error)

	// GetIntegrationMetadata extracts metadata from an integration package
	GetIntegrationMetadata(ctx context.Context, source string) (*IntegrationMetadata, error)

	// ExtractIntegration extracts and validates integration data from the given path and returns its identifier or an error.
	ExtractIntegration(ctx context.Context, path string) (string, error)

	// RegisterIntegration loads and registers an integration with the registry
	RegisterIntegration(ctx context.Context, source string, registry IntegrationRegistry) error
}
