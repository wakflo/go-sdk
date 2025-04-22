package sdk

import (
	"time"

	"github.com/rs/xid"
	"github.com/wakflo/go-sdk/oldcore"
)

type ExecuteMetadata struct {
	FlowVersionID xid.ID `json:"flowVersionId,omitempty"`
	// FlowID holds the value of the "flow_id" field.
	FlowID xid.ID  `json:"flowId"`
	RunID  *xid.ID `json:"runId,omitempty"`
	// Name holds the value of the "name" field.
	StepName string `json:"stepName,omitempty"`
	// LastRun field stores the timestamp of the last run of a flow version.
	LastRun *time.Time `json:"last_run,omitempty"`
	// Status indicates the current state of the flow version, represented by the FlowVersionState enumeration.
	FlowStatus oldcore.FlowVersionState `json:"flowStatus"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID xid.ID `json:"projectId"`
	// Mode specifies the mode of execution, represented by the ExecutionMode type, and serialized as "mode" in JSON.
	Mode ExecutionMode `json:"mode"`
}
