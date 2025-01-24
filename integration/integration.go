package integration

import (
	"context"
	"github.com/wakflo/go-sdk/core"
)

type Auth struct {
	Schema   *core.AutoFormSchema `json:"schema"`
	Required bool                 `json:"required"`
}

// Integration defines an interface for managing a set of triggers and actions for a specific system or service.
type Integration interface {
	// Name returns the human-readable name for the integration (e.g., "Slack", "Stripe").
	Name() string

	// Icon returns the icon URL or path for the integration.
	Icon() string

	// Version returns the version string of the integration, indicating its current release or iteration.
	Version() string

	// Triggers returns all triggers supported by this integration.
	Triggers() []Trigger

	// Actions returns all actions supported by this integration.
	Actions() []Action

	// The Description returns a detailed explanation or summary of the trigger or action being executed.
	Description() string

	// Documentation returns an OperationDocumentation instance, providing optional detailed documentation for the operation.
	Documentation() *OperationDocumentation

	// Auth returns the authentication schema configuration required by the integration.
	Auth() *Auth
}

type OperationMetadata struct {
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty" validate:"required"`
	// HelpText holds the value of the "helpText" field.
	HelpText *string `json:"helpText,omitempty"`
	// SampleOutput holds the value of the "sampleOutput" field.
	SampleOutput map[string]any `json:"sampleOutput,omitempty"`
	// Auth holds the value of the "auth" field.
	Auth *core.AutoFormSchema `json:"auth,omitempty"`
}

type OperationDocumentation struct {
	Documentation *string `json:"documentation,omitempty"`
}

type JSONOutput any

type OperationContextFn = func() (*JSONOutput, error)

type BaseOperation interface {
	// Name returns the human-readable name for the integration (e.g., "Slack", "Stripe").
	Name() string

	// The Description returns a detailed explanation or summary of the trigger or action being executed.
	Description() string

	// Documentation returns an OperationDocumentation instance, providing optional detailed documentation for the operation.
	Documentation() *OperationDocumentation

	// SampleData retrieves example or mock data related to the operation, often used for testing or integration scenarios.
	// Returns a pointer to the JSON data and an error if the retrieval fails.
	SampleData() (*core.JSON, error)

	// Test executes a test operation to validate connectivity or configuration, returning sample data or an error.
	Test() (*core.JSON, error)

	// Properties returns a map of property names to their corresponding AutoFormSchema definitions.
	Properties() map[string]core.AutoFormSchema

	// Auth returns the authentication schema required for the operation, defined as an AutoFormSchema.
	Auth() *Auth
}

// Trigger defines a generic interface for workflow triggers.
type Trigger interface {
	BaseOperation

	// Start prepares and activates the trigger (e.g., start polling, event listening, cron schedules, etc.).
	Start(ctx context.Context) error

	// Stop gracefully stops or disables the trigger (e.g., unsubscribe or clean up resources).
	Stop(ctx context.Context) error

	// Execute handles the trigger's action when manually invoked with an input schema.
	Execute(ctx context.Context) error

	// GetType returns the trigger type (e.g., POLLING, EVENT, WEBHOOK, MANUAL, SCHEDULED).
	GetType() core.TriggerType

	// Criteria Additional criteria or filter rules required to activate the trigger.
	Criteria(ctx context.Context) core.TriggerCriteria
}

// Action defines a generic interface for performing an operation within the integration.
type Action interface {
	BaseOperation

	// Perform executes the action with the given input schema and returns a result or an error.
	Perform(ctx context.Context) (map[string]any, error)
}
