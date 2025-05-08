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

	// AuthContext provides authentication context for the trigger execution.
	AuthContext() (*AuthContext, error)

	// Auth provides authentication context for the trigger execution.
	Auth() *AuthContext
}
