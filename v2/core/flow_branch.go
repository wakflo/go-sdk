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
	"github.com/rs/xid"
)

// NewFlowRouter creates a new branch with defaults
func NewFlowRouter(name string, nextStep string) *FlowRouter {
	return &FlowRouter{
		ID:       xid.New().String(),
		Name:     name,
		NextStep: nextStep,
		Order:    0,
	}
}

// WithCondition adds a condition to the branch
func (b *FlowRouter) WithCondition(condition string) *FlowRouter {
	b.Condition = condition
	return b
}

// WithValue adds a value to the branch
func (b *FlowRouter) WithValue(value interface{}) *FlowRouter {
	b.Value = value
	return b
}

// WithOrder sets the order of the branch
func (b *FlowRouter) WithOrder(order int) *FlowRouter {
	b.Order = order
	return b
}

// AsDefault sets this branch as the default
func (b *FlowRouter) AsDefault() *FlowRouter {
	b.IsDefault = true
	return b
}

// WithFieldCondition sets up a field-based condition
func (b *FlowRouter) WithFieldCondition(field string, operator string, value interface{}) *FlowRouter {
	b.ConditionField = field
	b.ConditionOperator = operator
	b.ConditionValue = value
	return b
}

// BranchSettings defines the configuration for a branching flow step.
type BranchSettings struct {
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
	BranchMode RouteMode `json:"branchMode,omitempty"`

	// Branches defines the specific branches available
	Branches []FlowRouter `json:"branches,omitempty"`
}

// NewBranchSettings creates a new BranchSettings with default values
func NewBranchSettings() *BranchSettings {
	return &BranchSettings{
		EvaluateAll:      false,
		AllowMultiPath:   false,
		BranchMode:       RouteModeExpression,
		ExpressionEngine: "javascript",
		Branches:         make([]FlowRouter, 0),
	}
}

// AddBranch adds a new branch to the BranchSettings
func (bs *BranchSettings) AddBranch(branch FlowRouter) {
	bs.Branches = append(bs.Branches, branch)
}

// GetBranchByID retrieves a branch by its ID
func (bs *BranchSettings) GetBranchByID(id string) *FlowRouter {
	for i, branch := range bs.Branches {
		if branch.ID == id {
			return &bs.Branches[i]
		}
	}
	return nil
}

// GetDefaultBranch gets the default branch
func (bs *BranchSettings) GetDefaultBranch() *FlowRouter {
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
func (bs *BranchSettings) SetDefaultBranch(id string) {
	bs.DefaultBranch = id

	// Also update IsDefault flags for clarity
	for i := range bs.Branches {
		bs.Branches[i].IsDefault = (bs.Branches[i].ID == id)
	}
}
