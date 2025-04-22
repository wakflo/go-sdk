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

// BooleanOperator represents the logical operator to apply to boolean values.
type BooleanOperator string

const (
	// OperatorAnd performs logical AND on the input values.
	OperatorAnd BooleanOperator = "and"

	// OperatorOr performs logical OR on the input values.
	OperatorOr BooleanOperator = "or"

	// OperatorNot performs logical NOT on a single input value.
	OperatorNot BooleanOperator = "not"

	// OperatorXor performs logical XOR (exclusive or) on the input values.
	OperatorXor BooleanOperator = "xor"

	// OperatorNand performs logical NAND (not and) on the input values.
	OperatorNand BooleanOperator = "nand"

	// OperatorNor performs logical NOR (not or) on the input values.
	OperatorNor BooleanOperator = "nor"
)

// BooleanCondition represents a condition to be evaluated.
type BooleanCondition struct {
	// Expression is the condition expression to evaluate
	Expression string `json:"expression"`

	// Description provides a human-readable description of this condition
	Description string `json:"description,omitempty"`

	// Weight allows for weighted evaluation when using certain operators
	Weight float64 `json:"weight,omitempty"`
}

// BooleanResult represents the result of a boolean evaluation.
type BooleanResult struct {
	// Result is the final boolean result of the evaluation
	Result bool `json:"result"`

	// Description explains how the result was determined
	Description string `json:"description,omitempty"`

	// ConditionResults contains the individual results of each condition
	ConditionResults map[string]bool `json:"conditionResults,omitempty"`

	// Metadata contains additional information about the evaluation
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// BooleanEvaluationContext provides context for boolean condition evaluation.
type BooleanEvaluationContext interface {
	// Context provides the perform context for the boolean operation
	Context() context.PerformContext

	// Operator returns the logical operator to apply
	Operator() BooleanOperator

	// Conditions returns the list of conditions to evaluate
	Conditions() []BooleanCondition

	// WorkflowData provides access to the workflow data for condition evaluation
	WorkflowData() map[string]interface{}

	// Logger returns a structured logger for the boolean evaluation
	Logger() core.Logger
}

// Boolean defines the interface for boolean logical operations.
type Boolean interface {
	// Evaluate determines the result of applying the boolean operator to the conditions
	Evaluate(ctx BooleanEvaluationContext) (*BooleanResult, error)

	// ValidateExpression validates that a condition expression is syntactically correct
	ValidateExpression(expression string) error

	// GetSupportedOperators returns all supported boolean operators
	GetSupportedOperators() []BooleanOperator
}
