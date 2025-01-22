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
)

type FlowSettings struct {
	Config  map[string]interface{} `json:"config"`
	LastRun *time.Time             `json:"LastRun"`
	// OutputSchema jsonschema4.JsonSchema4 `json:"output_schema"`
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
	TriggerStrategy  *TriggerType `json:"triggerStrategy,omitempty"`
}

type FlowStep struct {
	// Name of the schema
	Name string `json:"name,omitempty" validate:"required"`

	// Label of the step
	Label string `json:"label,omitempty" validate:"required"`

	// Icon of the step
	Icon string `json:"icon,omitempty" validate:"required"`

	// Icon of the step
	Type FlowStepType `json:"type,omitempty" validate:"required"`

	Meta map[string]any `json:"meta"`

	Settings StepNodeSettings `json:"settings"`

	OperationID *string `json:"operationId"`

	Auth *StepNodeNodeAuthInput `json:"auth"`

	Form StepNodeFormInput `json:"form"`

	Output any `json:"output"`

	SampleOutput any `json:"sampleOutput"`

	Tests StepNodeTest `json:"tests"`

	FirstLoopStep *FlowStep `json:"firstLoopStep,omitempty"`

	NextStep *FlowStep `json:"nextStep,omitempty"`

	// Data of the step
	Children []*FlowStep `json:"children,omitempty"`

	Valid bool `json:"valid,omitempty"`

	Skip bool `json:"skip,omitempty"`
}

func (s *FlowStep) Flatten() ([]FlowStep, error) {
	var dst FlowStep
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dst)
	return GetAllSteps(dst), nil
}

func (s *FlowStep) IsValid() bool {
	return s.Valid
}

func (s *FlowStep) IsTrigger() bool {
	return s.Type == FlowStepTypeStepTrigger
}

func (s *FlowStep) IsLoop() bool {
	return s.Type == FlowStepTypeLoop
}

func (s *FlowStep) IsRouter() bool {
	return s.Type == FlowStepTypeStepRouter
}

func (s *FlowStep) IsFlow() bool {
	return s.IsLoop() || s.IsRouter()
}

func (s *FlowStep) FlattenUnsafe() []FlowStep {
	var dst FlowStep
	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(b, &dst)
	return GetAllSteps(dst)
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

type TriggerFlowConfig struct {
	// FlowID of the trigger
	FlowID string `json:"flow_id,omitempty"`
}

type FlowVersion struct {
	ID uuid.UUID `json:"id,omitempty"`
	// FlowID holds the value of the "flow_id" field.
	FlowID uuid.UUID `json:"flow_id"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Steps holds the value of the "steps" field.
	Trigger FlowStep `json:"trigger,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID uuid.UUID `json:"project_id,omitempty"`
	// LastRun field stores the timestamp of the last run of a flow version.
	LastRun *time.Time `json:"last_run,omitempty"`
	// Status indicates the current state of the flow version, represented by the FlowVersionState enumeration.
	Status FlowVersionState `json:"status"`
}

type FlowMetadata struct {
	FlowVersionID uuid.UUID `json:"flow_version_id,omitempty"`
	// FlowID holds the value of the "flow_id" field.
	FlowID uuid.UUID `json:"flow_id"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LastRun field stores the timestamp of the last run of a flow version.
	LastRun *time.Time `json:"last_run,omitempty"`
	// Status indicates the current state of the flow version, represented by the FlowVersionState enumeration.
	Status FlowVersionState `json:"status"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID uuid.UUID `json:"project_id,omitempty"`
}

func (m FlowMetadata) ToSDK() {
}
