package integration

import (
	"github.com/google/uuid"
	"github.com/wakflo/go-sdk/core"
	"time"
)

type ExecuteMetadata struct {
	FlowVersionID uuid.UUID `json:"flowVersionId,omitempty"`
	// FlowID holds the value of the "flow_id" field.
	FlowID uuid.UUID  `json:"flowId"`
	RunID  *uuid.UUID `json:"runId,omitempty"`
	// Name holds the value of the "name" field.
	StepName string `json:"stepName,omitempty"`
	// LastRun field stores the timestamp of the last run of a flow version.
	LastRun *time.Time `json:"last_run,omitempty"`
	// Status indicates the current state of the flow version, represented by the FlowVersionState enumeration.
	FlowStatus core.FlowVersionState `json:"flowStatus"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID uuid.UUID `json:"projectId"`
	// Mode specifies the mode of execution, represented by the ExecutionMode type, and serialized as "mode" in JSON.
	Mode ExecutionMode `json:"mode"`
}
