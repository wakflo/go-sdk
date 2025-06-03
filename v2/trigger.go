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

// TriggerMetadata contains metadata about a trigger.
type TriggerMetadata struct {
	// ID is the unique identifier for this trigger within its integration
	ID string `json:"id"`

	// DisplayName is the human-readable name of the trigger
	DisplayName string `json:"displayName"`

	// Description provides details about the trigger's purpose
	Description string `json:"description"`

	// HelpText provides additional guidance for configuring the trigger
	HelpText string `json:"helpText,omitempty"`

	// Icon is a URL or base64-encoded image for the trigger icon
	Icon string `json:"icon,omitempty"`

	// Type specifies the trigger type (e.g., POLLING, WEBHOOK)
	Type core.TriggerType `json:"type"`

	// Documentation provides comprehensive usage instructions
	Documentation string `json:"documentation,omitempty"`

	// SampleOutput contains an example of the trigger's output
	SampleOutput core.JSON `json:"sampleOutput,omitempty"`

	// Criteria returns additional trigger criteria configuration
	Criteria *core.TriggerCriteria `json:"criteria"`
}

// TriggerDefinition contains metadata about a trigger.
type TriggerDefinition struct {
	// ID is the unique identifier for this trigger within its integration
	Name string `json:"name"`

	// DisplayName is the human-readable name of the trigger
	DisplayName string `json:"displayName"`

	// Description provides details about the trigger's purpose
	Description string `json:"description"`

	// HelpText provides additional guidance for configuring the trigger
	HelpText string `json:"helpText,omitempty"`

	// Icon is a URL or base64-encoded image for the trigger icon
	Icon string `json:"icon,omitempty"`

	// Type specifies the trigger type (e.g., POLLING, WEBHOOK)
	Type core.TriggerType `json:"type"`

	// Auth represents the authentication configuration required to perform the action, encapsulated in core.AuthMetadata.
	Auth *core.AuthMetadata `json:"auth"`

	// Documentation provides comprehensive usage instructions
	Documentation string `json:"documentation,omitempty"`

	// SampleOutput contains an example of the trigger's output
	SampleOutput core.JSON `json:"sampleOutput,omitempty"`

	// Properties defines the schema for additional configuration required for the trigger in the form of a smartform.
	Properties *smartform.FormSchema `json:"properties"`

	// Settings provides a settings for this trigger
	Settings core.TriggerSettings `json:"settings"`

	// Implementation specifies the action implementation logic or function to be executed.
	Implementation Trigger `json:"-"`
}

// Trigger defines the interface for workflow triggers.
type Trigger interface {
	// Metadata returns metadata about the trigger
	Metadata() TriggerMetadata

	// Props returns the schema for the trigger's input configuration
	Props() *smartform.FormSchema

	// Auth returns the authentication requirements for the trigger
	Auth() *core.AuthMetadata

	// Start prepares and activates the trigger (e.g., start polling, event listening, cron schedules)
	Start(ctx context.LifecycleContext) error

	// Stop gracefully stops or disables the trigger
	Stop(ctx context.LifecycleContext) error

	// Execute handles the trigger's action when manually invoked with an input schema
	Execute(ctx context.ExecuteContext) (core.JSON, error)
}

// TriggerRegistry manages trigger registration and discovery.
type TriggerRegistry interface {
	// RegisterTrigger adds a trigger to the registry
	RegisterTrigger(integrationID string, trigger Trigger) error

	// UnregisterTrigger removes a trigger from the registry
	UnregisterTrigger(integrationID string, triggerID string) error

	// GetTrigger retrieves a trigger by integration ID and trigger ID
	GetTrigger(integrationID string, triggerID string) (Trigger, error)

	// ListTriggers returns all registered triggers
	ListTriggers() map[string][]Trigger

	// ListTriggersByType returns triggers of a specific type
	ListTriggersByType(triggerType core.TriggerType) []Trigger

	// ListTriggersByIntegration returns all triggers for a specific integration
	ListTriggersByIntegration(integrationID string) []Trigger
}

// TriggerFactory creates instances of triggers.
type TriggerFactory interface {
	// CreateTrigger creates a trigger instance based on type
	CreateTrigger(triggerType core.TriggerType, config map[string]interface{}) (Trigger, error)

	// RegisterTriggerCreator registers a function to create a specific trigger type
	RegisterTriggerCreator(triggerType core.TriggerType, creator func(config map[string]interface{}) (Trigger, error)) error

	// ListSupportedTriggerTypes returns all trigger types supported by this factory
	ListSupportedTriggerTypes() []core.TriggerType
}
