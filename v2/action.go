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
	"github.com/juicycleff/smartform/v1"
	"github.com/wakflo/go-sdk/v2/context"
	"github.com/wakflo/go-sdk/v2/core"
)

// ActionMetadata contains metadata about an action.
type ActionMetadata struct {
	// ID is the unique identifier for this action within its integration
	ID string `json:"id"`

	// DisplayName is the human-readable name of the action
	DisplayName string `json:"displayName"`

	// Description provides details about the action's purpose
	Description string `json:"description"`

	// HelpText provides additional guidance for configuring the action
	HelpText string `json:"helpText,omitempty"`

	// Icon is a URL or base64-encoded image for the action icon
	Icon string `json:"icon,omitempty"`

	// Type specifies the action type (e.g., ACTION, BRANCH, BOOLEAN)
	Type core.ActionType `json:"type"`

	// Documentation provides comprehensive usage instructions
	Documentation string `json:"documentation,omitempty"`

	// SampleOutput contains an example of the action's output
	SampleOutput core.JSON `json:"sampleOutput,omitempty"`

	// Tags are searchable labels for the action
	Tags []string `json:"tags,omitempty"`

	// Category provides a categorization for this action
	Category string `json:"category,omitempty"`

	// Settings returns action-specific settings
	Settings core.ActionSettings
}

// ActionDefinition contains metadata about an action.
type ActionDefinition struct {
	// Name is the unique identifier for this action within its integration
	Name string `json:"id"`

	// DisplayName is the human-readable name of the action
	DisplayName string `json:"displayName"`

	// Description provides details about the action's purpose
	Description string `json:"description"`

	// HelpText provides additional guidance for configuring the action
	HelpText string `json:"helpText,omitempty"`

	// Icon is a URL or base64-encoded image for the action icon
	Icon string `json:"icon,omitempty"`

	// Type specifies the action type (e.g., ACTION, BRANCH, BOOLEAN)
	Type core.ActionType `json:"type"`

	// Auth represents the authentication configuration required to perform the action, encapsulated in core.AuthMetadata.
	Auth *core.AuthMetadata

	// Documentation provides comprehensive usage instructions
	Documentation string `json:"documentation,omitempty"`

	// SampleOutput contains an example of the action's output
	SampleOutput core.JSON `json:"sampleOutput,omitempty"`

	Properties *smartform.FormSchema

	// Tags are searchable labels for the action
	Tags []string `json:"tags,omitempty"`

	// Implementation specifies the action implementation logic or function to be executed.
	Implementation Action `json:"-"`
}

// ActionError represents a specific error that can occur during action execution.
type ActionError struct {
	// Code is a machine-readable error code
	Code string `json:"code"`

	// Message is a human-readable error message
	Message string `json:"message"`

	// Retryable indicates if the error is temporary and the action can be retried
	Retryable bool `json:"retryable"`

	// Details contains additional context for the error
	Details map[string]interface{} `json:"details,omitempty"`
}

// Action defines the interface for workflow actions.
type Action interface {
	// Metadata returns metadata about the action
	Metadata() ActionMetadata

	// Properties returns the schema for the action's input configuration
	Properties() *smartform.FormSchema

	// Auth returns the authentication requirements for the action
	Auth() *core.AuthMetadata

	// Perform executes the action with the given context and input
	Perform(ctx context.PerformContext) (core.JSON, error)
}

// ActionRegistry manages action registration and discovery.
type ActionRegistry interface {
	// RegisterAction adds an action to the registry
	RegisterAction(integrationID string, action Action) error

	// UnregisterAction removes an action from the registry
	UnregisterAction(integrationID string, actionID string) error

	// GetAction retrieves an action by integration ID and action ID
	GetAction(integrationID string, actionID string) (Action, error)

	// ListActions returns all registered actions
	ListActions() map[string][]Action

	// ListActionsByType returns actions of a specific type
	ListActionsByType(actionType core.ActionType) []Action

	// ListActionsByIntegration returns all actions for a specific integration
	ListActionsByIntegration(integrationID string) []Action

	// ListActionsByCategory returns actions in a specific category
	ListActionsByCategory(category string) []Action

	// SearchActions searches for actions matching criteria
	SearchActions(query string, filters map[string]interface{}) []Action
}

// ActionFactory creates instances of actions.
type ActionFactory interface {
	// CreateAction creates an action instance based on type
	CreateAction(actionType core.ActionType, config map[string]interface{}) (Action, error)

	// RegisterActionCreator registers a function to create a specific action type
	RegisterActionCreator(actionType core.ActionType, creator func(config map[string]interface{}) (Action, error)) error

	// ListSupportedActionTypes returns all action types supported by this factory
	ListSupportedActionTypes() []core.ActionType
}
