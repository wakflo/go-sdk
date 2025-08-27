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
	"github.com/wakflo/go-sdk/v2"
	"github.com/wakflo/go-sdk/v2/core"
)

// LoopType represents the type of loop to execute.
type LoopType string

const (
	// LoopTypeForEach iterates over each item in a collection.
	LoopTypeForEach LoopType = "forEach"

	// LoopTypeWhile continues execution while a condition is true.
	LoopTypeWhile LoopType = "while"

	// LoopTypeDoWhile executes at least once, then continues while a condition is true.
	LoopTypeDoWhile LoopType = "doWhile"

	// LoopTypeCount executes a fixed number of iterations.
	LoopTypeCount LoopType = "count"
)

// LoopIteration represents a single iteration of a loop.
type LoopIteration struct {
	// Index is the zero-based index of the current iteration
	Index int `json:"index"`

	// Item is the current item being processed (for forEach loops)
	Item interface{} `json:"item,omitempty"`

	// Key is the key of the current item (for map-like collections)
	Key interface{} `json:"key,omitempty"`

	// IsFirst indicates if this is the first iteration
	IsFirst bool `json:"isFirst"`

	// IsLast indicates if this is the last iteration (when known)
	IsLast bool `json:"isLast"`

	// Data contains any additional context for this iteration
	Data map[string]interface{} `json:"data,omitempty"`
}

// LoopExecutionContext provides context for loop execution.
type LoopExecutionContext interface {
	// Context provides the perform context for the loop operation
	Context() sdk.PerformContext

	// LoopType returns the type of loop being executed
	LoopType() LoopType

	// Collection returns the collection being iterated (for forEach loops)
	Collection() (interface{}, error)

	// Condition returns the condition expression (for while/doWhile loops)
	Condition() string

	// Count returns the total count for count-based loops
	Count() (int, error)

	// CurrentIteration returns information about the current iteration
	CurrentIteration() *LoopIteration

	// SetCurrentIteration updates the current iteration state
	SetCurrentIteration(*LoopIteration) error

	// MaxIterations returns the maximum allowed iterations to prevent infinite loops
	MaxIterations() int

	// Logger returns a structured logger for the loop execution
	Logger() core.Logger

	// WorkflowData provides access to the workflow data
	WorkflowData() map[string]interface{}

	// UpdateWorkflowData updates the workflow data
	UpdateWorkflowData(map[string]interface{}) error
}

// LoopController defines the interface for controlling loop execution.
type LoopController interface {
	// Initialize prepares the loop for execution
	Initialize(ctx LoopExecutionContext) error

	// HasNext determines if there are more iterations to execute
	HasNext(ctx LoopExecutionContext) (bool, error)

	// Next advances to the next iteration
	Next(ctx LoopExecutionContext) (*LoopIteration, error)

	// Reset resets the loop to its initial state
	Reset(ctx LoopExecutionContext) error

	// Break immediately exits the loop
	Break(ctx LoopExecutionContext) error

	// Continue skips to the next iteration
	Continue(ctx LoopExecutionContext) error

	// GetState returns the current state of the loop for persistence
	GetState(ctx LoopExecutionContext) (map[string]interface{}, error)

	// SetState restores the loop state for resuming execution
	SetState(ctx LoopExecutionContext, state map[string]interface{}) error
}

// Loop defines the interface for loop flow control.
type Loop interface {
	Flow

	// Execute runs the loop until completion, using the provided controller
	Execute(ctx LoopExecutionContext, controller LoopController) error

	// ValidateConfiguration validates the loop configuration
	ValidateConfiguration(ctx LoopExecutionContext) error

	// CreateController creates an appropriate controller for this loop type
	CreateController(loopType LoopType) (LoopController, error)
}
