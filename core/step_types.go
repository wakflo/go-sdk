package core

import (
	"github.com/google/uuid"
)

// StepNodeNodeAuthInput represents the authentication input for a node in a step, including a connection ID as a UUID.
type StepNodeNodeAuthInput struct {
	ConnectionID *uuid.UUID `json:"connectionId"`
}

// StepNodeFormInput represents the input structure for a form node, containing form state and input data.
type StepNodeFormInput struct {
	State    FormState    `json:"state,omitempty"`
	Input    JSONObject   `json:"input,omitempty"`
	Branches []FlowBranch `json:"branches,omitempty"`
}

// StepNodeTestResult represents the result of a test at a specific step in a process, including its timestamp, status, and output.
type StepNodeTestResult struct {
	Timestamp *int        `json:"timestamp"`
	Status    *TestStatus `json:"status"`
	Output    any         `json:"output"`
}

// StepNodeTest represents the testing data for a step node. It includes results, last test time, and last test status.
type StepNodeTest struct {
	Results        []StepNodeTestResult `json:"results"`
	LastTestTime   *int                 `json:"lastTestTime"`
	LastTestStatus *TestStatus          `json:"lastTestStatus"`
}

type StepNodeMeta struct {
	ConnectorName    string       `json:"connectorName,omitempty"`
	ConnectorVersion string       `json:"connectorVersion,omitempty"`
	TriggerType      *TriggerType `json:"triggerType,omitempty"`
	TriggerStrategy  *TriggerType `json:"triggerStrategy,omitempty"`
}

type StepNodeConnector struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type StepNodeSettings struct {
	Branch    *BranchSettings   `json:"branch,omitempty"`
	Connector StepNodeConnector `json:"connector"`
	Trigger   *TriggerSettings  `json:"trigger,omitempty"`
	Error     StepErrorSettings `json:"error,omitempty"`
}
