package context

import (
	"context"

	"github.com/wakflo/go-sdk/v2/core"
)

type BaseContext interface {
	// Context returns the underlying Go context for the execution operation.
	Context() context.Context

	// Logger returns a structured logger for the execution.
	Logger() core.Logger

	// Input returns the validated input data for the trigger execution.
	Input() core.JSONObject

	// SetInput updates the input data for the trigger execution.
	SetInput(input core.JSONObject) error
}
