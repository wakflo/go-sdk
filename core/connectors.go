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

package core

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/wakflo/go-sdk/core/authenums"
)

// JSONObject is a type alias for map[string]any.
type JSONObject = map[string]any

// ToJSONMap converts an interface of type `any` to a map[string]any.
// Returns the map and a boolean indicating success.
func ToJSONMap(input any) (JSONObject, bool) {
	if input == nil {
		return nil, false
	}

	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, false
	}

	var result map[string]any
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, false
	}

	return result, true
}

// ToJSON converts an interface of type `any` to a map[string]any.
// Returns the map and a boolean indicating success.
func ToJSON(input any) (JSON, bool) {
	if input == nil {
		return nil, false
	}

	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, false
	}

	var result any
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, false
	}

	return result, true
}

type FormState struct {
	Pristine      bool           `json:"pristine"`
	Dirty         bool           `json:"dirty"`
	Disabled      bool           `json:"disabled"`
	Submitted     bool           `json:"submitted"`
	Valid         bool           `json:"valid"`
	Invalid       bool           `json:"invalid"`
	Submitting    bool           `json:"submitting"`
	Validating    int            `json:"validating"`
	Gathering     int            `json:"gathering"`
	Values        map[string]any `json:"values"`
	MaskedValues  map[string]any `json:"maskedValues"`
	Errors        map[string]any `json:"errors"`
	Touched       map[string]any `json:"touched"`
	Modified      map[string]any `json:"modified"`
	Dirt          map[string]any `json:"dirt"`
	Focused       map[string]any `json:"focused"`
	InitialValues map[string]any `json:"initialValues"`
	Data          map[string]any `json:"data"`
	Memory        map[string]any `json:"memory"`
}

// ConnectorProperties is a Task operation type.
type ConnectorProperties struct {
	FormState      FormState   `json:"formState"`
	Input          JSONObject  `json:"input"`
	Output         any         `json:"output"`
	LastTestTime   *int        `json:"lastTestTime"`
	LastTestStatus *TestStatus `json:"lastTestStatus"`
}

// ConnectorAuthentication is a Task operation type.
type ConnectorAuthentication struct {
	AuthType string `json:"authType"`
}

// TaskOperation is a Task operation type.
type TaskOperation struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Input       map[string]any `json:"input"`
	Output      map[string]any `json:"output"`
}

// ConnectorPlugin is a ConnectorPlugin model.
type ConnectorPlugin struct {
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "deleted" field.
	Description string `json:"description"`
	// Version holds the value of the "version" field.
	Version string `json:"version"`
	// TeamType holds the value of the "language" field.
	Language PluginLanguage `json:"language,omitempty"`
}

// WorkspacePluginMetadata is a WorkspacePluginMetadata model.
type WorkspacePluginMetadata struct {
	Operations []TaskOperation `json:"properties,omitempty"`
	Compiler   PluginCompiler  `json:"compiler"`
	Language   PluginLanguage  `json:"language,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "deleted" field.
	Description string `json:"description"`
	// Documentation holds the value of the "documentation" field.
	Documentation *string `json:"documentation"`
	// Description holds the value of the "deleted" field.
	Category string `json:"category"`
	// Version holds the value of the "version" field.
	Version string `json:"version"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon"`
}

// PluginMetadata is a PluginMetadata model.
type PluginMetadata struct {
	Compiler PluginCompiler `json:"compiler"`
	Language PluginLanguage `json:"language,omitempty"`
}

// AuthOperation is an auth operation struct.
type AuthOperation struct {
	Type   authenums.AuthType `json:"type"`
	Config map[string]any     `json:"config"`
}

// Operation .
type Operation struct {
	// Key holds the value of the "key" field.
	Name string `json:"key,omitempty"`
	Icon string `json:"icon,omitempty"`
	// Name holds the value of the "name" field.
	DisplayName string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// HelpText holds the value of the "helpText" field.
	HelpText *string `json:"helpText,omitempty"`
	// Input holds the value of the "input" field.
	Input *AutoFormSchema `json:"input,omitempty"`
	// Auth holds the value of the "auth" field.
	Auth *AutoFormSchema `json:"auth,omitempty"`
	// Output holds the value of the "output" field.
	Output map[string]any `json:"output,omitempty"`
	// SampleOutput holds the value of the "sample_output" field.
	SampleOutput map[string]any `json:"sampleOutput,omitempty"`

	ErrorSettings StepErrorSettings `json:"errorSettings,omitempty"`

	Settings OperationSettings `json:"settings,omitempty"`

	RequireAuth bool `json:"requireAuth"`
	// Documentation represents the field used to store the connector's documentation in markdown.
	Documentation *string `json:"documentation,omitempty"`

	Type ActionType `json:"type" validate:"required,oneof=ACTION"`
}

type Operations = map[string]*Operation

type OperationsList = []*Operation

// Trigger .
type Trigger struct {
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" validate:"required"`
	Icon string `json:"icon,omitempty"`
	// DisplayName holds the value of the "displayName" field.
	DisplayName string `json:"displayName,omitempty" validate:"required"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty" validate:"required"`
	// HelpText holds the value of the "helpText" field.
	HelpText *string `json:"helpText,omitempty"`
	// Input holds the value of the "input" field.
	Input *AutoFormSchema `json:"input,omitempty"`
	// SampleOutput holds the value of the "sampleOutput" field.
	SampleOutput map[string]any `json:"sampleOutput,omitempty"`
	// Auth holds the value of the "auth" field.
	Auth *AutoFormSchema `json:"auth,omitempty"`

	// Type holds the value of the "description" field.
	Type TriggerType `json:"type,omitempty" validate:"required,oneof=SCHEDULED EVENT POLLING WEBHOOK"`

	Settings *TriggerSettings `json:"settings,omitempty"`

	ErrorSettings *StepErrorSettings `json:"errorSettings,omitempty"`

	RequireAuth bool `json:"requireAuth"`

	// Documentation represents the field used to store the connector's documentation in markdown.
	Documentation *string `json:"documentation,omitempty"`
}

type (
	Triggers     = map[string]*Trigger
	TriggersList = []*Trigger
)

type FlowRunMetadata struct {
	// FlowID holds the value of the "id" field.
	FlowID uuid.UUID `json:"flowId,omitempty"`
	// Name holds the value of the "name" field.
	FlowName string `json:"flowName,omitempty"`
	// StepName holds the value of the "name" field.
	StepName string `json:"stepName,omitempty"`
	// ConnectorName holds the value of the "connectorName" field.
	ConnectorName string `json:"connectorName,omitempty"`
	// ConnectorVersion holds the value of the "connectorVersion" field.
	ConnectorVersion string `json:"connectorVersion,omitempty"`
	// LastRun represents the timestamp of the last run of a flow.
	LastRun *time.Time `json:"lastRun"`
}

type OperationSettings struct {
	Branch *BranchSettings `json:"branch,omitempty"`
}

// ConnectorVersionMetadata is the model entity for the ConnectorVersion schema.
type ConnectorVersionMetadata struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// DeletedAt holds the value of the "delete_time" field.
	DeletedAt time.Time `json:"delete_time,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Auth holds the value of the "auth" field.
	Auth AuthOperation `json:"auth,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Namespace holds the value of the "namespace" field.
	RegistryName string `json:"registry_name,omitempty"`
	// Documentation holds the value of the "documentation" field.
	Documentation *string `json:"documentation,omitempty"`
	// ReleaseNotes holds the value of the "release_notes" field.
	ReleaseNotes *string `json:"release_notes,omitempty"`
	// ConnectorID holds the value of the "connector_id" field.
	ConnectorID uuid.UUID `json:"connector_id,omitempty"`
	// BuildMetadata holds the value of the "build_metadata" field.
	BuildMetadata *string `json:"build_metadata,omitempty"`
	// FileURL holds the value of the "file_url" field.
	FileURL *string `json:"file_url,omitempty"`
	// FileHash holds the value of the "file_hash" field.
	FileHash *string `json:"file_hash,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata PluginMetadata `json:"metadata,omitempty"`
	// Operations holds the value of the "operations" field.
	Operations map[string]*Operation `json:"operations,omitempty"`
	// Triggers holds the value of the "triggers" field.
	Triggers map[string]*Trigger `json:"triggers,omitempty"`
	// Approved holds the value of the "approved" field.
	Approved bool `json:"approved"`
}

type TestFlowStepInput struct {
	StepName string    `json:"stepName"`
	FlowID   uuid.UUID `json:"flowId"`
}

func FlattenSteps(node *FlowStep) []*FlowStep {
	nodes := []*FlowStep{node}
	if node.Children == nil {
		return nodes
	}

	for _, child := range node.Children {
		nodes = append(nodes, FlattenSteps(child)...)
	}
	return nodes
}
