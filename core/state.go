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

type Cursor string

type OffsetPaginationMeta struct {
	Offset     int  `json:"offset"`
	Limit      int  `json:"limit"`
	TotalItems int  `json:"totalItems"`
	HasMore    bool `json:"hasMore"`
}

type DynamicOptionsResponse struct {
	Metadata OffsetPaginationMeta `json:"metadata"`
	Items    any                  `json:"items"`
}

type DynamicOptionsFilterParams struct {
	Offset     int    `json:"offset"` // The offset of the first item to return (default: 0)
	Limit      int    `json:"limit"`  // The maximum number of items to return (default: 10)
	FilterTerm string `json:"filterTerm"`
}

type GetDynamicOptionsInput struct {
	ConnectorName    string                      `json:"connectorName"`
	ConnectorVersion string                      `json:"connectorVersion"`
	FieldName        string                      `json:"fieldName"`
	OperationName    string                      `json:"operationName"`
	StepName         string                      `json:"stepName"`
	FlowID           string                      `json:"flowId"`
	FlowVersionID    string                      `json:"flowVersionId,omitempty"`
	Input            map[string]interface{}      `json:"input,omitempty"`
	Filter           *DynamicOptionsFilterParams `json:"filter,omitempty"`
}

type StepRunData struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ProjectID holds the value of the "team_id" field.
	ProjectID uuid.UUID `json:"project_id,omitempty"`
	// Status holds the value of the "status" field.
	Status StepRunStatus `json:"status,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// FlowID holds the value of the "flow_id" field.
	FlowID uuid.UUID `json:"flow_id,omitempty"`
	// FlowVersionID holds the value of the "flow_version_id" field.
	FlowVersionID uuid.UUID `json:"flow_version_id,omitempty"`
	// RunID holds the value of the "run_id" field.
	RunID uuid.UUID `json:"run_id,omitempty"`
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
	// StartTime represents the start time of a step run.
	StartTime *time.Time `json:"start_time,omitempty"`
	// EndTime represents the end time of a step run.
	// It is a pointer to a time.Time value and can be nil.
	Duration *int64 `json:"duration,omitempty"`
	// EndTime represents the end time of a step run.
	// It is a pointer to a time.Time value and can be nil.
	EndTime *time.Time `json:"end_time,omitempty"`
}

type StepExecutionParams struct {
	// Status holds the value of the "status" field.
	Status StepRunStatus `json:"status,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// Input holds the value of the "input" field.
	Input map[string]interface{} `json:"input,omitempty"`
	// Errors holds the value of the "errors" field.
	Errors []string `json:"errors,omitempty"`
	// Output holds the value of the "output" field.
	Output any `json:"output,omitempty"`
	// StartTime represents the start time of a step run.
	StartTime *time.Time `json:"start_time,omitempty"`
	// EndTime represents the end time of a step run.
	// It is a pointer to a time.Time value and can be nil.
	EndTime *time.Time `json:"end_time,omitempty"`
}

type StepExecutionUpdateParams struct {
	// Status holds the value of the "status" field.
	Status *StepRunStatus `json:"status,omitempty"`
	// Errors holds the value of the "errors" field.
	Errors []string `json:"errors,omitempty"`
	// Output holds the value of the "output" field.
	Output any `json:"output,omitempty"`
}

type StepState = StepRunData
