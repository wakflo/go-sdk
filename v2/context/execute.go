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
	"time"

	"github.com/juicycleff/smartform/v1"
	"github.com/rs/xid"
	"github.com/wakflo/go-sdk/v2/core"
)

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
