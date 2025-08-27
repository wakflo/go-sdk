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

package context

import (
	"context"
	"time"

	"github.com/juicycleff/smartform/v1"
	"github.com/rs/xid"
	"github.com/wakflo/go-sdk/v2/core"
)

type BaseContext interface {
	// Context returns the underlying Go context for the execution operation.
	Context() context.Context

	// Logger returns a structured logger for the execution.
	Logger() core.Logger

	// Input returns the validated input data for the trigger execution.
	Input() core.JSONObject

	// AuthContext provides authentication context for the trigger execution.
	AuthContext() (*AuthContext, error)

	// Auth provides authentication context for the trigger execution.
	Auth() *AuthContext
}

// DynamicFieldContext defines the interface for performing an action in a workflow.
// It provides methods for handling action execution, including input validation,
// authentication, output processing, and error handling.
type DynamicFieldContext interface {
	BaseContext

	// Respond sends a response containing the provided data and total count of items, adhering to dynamic options structure.
	// Returns a `DynamicOptionsResponse` object with items and metadata or an error if processing the response fails.
	Respond(data any, totalItems int) (*core.DynamicOptionsResponse, error)

	// RespondJSON creates a JSON response containing the provided data and total item count. Returns the JSON or an error.
	RespondJSON(data any, totalItems int) (core.JSON, error)

	// WorkflowID returns the unique identifier of the workflow.
	WorkflowID() xid.ID

	// WorkflowVersionID returns the unique identifier of the workflow version.
	WorkflowVersionID() xid.ID

	// FieldName returns the name of the current field within the workflow context.
	FieldName() string

	// OperationID returns the unique identifier of the current operation within the workflow context.
	OperationID() string

	// StepID returns the unique identifier of the step in the workflow.
	StepID() string

	// Filter returns filtering parameters for dynamic options, including offset, limit, and filter term.
	Filter() *core.DynamicOptionsFilterParams
}

// LifecycleContext defines the interface for trigger lifecycle management.
// It provides methods for handling the setup, teardown, and lifecycle state
// of triggers within a workflow.
type LifecycleContext interface {
	// Context returns the underlying Go context for the lifecycle operation.
	Context() context.Context

	// WorkflowID returns the unique identifier of the workflow.
	WorkflowID() xid.ID

	// ProjectID returns the unique identifier of the project.
	ProjectID() xid.ID

	// TriggerID returns the unique identifier of the trigger being managed.
	TriggerID() string

	// Logger returns a structured logger for the trigger.
	Logger() core.Logger

	// Config returns the configuration for the trigger.
	Config() map[string]interface{}

	// Input returns the Input for the trigger.
	Input() map[string]interface{}

	// GetLastRunTime returns the timestamp of the last successful run.
	GetLastRunTime() (*time.Time, error)

	// SetLastRunTime updates the timestamp of the last successful run.
	SetLastRunTime(time.Time) error

	// GetState retrieves stored state for stateful triggers.
	GetState() (map[string]interface{}, error)

	// SetState stores state for stateful triggers.
	SetState(map[string]interface{}) error

	// TriggerCriteria returns the configured criteria for the trigger.
	TriggerCriteria() (*core.TriggerCriteria, error)

	// EmitEvent allows triggers to emit events that can be processed by the workflow runtime.
	EmitEvent(payload core.JSON) error

	// StoreMetadata allows triggers to store metadata about their execution.
	StoreMetadata(key string, value interface{}) error

	// GetMetadata retrieves stored metadata.
	GetMetadata(key string) (interface{}, error)

	// Cancel signals that the trigger should be canceled.
	Cancel() error

	// IsCanceled checks if the trigger has been canceled.
	IsCanceled() bool
}

// PerformContext defines the interface for performing an action in a workflow.
// It provides methods for handling action execution, including input validation,
// authentication, output processing, and error handling.
type PerformContext interface {
	BaseContext

	// WorkflowID returns the unique identifier of the workflow.
	WorkflowID() xid.ID

	// WorkflowVersionID returns the unique identifier of the workflow version.
	WorkflowVersionID() xid.ID

	// ProjectID returns the unique identifier of the project.
	ProjectID() xid.ID

	// StepID returns the unique identifier of the step in the workflow.
	StepID() string

	// RunID returns the unique identifier of the current workflow run.
	RunID() xid.ID

	// StepRunID returns the unique identifier of the current step run.
	StepRunID() xid.ID

	// PreviousStepOutput returns the output from the previous step in the workflow.
	PreviousStepOutput() (core.JSONObject, error)

	// Validate validates the input against the action's schema.
	Validate() error

	// SetOutput sets the output data from the action.
	SetOutput(output core.JSON) error

	// PauseExecution pauses the workflow execution, to be resumed later.
	PauseExecution(reason string, resumeAfter *time.Time) error

	// SetMetadata stores execution metadata for the workflow run.
	SetMetadata(key string, value interface{}) error

	// GetMetadata retrieves stored execution metadata.
	GetMetadata(key string) (interface{}, error)

	// ExecutionState returns the current state of the execution.
	ExecutionState() core.StepRunStatus

	// Schema returns the input schema for the action.
	Schema() *smartform.FormSchema

	// Retry schedules the action to be retried after the specified duration.
	Retry(after time.Duration, reason string) error

	// MarkFailed marks the action as failed with a reason.
	MarkFailed(reason string) error

	// Cancel signals that the action should be canceled.
	Cancel() error

	// IsCanceled checks if the action has been canceled.
	IsCanceled() bool

	// WorkflowContextData returns the workflow context data for the current run.
	WorkflowContextData() (map[string]interface{}, error)

	// UpdateWorkflowContext updates the workflow context data.
	UpdateWorkflowContext(data map[string]interface{}) error
}

// ExecuteContext defines the interface for executing a trigger in a workflow.
// It provides methods for handling trigger execution, including input validation,
// authentication, and output processing.
type ExecuteContext interface {
	BaseContext

	// WorkflowID returns the unique identifier of the workflow.
	WorkflowID() xid.ID

	// WorkflowVersionID returns the unique identifier of the workflow version.
	WorkflowVersionID() xid.ID

	// ProjectID returns the unique identifier of the project.
	ProjectID() xid.ID

	// TriggerID returns the unique identifier of the trigger being executed.
	TriggerID() string

	// RunID returns the unique identifier of the current workflow run.
	RunID() xid.ID

	// Logger returns a structured logger for the execution.
	Logger() core.Logger

	// Logger returns a structured logger for the execution.
	LastRun() *time.Time

	// SetInput updates the input data for the trigger execution.
	SetInput(input core.JSONObject) error

	// Environment returns the environment for the trigger execution.
	Environment() core.Environment

	// Validate validates the input against the trigger's schema.
	Validate() error

	// EmitEvent allows triggers to emit events during execution.
	EmitEvent(eventType string, payload core.JSON) error

	// SetOutput sets the output data from the trigger execution.
	SetOutput(output core.JSON) error

	// PauseExecution pauses the workflow execution, to be resumed later.
	PauseExecution(reason string) error

	// SetMetadata stores execution metadata for the workflow run.
	SetMetadata(key string, value interface{}) error

	// GetMetadata retrieves stored execution metadata.
	GetMetadata(key string) (interface{}, error)

	// ExecutionState returns the current state of the execution.
	ExecutionState() core.StepRunStatus

	// Schema returns the input schema for the trigger.
	Schema() *smartform.FormSchema

	// Cancel signals that the execution should be canceled.
	Cancel() error

	// IsCanceled checks if the execution has been canceled.
	IsCanceled() bool
}
