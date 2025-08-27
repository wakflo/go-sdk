package sdk

import (
	"github.com/wakflo/go-sdk/v2/context"
	"github.com/wakflo/go-sdk/v2/core"
	"github.com/wakflo/go-sdk/v2/worker"
)

type JSON = core.JSON
type JSONObject = core.JSONObject
type OffsetPaginationMeta = core.OffsetPaginationMeta
type Logger = core.Logger
type LogEntry = core.LogEntry
type LogLine = core.LogLine
type RetryPolicy = core.RetryPolicy
type TriggerCriteria = core.TriggerCriteria
type TriggerSettings = core.TriggerSettings
type ActionSettings = core.ActionSettings
type RouterSettings = core.RouterSettings
type BranchSettings = core.BranchSettings
type FlowRouter = core.FlowRouter
type APITriggerCriteria = core.APITriggerCriteria
type AuthMetadata = core.AuthMetadata
type DynamicOptionsFilterParams = core.DynamicOptionsFilterParams
type DynamicOptionsResponse = core.DynamicOptionsResponse
type EventTriggerCriteria = core.EventTriggerCriteria
type IntegrationBuildMetadata = core.IntegrationBuildMetadata
type IntegrationLanguage = core.IntegrationLanguage
type IntegrationPlatform = core.IntegrationPlatform
type LogBuilder = core.LogBuilder
type LogSink = core.LogSink
type ManualTriggerCriteria = core.ManualTriggerCriteria
type MessageTriggerCriteria = core.MessageTriggerCriteria
type PollingTriggerCriteria = core.PollingTriggerCriteria
type LogLevel = core.LogLevel
type LogLineLevel = core.LogLineLevel
type NoopLogger = core.NoopLogger
type ScheduleTriggerCriteria = core.ScheduleTriggerCriteria
type SystemActivityLog = core.SystemActivityLog
type SystemActivityLogs = core.SystemActivityLogs
type WebhookTriggerCriteria = core.WebhookTriggerCriteria
type WorkflowTriggerCriteria = core.WorkflowTriggerCriteria
type WriteLogLineOpts = core.WriteLogLineOpts

type Environment = core.Environment

const (
	// EnvironmentTest represents the test environment
	EnvironmentTest = core.EnvironmentTest
	// EnvironmentDebug represents the debug environment
	EnvironmentDebug = core.EnvironmentDebug
	// EnvironmentProd represents the production environment
	EnvironmentProd = core.EnvironmentProd
	EnvironmentDev  = core.EnvironmentDev
)

type FlowComponentType = core.FlowComponentType

const (
	// FlowComponentTypeBranch represents a branching flow component
	FlowComponentTypeBranch = core.FlowComponentTypeBranch

	// FlowComponentTypeLoop represents a loop flow component
	FlowComponentTypeLoop = core.FlowComponentTypeLoop

	// FlowComponentTypeCondition represents a condition flow component
	FlowComponentTypeCondition = core.FlowComponentTypeCondition

	// FlowComponentTypeRouter represents a router flow component
	FlowComponentTypeRouter = core.FlowComponentTypeRouter

	// FlowComponentTypeApproval represents an approval flow component
	FlowComponentTypeApproval = core.FlowComponentTypeApproval

	// FlowComponentTypeDelay represents a delay flow component
	FlowComponentTypeDelay = core.FlowComponentTypeDelay

	// FlowComponentTypeSequential represents a delay flow component
	FlowComponentTypeSequential = core.FlowComponentTypeSequential

	// FlowComponentTypeSubflow represents a subflow flow component
	FlowComponentTypeSubflow = core.FlowComponentTypeSubflow

	FlowComponentTypeParallel = core.FlowComponentTypeParallel
)

type AuthType = core.AuthType

const (
	// None indicates no authentication is required
	None = core.None

	// Basic indicates username and password authentication
	Basic = core.Basic

	// Secret indicates a shared secret or token authentication
	Secret = core.Secret

	// APIKey indicates API key authentication
	APIKey = core.APIKey

	// OAuth2 indicates OAuth 2.0 authentication
	OAuth2 = core.OAuth2

	// Custom indicates a custom authentication method
	Custom = core.Custom

	// JWT indicates JSON Web Token authentication
	JWT = core.JWT

	// ApiKeyHeader indicates an API key in the header
	ApiKeyHeader = core.ApiKeyHeader

	// ApiKeyQuery indicates an API key in the query string
	ApiKeyQuery = core.ApiKeyQuery

	// BearerToken indicates a bearer token authentication
	BearerToken = core.BearerToken

	// ClientCert indicates client certificate authentication
	ClientCert = core.ClientCert
)

type RouterType = core.RouterType

const (
	// RouterTypeSwitch routes based on evaluating an expression and matching the result.
	RouterTypeSwitch = core.RouterTypeSwitch

	// RouterTypeCondition routes based on evaluating multiple conditions and taking the first match.
	RouterTypeCondition = core.RouterTypeCondition

	// RouterTypeMultiPath allows taking multiple paths if their conditions are true.
	RouterTypeMultiPath = core.RouterTypeMultiPath
)

type BranchMode = core.RouteMode

const (
	// BranchModeExpression uses a JavaScript expression to determine the branch
	BranchModeExpression = core.RouteModeExpression

	// BranchModeCondition uses field-based condition matching
	BranchModeCondition = core.RouteModeCondition

	// BranchModeValue uses direct value matching
	BranchModeValue = core.RouteModeValue
)

type ActionType = core.ActionType

const (
	// ActionTypeAction represents a standard action.
	ActionTypeAction = core.ActionTypeAction

	// ActionTypeBranch represents a branch action that can conditionally execute different paths.
	ActionTypeBranch = core.ActionTypeBranch

	// ActionTypeBoolean represents a boolean-conditional action.
	ActionTypeBoolean = core.ActionTypeBoolean

	// ActionTypeLoop represents a loop action.
	ActionTypeLoop = core.ActionTypeLoop

	// ActionTypeRouter represents a router action that can direct flow based on conditions.
	ActionTypeRouter = core.ActionTypeRouter

	ActionTypeApproval = core.ActionTypeApproval

	ActionTypeDelay = core.ActionTypeDelay

	ActionTypeSubflow = core.ActionTypeSubflow

	ActionTypeParallel = core.ActionTypeParallel
)

type TriggerType = core.TriggerType

const (
	// TriggerTypeScheduled indicates a workflow triggered by a schedule
	TriggerTypeScheduled = core.TriggerTypeScheduled

	// TriggerTypeEvent indicates a workflow triggered by an event
	TriggerTypeEvent = core.TriggerTypeEvent

	TriggerTypePubsub = core.TriggerTypePubsub

	// TriggerTypePolling indicates a workflow triggered by polling for changes
	TriggerTypePolling = core.TriggerTypePolling

	// TriggerTypeWebhook indicates a workflow triggered by a webhook
	TriggerTypeWebhook = core.TriggerTypeWebhook

	// TriggerTypeManual indicates a workflow triggered manually by a user
	TriggerTypeManual = core.TriggerTypeManual

	// TriggerTypeAPI indicates a workflow triggered via the API
	TriggerTypeAPI = core.TriggerTypeAPI

	// TriggerTypeWorkflow indicates a workflow triggered by another workflow
	TriggerTypeWorkflow = core.TriggerTypeWorkflow

	// TriggerTypeMessage indicates a workflow triggered by a message
	TriggerTypeMessage = core.TriggerTypeMessage

	// TriggerTypeButton indicates a workflow triggered by a button click
	TriggerTypeButton = core.TriggerTypeButton
)

type AuthContext = context.AuthContext
type ExecuteContext = context.ExecuteContext
type BaseContext = context.BaseContext
type LifecycleContext = context.LifecycleContext
type PerformContext = context.PerformContext
type DynamicFieldContext = context.DynamicFieldContext

type Heartbeat = worker.Heartbeat
type HeartbeatOptions = worker.HeartbeatOptions
type HeartbeatStatus = worker.HeartbeatStatus
type Specialization = worker.Specialization
type SpecializationCapability = worker.SpecializationCapability
type SpecializationMatch = worker.SpecializationMatch
type Registration = worker.Registration
type SpecializationRequirement = worker.SpecializationRequirement
type RegistrationInfo = worker.RegistrationInfo
type WorkerMetadata = worker.WorkerMetadata

type SpecializationType = worker.SpecializationType

const (
	SpecializationTypeIntegration SpecializationType = worker.SpecializationTypeIntegration
	SpecializationTypeWorkflow    SpecializationType = worker.SpecializationTypeWorkflow
	SpecializationTypeRegion      SpecializationType = worker.SpecializationTypeRegion
	SpecializationTypeProject     SpecializationType = worker.SpecializationTypeProject
)

type WorkerStatus = worker.WorkerStatus

const (
	// WorkerStatusIdle indicates the worker is available but not processing any tasks.
	WorkerStatusIdle WorkerStatus = worker.WorkerStatusIdle

	// WorkerStatusBusy indicates the worker is currently processing tasks.
	WorkerStatusBusy WorkerStatus = worker.WorkerStatusBusy

	WorkerStatusOffline = worker.WorkerStatusOffline

	// WorkerStatusStarting indicates the worker is in the process of starting up.
	WorkerStatusStarting WorkerStatus = worker.WorkerStatusStarting

	// WorkerStatusStopping indicates the worker is in the process of shutting down.
	WorkerStatusStopping WorkerStatus = worker.WorkerStatusStopping

	// WorkerStatusError indicates the worker has encountered an error.
	WorkerStatusError WorkerStatus = worker.WorkerStatusError

	// WorkerStatusMaintenance indicates the worker is in maintenance mode.
	WorkerStatusMaintenance WorkerStatus = worker.WorkerStatusMaintenance

	// WorkerStatusOnline indicates the worker is available and actively connected.
	WorkerStatusOnline WorkerStatus = worker.WorkerStatusOnline

	// WorkerStatusDraining indicates the worker is completing current tasks and not accepting new ones.
	WorkerStatusDraining WorkerStatus = worker.WorkerStatusDraining
)

// StepRunStatus represents the status of a step run.
type StepRunStatus = core.StepRunStatus

// Enum values for StepRunStatus.
const (
	// StepRunStatusPending indicates a step is queued to run but hasn't started
	StepRunStatusPending StepRunStatus = core.StepRunStatusPending

	// StepRunStatusPaused indicates a step is paused and waiting for manual continuation
	StepRunStatusPaused StepRunStatus = core.StepRunStatusPaused

	// StepRunStatusRunning indicates a step is currently executing
	StepRunStatusRunning StepRunStatus = core.StepRunStatusRunning

	// StepRunStatusCompleted indicates a step has completed successfully
	StepRunStatusCompleted StepRunStatus = core.StepRunStatusCompleted

	// StepRunStatusFailed indicates a step has failed
	StepRunStatusFailed StepRunStatus = core.StepRunStatusFailed

	// StepRunStatusCancelled indicates a step was manually cancelled
	StepRunStatusCancelled StepRunStatus = core.StepRunStatusCancelled

	// StepRunStatusSkipped indicates a step was skipped due to conditions or branching
	StepRunStatusSkipped StepRunStatus = core.StepRunStatusSkipped

	// StepRunStatusTimeout indicates a step exceeded its allowed execution time
	StepRunStatusTimeout StepRunStatus = core.StepRunStatusTimeout

	// StepRunStatusWaiting indicates a step is waiting for an external event or condition
	StepRunStatusWaiting StepRunStatus = core.StepRunStatusWaiting

	// StepRunStatusBlocked indicates a step is blocked by a dependency or condition
	StepRunStatusBlocked StepRunStatus = core.StepRunStatusBlocked

	// StepRunStatusApproved indicates a step has been manually approved
	StepRunStatusApproved StepRunStatus = core.StepRunStatusApproved

	// StepRunStatusRejected indicates a step was manually rejected
	StepRunStatusRejected StepRunStatus = core.StepRunStatusRejected

	StepRunStatusRetrying StepRunStatus = core.StepRunStatusRetrying
)
