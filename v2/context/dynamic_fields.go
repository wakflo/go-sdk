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

package context

import (
	"github.com/rs/xid"
	"github.com/wakflo/go-sdk/v2/core"
)

// DynamicFieldContext defines the interface for performing an action in a workflow.
// It provides methods for handling action execution, including input validation,
// authentication, output processing, and error handling.
type DynamicFieldContext interface {
	BaseContext

	// Respond sends a response containing the provided data and total count of items, adhering to dynamic options structure.
	// Returns a `DynamicOptionsResponse` object with items and metadata or an error if processing the response fails.
	Respond(data any, totalItems int) (*core.DynamicOptionsResponse, error)

	// RespondJSON creates a JSON response containing the provided data and total item count. Returns the JSON or an error.
	RespondJSON(data any, totalItems int) (core.JSON, error)

	// WorkflowID returns the unique identifier of the workflow.
	WorkflowID() xid.ID

	// WorkflowVersionID returns the unique identifier of the workflow version.
	WorkflowVersionID() xid.ID

	// FieldName returns the name of the current field within the workflow context.
	FieldName() string

	// OperationID returns the unique identifier of the current operation within the workflow context.
	OperationID() string

	// StepID returns the unique identifier of the step in the workflow.
	StepID() string

	// Filter returns filtering parameters for dynamic options, including offset, limit, and filter term.
	Filter() *core.DynamicOptionsFilterParams
}
