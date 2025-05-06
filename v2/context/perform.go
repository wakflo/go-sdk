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

	// AuthContext provides authentication context for the action.
	AuthContext() (*AuthContext, error)

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

	// WorkflowContext returns the workflow context data for the current run.
	WorkflowContext() (map[string]interface{}, error)

	// UpdateWorkflowContext updates the workflow context data.
	UpdateWorkflowContext(data map[string]interface{}) error
}
