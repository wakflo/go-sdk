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
	"time"

	"github.com/google/uuid"
)

type WorkflowSettings struct {
	Config  map[string]interface{} `json:"config"`
	LastRun *time.Time             `json:"LastRun"`
	// OutputSchema jsonschema4.JsonSchema4 `json:"output_schema"`
}

// StepEdge is the graph edge reference of the steps
type StepEdge struct {
	// ID of the edge
	ID string `json:"id,omitempty"`

	// Source id of source step
	Source string `json:"source,omitempty"`

	// Target id of target step
	Target string `json:"target,omitempty"`

	// Type of step edge
	Type string `json:"type,omitempty"`

	// Style of step edge
	Style *map[string]interface{} `json:"style,omitempty"`
}

type EntityReference struct {
	// Field such is the unique field name of the entity
	Field string `json:"field,omitempty"`

	// Value such is the unique field value of the entity
	Value string `json:"value,omitempty"`

	// Value such is the unique field value of the entity
	Version *string `json:"version,omitempty"`

	// Type of the entity
	Entity string `json:"entity,omitempty"`
}

type ConnectorStepPosition struct {
	// Field of the step
	X int32 `json:"x"`

	Y int32 `json:"y"`
}

type ConnectorStepData struct {
	OperationID      *string             `json:"operationId"`
	AuthConnectionID *uuid.UUID          `json:"authConnectionId"`
	Properties       ConnectorProperties `json:"properties"`
}

type StepErrorSettings struct {
	// ContinueOnError of the step
	ContinueOnError bool `json:"continueOnError,omitempty"`

	// RetryOnError of the step
	RetryOnError bool `json:"retryOnError,omitempty"`
}

type ConnectorStepMetadata struct {
	ConnectorName    string       `json:"connectorName,omitempty"`
	ConnectorVersion string       `json:"connectorVersion,omitempty"`
	TriggerType      *TriggerType `json:"triggerType,omitempty"`
}

type ConnectorStep struct {
	// Label of the step
	Label string `json:"label,omitempty" validate:"required"`

	// Icon of the step
	Icon string `json:"icon,omitempty" validate:"required"`

	// Name of the schema
	Name string `json:"name,omitempty" validate:"required"`

	// IsTrigger of the schema
	IsTrigger bool `json:"isTrigger" validate:"required"`

	// Path of the schema
	Path []string `json:"path,omitempty" validate:"required"`

	NodeIndex int `json:"nodeIndex" validate:"required"`

	// Icon of the step
	Type StepType `json:"type,omitempty" validate:"required"`

	// Data of the step
	Data ConnectorStepData `json:"data,omitempty"`

	// Data of the step
	Children *[]ConnectorStep `json:"children,omitempty"`

	Reference *EntityReference `json:"reference,omitempty"`

	Metadata ConnectorStepMetadata `json:"metadata,omitempty" validate:"required"`

	TriggerSettings *StepTriggerSettings `json:"triggerSettings,omitempty"`

	// ParentID of the step
	ParentID *string `json:"parentId,omitempty"`

	// ErrorSettings of the step
	ErrorSettings StepErrorSettings `json:"errorSettings,omitempty"`

	// ContinueOnError of the step
	Valid bool `json:"valid,omitempty"`

	// IsFolded if step is folded
	IsFolded bool `json:"isFolded,omitempty"`
}

// SignLog sign log
type SignLog struct {
	UserAgent string
	At        *time.Time
	IP        string
}

// SignLogs record sign in logs
type SignLogs struct {
	Log         string
	SignInCount uint
	Logs        []SignLog
}

type TriggerPubSubConfig struct {
	// Topic of the pubsub
	Topic string `json:"id,omitempty"`
	// Topic of the pubsub
	Endpoint string `json:"endpoint,omitempty"`
}

type TriggerManualConfig struct{}

type TriggerScheduledConfig struct {
	// Field of the step
	RunAt string `json:"runAt,omitempty"`
}

type TriggerWebhookConfig struct {
	// Endpoint of the trigger
	Endpoint string `json:"endpoint,omitempty"`
}

type TriggerWorkflowConfig struct {
	// WorkflowID of the trigger
	WorkflowID string `json:"workflow_id,omitempty"`
}

type WorkflowVersion struct {
	ID uuid.UUID `json:"id,omitempty"`
	// WorkflowID holds the value of the "workflow_id" field.
	WorkflowID uuid.UUID `json:"workflow_id,omitempty"`
	// Name holds the value of the "workflow_id" field.
	Name string `json:"name,omitempty"`
	// Steps holds the value of the "steps" field.
	Steps map[string]ConnectorStep `json:"steps,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID uuid.UUID `json:"project_id,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// LastRun field stores the timestamp of the last run of a workflow version.
	LastRun *time.Time `json:"last_run,omitempty"`
}
