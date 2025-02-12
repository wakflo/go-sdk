package core

import (
	"time"

	"github.com/robfig/cron/v3"
)

// PollingTriggerCriteria defines criteria for a polling-based trigger.
type PollingTriggerCriteria struct {
	// Interval defines the duration between each poll (e.g., every 5 minutes).
	Interval time.Duration `json:"interval" validate:"required,gt=0"`

	// MaxRetries specifies the maximum number of retries in case polling fails.
	MaxRetries int `json:"maxRetries,omitempty" validate:"min=0"`

	// FetchLimit defines the maximum number of items to fetch per poll.
	FetchLimit int `json:"fetchLimit,omitempty" validate:"min=1"`

	// DataFilters allows specifying conditions for the data being fetched.
	DataFilters map[string]any `json:"dataFilters,omitempty"`

	// AllowEmptyData determines if polling should proceed without error when no data is returned in a poll.
	AllowEmptyData bool `json:"allowEmptyData"`

	// LastExecutionTime holds the timestamp of the last successful poll, useful for incremental polling.
	LastExecutionTime *time.Time `json:"lastExecutionTime,omitempty"`

	// Enabled determines whether the polling trigger is active.
	Enabled bool `json:"enabled" validate:"required"`

	// ExcludedDays allows specifying days when polling should be skipped (e.g., holidays).
	ExcludedDays []time.Weekday `json:"excludedDays,omitempty"`
}

// NewPollingTriggerCriteria initializes and returns a new PollingTriggerCriteria instance with default values applied.
func NewPollingTriggerCriteria() *PollingTriggerCriteria {
	c := &PollingTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values for PollingTriggerCriteria.
func (c *PollingTriggerCriteria) SetDefaults() {
	if c.Interval == 0 {
		c.Interval = 5 * time.Minute // Default to 5 minutes.
	}
	if c.MaxRetries == 0 {
		c.MaxRetries = 3 // Default to 3 retries.
	}
	if c.FetchLimit == 0 {
		c.FetchLimit = 10 // Default to fetching 10 items per poll.
	}
	if c.DataFilters == nil {
		c.DataFilters = make(map[string]any) // Default to an empty filter map.
	}
	if c.ExcludedDays == nil {
		c.ExcludedDays = []time.Weekday{} // Default to no excluded days.
	}
}

// WebhookTriggerCriteria defines the criteria for triggering via webhook.
type WebhookTriggerCriteria struct {
	// Endpoint is the URL where the webhook should listen for incoming requests.
	Endpoint string `json:"endpoint" validate:"required,url"`

	// HttpMethod specifies the HTTP method for the webhook (e.g., POST, GET).
	HttpMethod string `json:"httpMethod" validate:"required,oneof=GET POST PUT DELETE"`

	// AuthEnabled determines if the webhook requires authentication.
	AuthEnabled bool `json:"authEnabled"`

	// Headers specify any required headers that must be included in the webhook requests.
	Headers map[string]string `json:"headers,omitempty"`

	// QueryParams define any required query parameters for validation.
	QueryParams map[string]string `json:"queryParams,omitempty"`

	// ValidationSecret is an optional field for validating webhook requests (e.g., HMAC secret).
	ValidationSecret *string `json:"validationSecret,omitempty"`

	// Enabled determines if the webhook trigger is active.
	Enabled bool `json:"enabled" validate:"required"`
}

// NewWebhookTriggerCriteria creates a new instance of WebhookTriggerCriteria and initializes default values.
func NewWebhookTriggerCriteria() *WebhookTriggerCriteria {
	c := &WebhookTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values for WebhookTriggerCriteria.
func (c *WebhookTriggerCriteria) SetDefaults() {
	if c.HttpMethod == "" {
		c.HttpMethod = "POST" // Default to POST method.
	}
	if c.Headers == nil {
		c.Headers = make(map[string]string) // Default to an empty header map.
	}
	if c.QueryParams == nil {
		c.QueryParams = make(map[string]string) // Default to an empty query parameter map.
	}
}

// ManualTriggerCriteria defines criteria for manually triggered workflows.
type ManualTriggerCriteria struct {
	// Description is an optional field to describe the purpose of the manual trigger.
	Description string `json:"description,omitempty"`

	// UserInputSchema defines a schema for the input required when triggering manually.
	UserInputSchema *AutoFormSchema `json:"userInputSchema,omitempty"`

	// RequireApproval specifies if the manual trigger requires prior approval to execute.
	RequireApproval bool `json:"requireApproval"`

	// ExecutionTimeout defines the maximum allowed time for the manual step to be completed.
	ExecutionTimeout *time.Duration `json:"executionTimeout,omitempty"`
}

// NewManualTriggerCriteria creates a new instance of ManualTriggerCriteria and sets its default values.
func NewManualTriggerCriteria() *ManualTriggerCriteria {
	c := &ManualTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values for ManualTriggerCriteria.
func (c *ManualTriggerCriteria) SetDefaults() {
	if c.Description == "" {
		c.Description = "Manual Trigger" // Default description.
	}

	if c.ExecutionTimeout == nil {
		defaultTimeout := 1 * time.Hour
		c.ExecutionTimeout = &defaultTimeout // Default to 1 hour execution timeout.
	}
}

// EventTriggerCriteria defines trigger criteria based on system or external events.
type EventTriggerCriteria struct {
	// EventName specifies the name of the event to subscribe to (e.g., "file_uploaded").
	EventName string `json:"eventName" validate:"required"`

	// Source specifies the source of the event (e.g., "system", "integration_name").
	Source string `json:"source" validate:"required"`

	// Filters allow refining event criteria (e.g., event-specific parameters or payload conditions).
	Filters map[string]any `json:"filters,omitempty"`
}

// NewEventTriggerCriteria initializes a new EventTriggerCriteria instance with default values applied.
func NewEventTriggerCriteria() *EventTriggerCriteria {
	c := &EventTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values for EventTriggerCriteria.
func (c *EventTriggerCriteria) SetDefaults() {
	if c.Filters == nil {
		c.Filters = make(map[string]any) // Default to an empty filter map.
	}
}

// ScheduleTriggerCriteria defines trigger criteria for scheduled workflows.
type ScheduleTriggerCriteria struct {
	// CronExpression specifies the schedule using a cron expression.
	CronExpression string `json:"cronExpression" validate:"required,cron"`

	// StartTime specifies the time to begin the schedule execution.
	StartTime *time.Time `json:"startTime,omitempty"`

	// EndTime specifies the time to stop the schedule execution.
	EndTime *time.Time `json:"endTime,omitempty"`

	// TimeZone allows defining the time zone for the cron schedule.
	TimeZone string `json:"timeZone" validate:"required"`

	// Enabled determines if the schedule trigger is active.
	Enabled bool `json:"enabled" validate:"required"`
}

// ScheduledInterval derives the scheduled interval as a time.Duration based on the cron expression.
func (c *ScheduleTriggerCriteria) ScheduledInterval() time.Duration {
	// This is a simplified derivation. Use a cron library for precise parsing.
	schedule, err := cron.ParseStandard(c.CronExpression)
	if err != nil {
		return -1 // Interval unknown or invalid cron expression
	}

	now := time.Now()
	next := schedule.Next(now)
	interval := next.Sub(now)

	return interval
}

// NewScheduleTriggerCriteria creates and initializes a new ScheduleTriggerCriteria instance with default values applied.
func NewScheduleTriggerCriteria() *ScheduleTriggerCriteria {
	c := &ScheduleTriggerCriteria{}
	c.ApplyDefaults()
	return c
}

// ApplyDefaults applies default values to a ScheduleTriggerCriteria struct.
func (c *ScheduleTriggerCriteria) ApplyDefaults() {
	if c.CronExpression == "" {
		c.CronExpression = "0 0 * * *" // Default to daily execution.
	}
	if c.TimeZone == "" {
		c.TimeZone = "UTC" // Default to UTC timezone.
	}
	if !c.Enabled {
		c.Enabled = true // Enable the schedule trigger by default.
	}
}

// RetryPolicy defines the retry policy for failed event triggers.
type RetryPolicy struct {
	Enabled bool `json:"enabled" validate:"required"`

	// MaxRetries specifies the maximum number of retry attempts.
	MaxRetries int `json:"maxRetries" validate:"min=0"`

	// RetryInterval defines the time to wait between retry attempts.
	RetryInterval time.Duration `json:"retryInterval" validate:"gt=0"`

	// ExponentialBackoff determines if exponential backoff is used for retries.
	ExponentialBackoff bool `json:"exponentialBackoff"`
}

// NewRetryPolicy creates and returns a new RetryPolicy instance with default values applied.
func NewRetryPolicy() *RetryPolicy {
	rp := &RetryPolicy{}
	rp.SetDefaults()
	return rp
}

// SetDefaults applies default values for RetryPolicy.
func (rp *RetryPolicy) SetDefaults() {
	rp.Enabled = false

	if rp.MaxRetries == 0 {
		rp.MaxRetries = 3 // Default to 3 retries.
	}
	if rp.RetryInterval == 0 {
		rp.RetryInterval = 1 * time.Minute // Default to 1 minute retry interval.
	}
}

// TriggerCriteria defines the settings and criteria for configuring step-based triggers in the system.
type TriggerCriteria struct {
	// Polling specifies the configuration used for defining triggers based on polling mechanisms.
	Polling *PollingTriggerCriteria `json:"pollingCriteria,omitempty"`

	// Event specifies the conditions that define event-based triggers in the system.
	Event *EventTriggerCriteria `json:"eventCriteria,omitempty"`

	// Manual defines the criteria needed for a manual trigger configuration in the system.
	Manual *ManualTriggerCriteria `json:"manualCriteria,omitempty"`

	// Webhook specifies the conditions required for triggering actions via a webhook.
	Webhook *WebhookTriggerCriteria `json:"webhookCriteria,omitempty"`

	// Schedule specifies the criteria for configuring scheduled triggers in the system.
	Schedule *ScheduleTriggerCriteria `json:"scheduleCriteria,omitempty"`
}

// NewTriggerCriteria creates a new instance of TriggerSettings and applies default values using SetDefaults().
func NewTriggerCriteria() *TriggerCriteria {
	ts := &TriggerCriteria{}
	ts.SetDefaults()
	return ts
}

// SetDefaults applies default values to TriggerSettings.
func (ts *TriggerCriteria) SetDefaults() {
	if ts.Polling != nil {
		ts.Polling.SetDefaults()
	}
	if ts.Event != nil {
		ts.Event.SetDefaults()
	}
	if ts.Manual != nil {
		ts.Manual.SetDefaults()
	}
	if ts.Webhook != nil {
		ts.Webhook.SetDefaults()
	}
	if ts.Schedule != nil {
		ts.Schedule.ApplyDefaults()
	}
}

// SetDefaultByType applies default values to TriggerSettings.
func (ts *TriggerCriteria) SetDefaultByType(strategy TriggerType) {
	switch strategy {
	case TriggerTypeManual:
		ts.Manual = NewManualTriggerCriteria()
		break
	case TriggerTypeEvent:
		ts.Event = NewEventTriggerCriteria()
		break
	case TriggerTypePolling:
		ts.Polling = NewPollingTriggerCriteria()
		break
	case TriggerTypeWebhook:
		ts.Webhook = NewWebhookTriggerCriteria()
	case TriggerTypeScheduled:
		ts.Schedule = NewScheduleTriggerCriteria()
		break
	}
}

// TriggerSettings defines the settings and criteria for configuring step-based triggers in the system.
type TriggerSettings struct {
	// Type holds the value of the "type" field.
	Type TriggerType `json:"strategy,omitempty" validate:"required,oneof=SCHEDULED EVENT POLLING WEBHOOK MANUAL"`

	// Criteria specifies the configuration settings and conditions for initializing step-based triggers.
	Criteria *TriggerCriteria `json:"criteria,omitempty"`

	// RetryPolicy specifies the retry behavior for handling failed event triggers, including limits and intervals.
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`
}

// SetDefaults applies default values to TriggerSettings.
func (ts *TriggerSettings) SetDefaults() {
	if ts.Criteria == nil {
		ts.Criteria = NewTriggerCriteria()
	}

	if ts.Type == "" {
		ts.Type = TriggerTypeManual
	}

	if ts.RetryPolicy != nil {
		ts.RetryPolicy.SetDefaults()
	}

	ts.Criteria.SetDefaultByType(ts.Type)
}

// NewTriggerSettings creates a new instance of TriggerSettings and applies default values using SetDefaults().
func NewTriggerSettings() *TriggerSettings {
	ts := &TriggerSettings{}
	ts.SetDefaults()
	return ts
}

// SetDefaultByType applies default values to TriggerSettings.
func (ts *TriggerSettings) SetDefaultByType(strategy TriggerType) {
	switch strategy {
	case TriggerTypeEvent:
	case TriggerTypePolling:
	case TriggerTypeScheduled:
	case TriggerTypeWebhook:
		ts.RetryPolicy = NewRetryPolicy()
		break
	case TriggerTypeManual:
		break
	}
}
