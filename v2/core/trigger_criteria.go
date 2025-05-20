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
	"time"

	"github.com/juicycleff/smartform/v1"
	"github.com/robfig/cron/v3"
)

// PollingTriggerCriteria defines criteria for a polling-based trigger
type PollingTriggerCriteria struct {
	// Interval defines the duration between each poll
	Interval time.Duration `json:"interval" validate:"required,gt=0"`

	// MinInterval defines the duration between each poll
	MinInterval time.Duration `json:"minInterval" validate:"required,gt=0"`

	// MaxInterval defines the duration between each poll
	MaxInterval time.Duration `json:"maxInterval" validate:"required,gt=0"`

	// MaxRetries specifies the maximum number of retries in case polling fails
	MaxRetries int `json:"maxRetries,omitempty" validate:"min=0"`

	// FetchLimit defines the maximum number of items to fetch per poll
	FetchLimit int `json:"fetchLimit,omitempty" validate:"min=1"`

	// DataFilters allows specifying conditions for the data being fetched
	DataFilters map[string]any `json:"dataFilters,omitempty"`

	// AllowEmptyData determines if polling should proceed without error when no data is returned
	AllowEmptyData bool `json:"allowEmptyData"`

	// LastExecutionTime holds the timestamp of the last successful poll
	LastExecutionTime *time.Time `json:"lastExecutionTime,omitempty"`

	// Enabled determines whether the polling trigger is active
	Enabled bool `json:"enabled" validate:"required"`

	// ExcludedDays allows specifying days when polling should be skipped
	ExcludedDays []time.Weekday `json:"excludedDays,omitempty"`

	// ExcludedHours allows specifying hours when polling should be skipped
	ExcludedHours []int `json:"excludedHours,omitempty"`

	// TimeoutSeconds is the maximum time allowed for a polling operation
	TimeoutSeconds int `json:"timeoutSeconds,omitempty"`

	// Deduplicate determines if duplicate items should be filtered
	Deduplicate bool `json:"deduplicate"`

	// DedupKeyPath is the path to the field used for deduplication
	DedupKeyPath string `json:"dedupKeyPath,omitempty"`
}

// NewPollingTriggerCriteria initializes and returns a new PollingTriggerCriteria instance with default values
func NewPollingTriggerCriteria() *PollingTriggerCriteria {
	c := &PollingTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values for PollingTriggerCriteria
func (c *PollingTriggerCriteria) SetDefaults() {
	if c.Interval == 0 {
		c.Interval = 5 * time.Minute // Default to 5 minutes
	}
	if c.MaxRetries == 0 {
		c.MaxRetries = 3 // Default to 3 retries
	}
	if c.FetchLimit == 0 {
		c.FetchLimit = 10 // Default to fetching 10 items per poll
	}
	if c.DataFilters == nil {
		c.DataFilters = make(map[string]any) // Default to an empty filter map
	}
	if c.ExcludedDays == nil {
		c.ExcludedDays = []time.Weekday{} // Default to no excluded days
	}
	if c.ExcludedHours == nil {
		c.ExcludedHours = []int{} // Default to no excluded hours
	}
	if c.TimeoutSeconds == 0 {
		c.TimeoutSeconds = 60 // Default to 60 seconds timeout
	}
	c.Enabled = true // Default to enabled
}

// WebhookTriggerCriteria defines the criteria for triggering via webhook
type WebhookTriggerCriteria struct {
	// Endpoint is the URL where the webhook should listen for incoming requests
	Endpoint string `json:"endpoint" validate:"required,url"`

	// HttpMethod specifies the HTTP method for the webhook
	HttpMethod string `json:"httpMethod" validate:"required,oneof=GET POST PUT DELETE"`

	// AuthEnabled determines if the webhook requires authentication
	AuthEnabled bool `json:"authEnabled"`

	// Headers specify required headers that must be included in the webhook requests
	Headers map[string]string `json:"headers,omitempty"`

	// QueryParams define required query parameters for validation
	QueryParams map[string]string `json:"queryParams,omitempty"`

	// ValidationSecret is used to validate webhook requests
	ValidationSecret *string `json:"validationSecret,omitempty"`

	// Enabled determines if the webhook trigger is active
	Enabled bool `json:"enabled" validate:"required"`

	// ResponseTemplate is a template for the webhook response
	ResponseTemplate string `json:"responseTemplate,omitempty"`

	// StatusCode is the HTTP status code to return
	StatusCode int `json:"statusCode,omitempty"`

	// PayloadPath is the path to extract the payload from the request
	PayloadPath string `json:"payloadPath,omitempty"`

	// ContentType is the expected content type of the webhook request
	ContentType string `json:"contentType,omitempty"`

	BasePath     string `json:"basePath,omitempty"`
	SecretHeader string `json:"secretHeader,omitempty"`
}

// NewWebhookTriggerCriteria creates a new instance of WebhookTriggerCriteria with default values
func NewWebhookTriggerCriteria() *WebhookTriggerCriteria {
	c := &WebhookTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values for WebhookTriggerCriteria
func (c *WebhookTriggerCriteria) SetDefaults() {
	if c.HttpMethod == "" {
		c.HttpMethod = "POST" // Default to POST method
	}
	if c.Headers == nil {
		c.Headers = make(map[string]string) // Default to an empty header map
	}
	if c.QueryParams == nil {
		c.QueryParams = make(map[string]string) // Default to an empty query parameter map
	}
	if c.StatusCode == 0 {
		c.StatusCode = 200 // Default to 200 OK
	}
	if c.ContentType == "" {
		c.ContentType = "application/json" // Default to JSON content type
	}
	c.Enabled = true // Default to enabled
}

// ManualTriggerCriteria defines criteria for manually triggered workflows
type ManualTriggerCriteria struct {
	// Description is an optional field to describe the purpose of the manual trigger
	Description string `json:"description,omitempty"`

	// UserInputSchema defines a schema for the input required when triggering manually
	UserInputSchema *smartform.FormSchema `json:"userInputSchema,omitempty"`

	// RequireApproval specifies if the manual trigger requires prior approval to execute
	RequireApproval bool `json:"requireApproval"`

	// ExecutionTimeout defines the maximum allowed time for the manual step to be completed
	ExecutionTimeout *time.Duration `json:"executionTimeout,omitempty"`

	// Approvers is a list of users who can approve the trigger
	Approvers []string `json:"approvers,omitempty"`

	// ApproverGroups is a list of groups who can approve the trigger
	ApproverGroups []string `json:"approverGroups,omitempty"`

	// Tags provide searchable metadata for this trigger
	Tags []string `json:"tags,omitempty"`

	// ShowInUI determines if this trigger should be visible in the UI
	ShowInUI bool `json:"showInUI"`

	// ButtonLabel is the text to display on the trigger button
	ButtonLabel string `json:"buttonLabel,omitempty"`

	// ButtonIcon is the icon to display on the trigger button
	ButtonIcon string `json:"buttonIcon,omitempty"`
}

// NewManualTriggerCriteria creates a new instance of ManualTriggerCriteria with default values
func NewManualTriggerCriteria() *ManualTriggerCriteria {
	c := &ManualTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values for ManualTriggerCriteria
func (c *ManualTriggerCriteria) SetDefaults() {
	if c.Description == "" {
		c.Description = "Manual Trigger" // Default description
	}

	if c.ExecutionTimeout == nil {
		defaultTimeout := 1 * time.Hour
		c.ExecutionTimeout = &defaultTimeout // Default to 1 hour execution timeout
	}

	if c.ButtonLabel == "" {
		c.ButtonLabel = "Run Workflow" // Default button label
	}

	if c.ButtonIcon == "" {
		c.ButtonIcon = "play-circle" // Default button icon
	}

	c.ShowInUI = true // Default to showing in UI
}

// EventTriggerCriteria defines trigger criteria based on system or external events
type EventTriggerCriteria struct {
	// EventName specifies the name of the event to subscribe to
	EventName string `json:"eventName" validate:"required"`

	// Source specifies the source of the event
	Source string `json:"source" validate:"required"`

	// Filters allow refining event criteria
	Filters map[string]any `json:"filters,omitempty"`

	// EventPattern is a pattern to match against the event
	EventPattern string `json:"eventPattern,omitempty"`

	// MaxEvents is the maximum number of events to process at once
	MaxEvents int `json:"maxEvents,omitempty"`

	// Deduplicate determines if duplicate events should be filtered
	Deduplicate bool `json:"deduplicate"`

	// Batching determines if events should be batched
	Batching bool `json:"batching"`

	// BatchSize is the number of events to include in a batch
	BatchSize int `json:"batchSize,omitempty"`

	// BatchWindow is the time window for event batching
	BatchWindow string `json:"batchWindow,omitempty"`
}

// NewEventTriggerCriteria initializes a new EventTriggerCriteria instance with default values
func NewEventTriggerCriteria() *EventTriggerCriteria {
	c := &EventTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values for EventTriggerCriteria
func (c *EventTriggerCriteria) SetDefaults() {
	if c.Filters == nil {
		c.Filters = make(map[string]any) // Default to an empty filter map
	}
	if c.MaxEvents == 0 {
		c.MaxEvents = 10 // Default to 10 events
	}
	if c.BatchSize == 0 {
		c.BatchSize = 10 // Default to 10 events per batch
	}
	if c.BatchWindow == "" {
		c.BatchWindow = "1m" // Default to 1 minute batch window
	}
}

// ScheduleTriggerCriteria defines trigger criteria for scheduled workflows
type ScheduleTriggerCriteria struct {
	// CronExpression specifies the schedule using a cron expression
	CronExpression string `json:"cronExpression" validate:"required,cron"`

	// StartTime specifies the time to begin the schedule execution
	StartTime *time.Time `json:"startTime,omitempty"`

	// EndTime specifies the time to stop the schedule execution
	EndTime *time.Time `json:"endTime,omitempty"`

	// TimeZone allows defining the time zone for the cron schedule
	TimeZone string `json:"timeZone" validate:"required"`

	// Enabled determines if the schedule trigger is active
	Enabled bool `json:"enabled" validate:"required"`

	// MaxConcurrentRuns is the maximum number of concurrent runs
	MaxConcurrentRuns int `json:"maxConcurrentRuns,omitempty"`

	// SkipIfPreviousRunning determines if a run should be skipped if the previous one is still running
	SkipIfPreviousRunning bool `json:"skipIfPreviousRunning"`

	// MissedRunPolicy defines what to do with missed runs (skip, run-once, run-all)
	MissedRunPolicy string `json:"missedRunPolicy,omitempty"`

	// Description provides additional information about the schedule
	Description string `json:"description,omitempty"`
}

// ScheduledInterval derives the scheduled interval as a time.Duration based on the cron expression
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

// NewScheduleTriggerCriteria creates a new ScheduleTriggerCriteria with default values
func NewScheduleTriggerCriteria() *ScheduleTriggerCriteria {
	c := &ScheduleTriggerCriteria{}
	c.ApplyDefaults()
	return c
}

// ApplyDefaults applies default values to a ScheduleTriggerCriteria
func (c *ScheduleTriggerCriteria) ApplyDefaults() {
	if c.CronExpression == "" {
		c.CronExpression = "0 0 * * *" // Default to daily execution
	}
	if c.TimeZone == "" {
		c.TimeZone = "UTC" // Default to UTC timezone
	}
	if c.MaxConcurrentRuns == 0 {
		c.MaxConcurrentRuns = 1 // Default to 1 concurrent run
	}
	if c.MissedRunPolicy == "" {
		c.MissedRunPolicy = "run-once" // Default to run-once policy
	}
	c.Enabled = true // Default to enabled
}

// APITriggerCriteria defines criteria for API-triggered workflows
type APITriggerCriteria struct {
	// Enabled determines if the API trigger is active
	Enabled bool `json:"enabled"`

	// Endpoint is the API endpoint path
	Endpoint string `json:"endpoint,omitempty"`

	// Method is the HTTP method for the API endpoint
	Method string `json:"method,omitempty"`

	// Authentication determines if authentication is required
	Authentication bool `json:"authentication"`

	// RequiredScope is the required OAuth scope
	RequiredScope string `json:"requiredScope,omitempty"`

	// InputSchema defines the expected input schema
	InputSchema *smartform.FormSchema `json:"inputSchema,omitempty"`

	// RateLimit is the maximum number of requests per minute
	RateLimit int `json:"rateLimit,omitempty"`
}

// NewAPITriggerCriteria creates a new APITriggerCriteria with default values
func NewAPITriggerCriteria() *APITriggerCriteria {
	c := &APITriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values to APITriggerCriteria
func (c *APITriggerCriteria) SetDefaults() {
	if c.Method == "" {
		c.Method = "POST" // Default to POST method
	}
	if c.RateLimit == 0 {
		c.RateLimit = 60 // Default to 60 requests per minute
	}
	c.Enabled = true // Default to enabled
}

// WorkflowTriggerCriteria defines criteria for workflow-triggered workflows
type WorkflowTriggerCriteria struct {
	// SourceWorkflowID is the ID of the source workflow
	SourceWorkflowID string `json:"sourceWorkflowId,omitempty"`

	// TriggerOnEvents determines which events trigger the workflow
	TriggerOnEvents []string `json:"triggerOnEvents,omitempty"`

	// InputMapping defines how to map source workflow outputs to this workflow inputs
	InputMapping map[string]string `json:"inputMapping,omitempty"`

	// Condition is a condition that must be met for the trigger to fire
	Condition string `json:"condition,omitempty"`

	// Enabled determines if the workflow trigger is active
	Enabled bool `json:"enabled"`
}

// NewWorkflowTriggerCriteria creates a new WorkflowTriggerCriteria with default values
func NewWorkflowTriggerCriteria() *WorkflowTriggerCriteria {
	c := &WorkflowTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values to WorkflowTriggerCriteria
func (c *WorkflowTriggerCriteria) SetDefaults() {
	if c.TriggerOnEvents == nil {
		c.TriggerOnEvents = []string{"success", "failure"} // Default to success and failure events
	}
	if c.InputMapping == nil {
		c.InputMapping = make(map[string]string) // Default to empty input mapping
	}
	c.Enabled = true // Default to enabled
}

// MessageTriggerCriteria defines criteria for message-triggered workflows
type MessageTriggerCriteria struct {
	// MessageType is the type of message that triggers the workflow
	MessageType string `json:"messageType,omitempty"`

	// Source is the source of the message
	Source string `json:"source,omitempty"`

	// Destination is the destination of the message
	Destination string `json:"destination,omitempty"`

	// Filter is a filter expression applied to the message
	Filter string `json:"filter,omitempty"`

	// Enabled determines if the message trigger is active
	Enabled bool `json:"enabled"`
}

// NewMessageTriggerCriteria creates a new MessageTriggerCriteria with default values
func NewMessageTriggerCriteria() *MessageTriggerCriteria {
	c := &MessageTriggerCriteria{}
	c.SetDefaults()
	return c
}

// SetDefaults applies default values to MessageTriggerCriteria
func (c *MessageTriggerCriteria) SetDefaults() {
	c.Enabled = true // Default to enabled
}

// RetryPolicy defines the retry policy for failed event triggers
type RetryPolicy struct {
	// Enabled determines if retry is enabled
	Enabled bool `json:"enabled" validate:"required"`

	// MaxRetries specifies the maximum number of retry attempts
	MaxRetries int `json:"maxRetries" validate:"min=0"`

	// RetryInterval defines the time to wait between retry attempts
	RetryInterval time.Duration `json:"retryInterval" validate:"gt=0"`

	// ExponentialBackoff determines if exponential backoff is used for retries
	ExponentialBackoff bool `json:"exponentialBackoff"`

	// MaxInterval is the maximum interval between retries when using exponential backoff
	MaxInterval time.Duration `json:"maxInterval,omitempty"`

	// RetryableErrors defines which errors should be retried
	RetryableErrors []string `json:"retryableErrors,omitempty"`
}

// NewRetryPolicy creates a new RetryPolicy with default values
func NewRetryPolicy() *RetryPolicy {
	rp := &RetryPolicy{}
	rp.SetDefaults()
	return rp
}

// SetDefaults applies default values for RetryPolicy
func (rp *RetryPolicy) SetDefaults() {
	rp.Enabled = false // Default to disabled

	if rp.MaxRetries == 0 {
		rp.MaxRetries = 3 // Default to 3 retries
	}
	if rp.RetryInterval == 0 {
		rp.RetryInterval = 1 * time.Minute // Default to 1 minute retry interval
	}
	if rp.MaxInterval == 0 {
		rp.MaxInterval = 1 * time.Hour // Default to 1 hour maximum interval
	}
}

// TriggerCriteria defines the settings and criteria for configuring triggers
type TriggerCriteria struct {
	// Polling specifies the configuration for polling-based triggers
	Polling *PollingTriggerCriteria `json:"pollingCriteria,omitempty"`

	// Event specifies the configuration for event-based triggers
	Event *EventTriggerCriteria `json:"eventCriteria,omitempty"`

	// Manual specifies the configuration for manual triggers
	Manual *ManualTriggerCriteria `json:"manualCriteria,omitempty"`

	// Webhook specifies the configuration for webhook triggers
	Webhook *WebhookTriggerCriteria `json:"webhookCriteria,omitempty"`

	// Schedule specifies the configuration for scheduled triggers
	Schedule *ScheduleTriggerCriteria `json:"scheduleCriteria,omitempty"`

	// API specifies the configuration for API triggers
	API *APITriggerCriteria `json:"apiCriteria,omitempty"`

	// Workflow specifies the configuration for workflow triggers
	Workflow *WorkflowTriggerCriteria `json:"workflowCriteria,omitempty"`

	// Message specifies the configuration for message triggers
	Message *MessageTriggerCriteria `json:"messageCriteria,omitempty"`
}

// NewTriggerCriteria creates a new TriggerCriteria with default values
func NewTriggerCriteria() *TriggerCriteria {
	ts := &TriggerCriteria{}
	ts.SetDefaults()
	return ts
}

// SetDefaults applies default values to TriggerCriteria
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
	if ts.API != nil {
		ts.API.SetDefaults()
	}
	if ts.Workflow != nil {
		ts.Workflow.SetDefaults()
	}
	if ts.Message != nil {
		ts.Message.SetDefaults()
	}
}

// SetDefaultByType applies default values to TriggerCriteria based on trigger type
func (ts *TriggerCriteria) SetDefaultByType(strategy TriggerType) {
	switch strategy {
	case TriggerTypeManual:
		ts.Manual = NewManualTriggerCriteria()
	case TriggerTypeEvent:
		ts.Event = NewEventTriggerCriteria()
	case TriggerTypePolling:
		ts.Polling = NewPollingTriggerCriteria()
	case TriggerTypeWebhook:
		ts.Webhook = NewWebhookTriggerCriteria()
	case TriggerTypeScheduled:
		ts.Schedule = NewScheduleTriggerCriteria()
	case TriggerTypeAPI:
		ts.API = NewAPITriggerCriteria()
	case TriggerTypeWorkflow:
		ts.Workflow = NewWorkflowTriggerCriteria()
	case TriggerTypeMessage:
		ts.Message = NewMessageTriggerCriteria()
	}
}

// TriggerSettings defines the settings for a trigger
type TriggerSettings struct {
	// Type defines the trigger type
	Type TriggerType `json:"strategy,omitempty" validate:"required,oneof=SCHEDULED EVENT POLLING WEBHOOK MANUAL API WORKFLOW MESSAGE BUTTON"`

	// Criteria contains the configuration settings for the trigger
	Criteria *TriggerCriteria `json:"criteria,omitempty"`

	// RetryPolicy specifies the retry behavior for failed triggers
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`

	// Description provides additional information about the trigger
	Description string `json:"description,omitempty"`

	// Tags provide searchable metadata for the trigger
	Tags []string `json:"tags,omitempty"`

	// Priority determines the execution priority when multiple triggers fire
	Priority int `json:"priority,omitempty"`
}

// SetDefaults applies default values to TriggerSettings
func (ts *TriggerSettings) SetDefaults() {
	if ts.Criteria == nil {
		ts.Criteria = NewTriggerCriteria()
	}

	if ts.Type == "" {
		ts.Type = TriggerTypeManual
	}

	if ts.RetryPolicy == nil {
		ts.RetryPolicy = NewRetryPolicy()
	}

	ts.Criteria.SetDefaultByType(ts.Type)
}

// NewTriggerSettings creates a new TriggerSettings with default values
func NewTriggerSettings() *TriggerSettings {
	ts := &TriggerSettings{}
	ts.SetDefaults()
	return ts
}

// SetDefaultByType applies default values to TriggerSettings based on trigger type
func (ts *TriggerSettings) SetDefaultByType(strategy TriggerType) {
	ts.Type = strategy
	ts.Criteria = NewTriggerCriteria()
	ts.Criteria.SetDefaultByType(strategy)

	// Only create retry policy for certain trigger types
	switch strategy {
	case TriggerTypeEvent, TriggerTypeWebhook, TriggerTypePolling, TriggerTypeAPI:
		ts.RetryPolicy = NewRetryPolicy()
	}
}
