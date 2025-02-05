package integration

import (
	"github.com/google/uuid"
	"github.com/wakflo/go-sdk/core"
	"time"
)

type StepRunState struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ProjectID holds the value of the "team_id" field.
	ProjectID uuid.UUID `json:"project_id,omitempty"`
	// Status holds the value of the "status" field.
	Status core.StepRunStatus `json:"status,omitempty"`
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
	Errors []core.StateError `json:"errors,omitempty"`
	// Output holds the value of the "output" field.
	Output core.JSON `json:"output,omitempty"`
	// StartTime represents the start time of a step run.
	StartTime *time.Time `json:"start_time,omitempty"`
	// EndTime represents the end time of a step run.
	// It is a pointer to a time.Time value and can be nil.
	Duration *int64 `json:"duration,omitempty"`
	// EndTime represents the end time of a step run.
	// It is a pointer to a time.Time value and can be nil.
	EndTime *time.Time `json:"end_time,omitempty"`
	// Type represents the type of the flow step, defined by the FlowStepType enumeration, such as STEP, LOOP, or ROUTER.
	Type          core.FlowStepType `json:"type,omitempty"`
	ExecutionMode ExecutionMode     `json:"mode,omitempty"`
}
