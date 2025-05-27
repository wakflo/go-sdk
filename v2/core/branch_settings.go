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

// RouterType defines the type of router.
type RouterType string

const (
	// RouterTypeSwitch routes based on evaluating an expression and matching the result.
	RouterTypeSwitch RouterType = "switch"

	// RouterTypeCondition routes based on evaluating multiple conditions and taking the first match.
	RouterTypeCondition RouterType = "condition"

	// RouterTypeMultiPath allows taking multiple paths if their conditions are true.
	RouterTypeMultiPath RouterType = "multipath"
)

// RouterSettings defines the configuration for a branching flow step.
type RouterSettings struct {
	Type RouterType `json:"type,omitempty"`

	// DefaultBranch specifies which branch to follow when no condition is met
	DefaultBranch string `json:"defaultBranch,omitempty"`

	// EvaluateAll determines if all branch conditions should be evaluated
	// or only until the first match
	EvaluateAll bool `json:"evaluateAll,omitempty"`

	// Expression is a JavaScript expression that determines which branch to take
	Expression string `json:"expression,omitempty"`

	// ExpressionEngine defines which engine to use for expression evaluation
	// Default is "javascript"
	ExpressionEngine string `json:"expressionEngine,omitempty"`

	// ConditionField is the specific field to evaluate in the step's output
	ConditionField string `json:"conditionField,omitempty"`

	// ConditionOperator defines how to compare the field value (equals, contains, etc.)
	ConditionOperator string `json:"conditionOperator,omitempty"`

	// ConditionValue is the value to compare against
	ConditionValue interface{} `json:"conditionValue,omitempty"`

	// AllowMultiPath determines if multiple branches can be activated simultaneously
	AllowMultiPath bool `json:"allowMultiPath,omitempty"`

	// BranchMode defines how branches are evaluated (expression, condition, or value)
	BranchMode BranchMode `json:"branchMode,omitempty"`

	// Branches defines the specific branches available
	Branches []FlowRouter `json:"branches,omitempty"`
}

// BranchMode defines how branches in a workflow are selected.
type BranchMode string

const (
	// BranchModeExpression uses a JavaScript expression to determine the branch
	BranchModeExpression BranchMode = "expression"

	// BranchModeCondition uses field-based condition matching
	BranchModeCondition BranchMode = "condition"

	// BranchModeValue uses direct value matching
	BranchModeValue BranchMode = "value"
)

// FlowRouter represents a single branch option in a branching flow.
type FlowRouter struct {
	// ID is the unique identifier of the branch
	ID string `json:"id,omitempty"`

	// Name is the display name of the branch
	Name string `json:"name,omitempty"`

	// Condition is the condition expression for this branch
	Condition any `json:"condition,omitempty"`

	// Value is the expected value for this branch
	// (used with BranchModeValue)
	Value interface{} `json:"value,omitempty"`

	// Order determines the evaluation order of branches
	Order int `json:"order,omitempty"`

	// NextStep is the ID of the step to execute if this branch is selected
	NextStep string `json:"nextStep,omitempty"`

	// IsDefault indicates if this is the default branch
	IsDefault bool `json:"isDefault,omitempty"`

	// ConditionField is the specific field to evaluate for this branch
	ConditionField string `json:"conditionField,omitempty"`

	// ConditionOperator defines how to compare values for this branch
	ConditionOperator string `json:"conditionOperator,omitempty"`

	// ConditionValue is the value to compare against for this branch
	ConditionValue interface{} `json:"conditionValue,omitempty"`

	// Description is the description of the branch
	Description string `json:"description,omitempty"`
}

// NewRouterSettings creates a new BranchSettings with default values
func NewRouterSettings() *RouterSettings {
	return &RouterSettings{
		EvaluateAll:    false,
		AllowMultiPath: false,
		BranchMode:     BranchModeCondition,
		Branches:       make([]FlowRouter, 0),
	}
}

// AddBranch adds a new branch to the RouterSettings
func (bs *RouterSettings) AddBranch(branch FlowRouter) {
	bs.Branches = append(bs.Branches, branch)
}

// GetBranchByID retrieves a branch by its ID
func (bs *RouterSettings) GetBranchByID(id string) *FlowRouter {
	for i, branch := range bs.Branches {
		if branch.ID == id {
			return &bs.Branches[i]
		}
	}
	return nil
}

// GetDefaultBranch gets the default branch
func (bs *RouterSettings) GetDefaultBranch() *FlowRouter {
	// First check for explicitly marked default branch
	for i, branch := range bs.Branches {
		if branch.IsDefault {
			return &bs.Branches[i]
		}
	}

	// Then check the DefaultBranch field
	if bs.DefaultBranch != "" {
		return bs.GetBranchByID(bs.DefaultBranch)
	}

	// If no default is specified and there's only one branch, use that
	if len(bs.Branches) == 1 {
		return &bs.Branches[0]
	}

	return nil
}

// SetDefaultBranch sets a branch as the default
func (bs *RouterSettings) SetDefaultBranch(id string) {
	bs.DefaultBranch = id

	// Also update IsDefault flags for clarity
	for i := range bs.Branches {
		bs.Branches[i].IsDefault = (bs.Branches[i].ID == id)
	}
}
