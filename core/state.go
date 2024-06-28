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

type (
	StepsState       struct{}
	StepsRunSnapshot struct {
		// Steps          map[string]*StepState `json:"steps"`
		CurrentStepID  string  `json:"currentStepId"`
		PreviousStepID *string `json:"previousStepId"`
	}
)

type GetDynamicOptionsInput struct {
	ConnectorName     string                 `json:"connectorName"`
	ConnectorVersion  string                 `json:"connectorVersion"`
	FieldName         string                 `json:"fieldName"`
	OperationName     string                 `json:"operationName"`
	StepName          string                 `json:"stepName"`
	Filter            *string                `json:"filter,omitempty"`
	WorkflowID        string                 `json:"workflowId"`
	WorkflowVersionID *string                `json:"workflowVersionId,omitempty"`
	Input             map[string]interface{} `json:"input,omitempty"`
}

type StepRunData struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID uuid.UUID `json:"team_id,omitempty"`
	// Status holds the value of the "status" field.
	Status StepRunStatus `json:"status,omitempty"`
	// JobID holds the value of the "job_id" field.
	JobID int `json:"job_id,omitempty"`
	// Index holds the value of the "index" field.
	Index int `json:"index,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// WorkflowID holds the value of the "workflow_id" field.
	WorkflowID uuid.UUID `json:"workflow_id,omitempty"`
	// WorkflowVersionID holds the value of the "workflow_version_id" field.
	WorkflowVersionID *uuid.UUID `json:"workflow_version_id,omitempty"`
	// WorkflowRunID holds the value of the "workflow_run_id" field.
	WorkflowRunID uuid.UUID `json:"workflow_run_id,omitempty"`
	// ConnectorName holds the value of the "connector_name" field.
	ConnectorName string `json:"connector_name,omitempty"`
	// ConnectorVersion holds the value of the "connector_version" field.
	ConnectorVersion string `json:"connector_version,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Input holds the value of the "input" field.
	Input map[string]interface{} `json:"input,omitempty"`
	// Errors holds the value of the "errors" field.
	Errors []string `json:"errors,omitempty"`
	// Output holds the value of the "output" field.
	Output any `json:"output,omitempty"`
}

type StepState = StepRunData
