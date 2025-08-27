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

package flow

import (
	"github.com/wakflo/go-sdk/v2/context"
	"github.com/wakflo/go-sdk/v2/core"
)

// BranchPathResult represents the result of a branch path evaluation.
type BranchPathResult struct {
	// PathName is the identifier of the chosen path
	PathName string `json:"pathName"`

	// Description provides a human-readable description of why this path was chosen
	Description string `json:"description,omitempty"`

	// Metadata contains additional information about the branch decision
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// BranchPath represents a potential path in a branch flow.
type BranchPath struct {
	// Name is the unique identifier for this path
	Name string `json:"name"`

	// DisplayName is a human-readable name for this path
	DisplayName string `json:"displayName"`

	// Description provides details about this path
	Description string `json:"description,omitempty"`

	// Condition is a serialized condition expression that determines if this path should be taken
	Condition string `json:"condition,omitempty"`

	// IsDefault indicates if this is the default path when no conditions match
	IsDefault bool `json:"isDefault,omitempty"`

	// Order specifies the evaluation order of paths (lower numbers evaluated first)
	Order int `json:"order"`
}

// BranchEvaluationContext provides context for branch condition evaluation.
type BranchEvaluationContext interface {
	// Context provides the perform context for the branch operation
	Context() context.PerformContext

	// Paths returns the available branch paths to evaluate
	Paths() []BranchPath

	// GetPathByName returns a specific path by name
	GetPathByName(name string) (*BranchPath, error)

	// WorkflowData provides access to the workflow data for condition evaluation
	WorkflowData() map[string]interface{}

	// Logger returns a structured logger for the branch evaluation
	Logger() core.Logger
}

// BranchConditionEvaluator defines the interface for evaluating branch conditions.
type BranchConditionEvaluator interface {
	// EvaluateCondition determines if a condition is satisfied
	EvaluateCondition(condition any, data map[string]interface{}) (bool, error)

	// ValidateCondition validates that a condition expression is syntactically correct and suitable for evaluation.
	ValidateCondition(condition any) error

	// GetFunctions returns the available functions that can be used in conditions
	GetFunctions() map[string]interface{}
}

// Branch defines the interface for branch flow control.
type Branch interface {
	Flow

	// Evaluate determines which branch path should be taken
	Evaluate(ctx BranchEvaluationContext) (*BranchPathResult, error)

	// ValidateCondition validates that a condition expression is syntactically correct
	ValidateCondition(condition any) error

	// GetConditionEvaluator returns the condition evaluator used by this branch
	GetConditionEvaluator() BranchConditionEvaluator
}
