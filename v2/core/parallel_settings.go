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
	"errors"
	"time"
)

// Error definitions
var (
	ErrInvalidConfiguration = errors.New("invalid configuration")
	ErrDuplicateBranchID    = errors.New("duplicate branch ID")
	ErrInvalidBranchID      = errors.New("invalid or empty branch ID")
)

// ParallelType defines the type of parallel execution.
type ParallelType string

const (
	// ParallelTypeAll executes all branches simultaneously and waits for all to complete
	ParallelTypeAll ParallelType = "all"

	// ParallelTypeRace executes all branches simultaneously but returns on first completion
	ParallelTypeRace ParallelType = "race"

	// ParallelTypeBatch executes branches in batches with configurable batch size
	ParallelTypeBatch ParallelType = "batch"

	// ParallelTypeThrottled executes branches with rate limiting
	ParallelTypeThrottled ParallelType = "throttled"
)

// ParallelErrorHandling defines how to handle errors in parallel execution.
type ParallelErrorHandling string

const (
	// ErrorHandlingFailFast stops all execution on first error
	ErrorHandlingFailFast ParallelErrorHandling = "fail_fast"

	// ErrorHandlingContinue continues execution despite errors
	ErrorHandlingContinue ParallelErrorHandling = "continue"

	// ErrorHandlingCollect collects all errors and fails at the end
	ErrorHandlingCollect ParallelErrorHandling = "collect"

	// ErrorHandlingPartial allows partial success with some failures
	ErrorHandlingPartial ParallelErrorHandling = "partial"
)

// ParallelResultHandling defines how to handle and aggregate results.
type ParallelResultHandling string

const (
	// ResultHandlingMerge merges all results into a single output
	ResultHandlingMerge ParallelResultHandling = "merge"

	// ResultHandlingArray collects results as an array
	ResultHandlingArray ParallelResultHandling = "array"

	// ResultHandlingMap collects results as a keyed map
	ResultHandlingMap ParallelResultHandling = "map"

	// ResultHandlingFirst returns only the first successful result
	ResultHandlingFirst ParallelResultHandling = "first"
)

// ParallelSettings defines the configuration for parallel flow execution.
type ParallelSettings struct {
	Type ParallelType `json:"type,omitempty"`

	// MaxConcurrency limits the maximum number of concurrent executions
	// 0 means unlimited (use with caution)
	MaxConcurrency int `json:"maxConcurrency,omitempty"`

	// BatchSize defines the number of branches to execute per batch (for batch type)
	BatchSize int `json:"batchSize,omitempty"`

	// ThrottleRate defines the rate limit for throttled execution (executions per second)
	ThrottleRate float64 `json:"throttleRate,omitempty"`

	// Timeout defines the maximum time to wait for all branches to complete
	Timeout *time.Duration `json:"timeout,omitempty"`

	// ErrorHandling defines how to handle errors during parallel execution
	ErrorHandling ParallelErrorHandling `json:"errorHandling,omitempty"`

	// ResultHandling defines how to aggregate and return results
	ResultHandling ParallelResultHandling `json:"resultHandling,omitempty"`

	// WaitForAll determines if we should wait for all branches to complete
	// (only applies to certain types)
	WaitForAll bool `json:"waitForAll,omitempty"`

	// FailOnError determines if the entire parallel execution should fail on any error
	FailOnError bool `json:"failOnError,omitempty"`

	// MinSuccessCount defines minimum successful branches required for overall success
	MinSuccessCount int `json:"minSuccessCount,omitempty"`

	// MaxFailureCount defines maximum failures allowed before stopping
	MaxFailureCount int `json:"maxFailureCount,omitempty"`

	// RetryPolicy defines retry behavior for failed branches
	RetryPolicy *ParallelRetryPolicy `json:"retryPolicy,omitempty"`

	// Priority defines execution priority for branches
	Priority ParallelPriority `json:"priority,omitempty"`

	// EnableLoadBalancing enables automatic load balancing across available resources
	EnableLoadBalancing bool `json:"enableLoadBalancing,omitempty"`

	// ResourceConstraints defines resource limits for parallel execution
	ResourceConstraints *ResourceConstraints `json:"resourceConstraints,omitempty"`

	// CollectMetrics determines if execution metrics should be collected
	CollectMetrics bool `json:"collectMetrics,omitempty"`

	// EnableTracing enables distributed tracing for parallel branches
	EnableTracing bool `json:"enableTracing,omitempty"`

	// BranchTimeout defines timeout for individual branches
	BranchTimeout *time.Duration `json:"branchTimeout,omitempty"`

	// Branches defines the parallel execution branches
	Branches []ParallelBranch `json:"branches,omitempty"`

	// PreCondition is an expression that must evaluate to true before parallel execution
	PreCondition string `json:"preCondition,omitempty"`

	// PostCondition is an expression to evaluate after parallel execution
	PostCondition string `json:"postCondition,omitempty"`

	// OnSuccess defines actions to take when parallel execution succeeds
	OnSuccess *ParallelCallback `json:"onSuccess,omitempty"`

	// OnFailure defines actions to take when parallel execution fails
	OnFailure *ParallelCallback `json:"onFailure,omitempty"`

	// OnPartialSuccess defines actions to take when some branches succeed
	OnPartialSuccess *ParallelCallback `json:"onPartialSuccess,omitempty"`
}

// ParallelBranch represents a single branch in parallel execution.
type ParallelBranch struct {
	// ID is the unique identifier of the branch
	ID string `json:"id,omitempty"`

	// Name is the display name of the branch
	Name string `json:"name,omitempty"`

	// Description describes the branch purpose
	Description string `json:"description,omitempty"`

	// Weight defines the relative importance/priority of this branch
	Weight int `json:"weight,omitempty"`

	// Condition is an optional condition that must be true to execute this branch
	Condition string `json:"condition,omitempty"`

	// NextStep is the ID of the step to execute in this branch
	NextStep string `json:"nextStep,omitempty"`

	// Timeout defines timeout specific to this branch
	Timeout *time.Duration `json:"timeout,omitempty"`

	// RetryPolicy defines retry behavior specific to this branch
	RetryPolicy *ParallelRetryPolicy `json:"retryPolicy,omitempty"`

	// Required indicates if this branch is required for overall success
	Required bool `json:"required,omitempty"`

	// Order defines execution order (for ordered parallel execution)
	Order int `json:"order,omitempty"`

	// ResourceRequirements defines resource needs for this branch
	ResourceRequirements map[string]interface{} `json:"resourceRequirements,omitempty"`

	// Tags for categorizing and filtering branches
	Tags []string `json:"tags,omitempty"`

	// Metadata for additional branch information
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// ParallelRetryPolicy defines retry behavior for parallel branches.
type ParallelRetryPolicy struct {
	// MaxAttempts defines maximum retry attempts
	MaxAttempts int `json:"maxAttempts,omitempty"`

	// BackoffType defines the backoff strategy
	BackoffType string `json:"backoffType,omitempty"` // linear, exponential, fixed

	// InitialDelay is the initial delay before first retry
	InitialDelay time.Duration `json:"initialDelay,omitempty"`

	// MaxDelay is the maximum delay between retries
	MaxDelay time.Duration `json:"maxDelay,omitempty"`

	// BackoffMultiplier for exponential backoff
	BackoffMultiplier float64 `json:"backoffMultiplier,omitempty"`

	// RetriableErrors defines which errors should trigger a retry
	RetriableErrors []string `json:"retriableErrors,omitempty"`

	// NonRetriableErrors defines which errors should not trigger a retry
	NonRetriableErrors []string `json:"nonRetriableErrors,omitempty"`
}

// ParallelPriority defines execution priority settings.
type ParallelPriority string

const (
	// PriorityHigh for high priority execution
	PriorityHigh ParallelPriority = "high"

	// PriorityNormal for normal priority execution
	PriorityNormal ParallelPriority = "normal"

	// PriorityLow for low priority execution
	PriorityLow ParallelPriority = "low"

	// PriorityWeighted for weight-based priority execution
	PriorityWeighted ParallelPriority = "weighted"
)

// ResourceConstraints defines resource limits for parallel execution.
type ResourceConstraints struct {
	// MaxMemory defines maximum memory usage (in bytes)
	MaxMemory int64 `json:"maxMemory,omitempty"`

	// MaxCPU defines maximum CPU usage (as percentage)
	MaxCPU float64 `json:"maxCPU,omitempty"`

	// MaxDiskIO defines maximum disk I/O usage
	MaxDiskIO int64 `json:"maxDiskIO,omitempty"`

	// MaxNetworkIO defines maximum network I/O usage
	MaxNetworkIO int64 `json:"maxNetworkIO,omitempty"`

	// MaxExecutionTime defines maximum total execution time
	MaxExecutionTime *time.Duration `json:"maxExecutionTime,omitempty"`
}

// ParallelCallback defines callback actions for parallel execution events.
type ParallelCallback struct {
	// Expression to evaluate
	Expression string `json:"expression,omitempty"`

	// NextStep to execute
	NextStep string `json:"nextStep,omitempty"`

	// NotificationSettings for sending notifications
	NotificationSettings map[string]interface{} `json:"notificationSettings,omitempty"`
}

// NewParallelSettings creates a new ParallelSettings with default values.
func NewParallelSettings() *ParallelSettings {
	return &ParallelSettings{
		Type:                ParallelTypeAll,
		MaxConcurrency:      10,
		BatchSize:           5,
		ThrottleRate:        10.0,
		ErrorHandling:       ErrorHandlingFailFast,
		ResultHandling:      ResultHandlingMerge,
		WaitForAll:          true,
		FailOnError:         true,
		MinSuccessCount:     1,
		MaxFailureCount:     0,
		Priority:            PriorityNormal,
		EnableLoadBalancing: false,
		CollectMetrics:      true,
		EnableTracing:       true,
		Branches:            make([]ParallelBranch, 0),
		ResourceConstraints: &ResourceConstraints{
			MaxMemory: 1024 * 1024 * 1024, // 1GB
			MaxCPU:    80.0,               // 80%
		},
	}
}

// AddBranch adds a new branch to the ParallelSettings.
func (ps *ParallelSettings) AddBranch(branch ParallelBranch) {
	ps.Branches = append(ps.Branches, branch)
}

// GetBranchByID retrieves a branch by its ID.
func (ps *ParallelSettings) GetBranchByID(id string) *ParallelBranch {
	for i, branch := range ps.Branches {
		if branch.ID == id {
			return &ps.Branches[i]
		}
	}
	return nil
}

// GetRequiredBranches returns all branches marked as required.
func (ps *ParallelSettings) GetRequiredBranches() []ParallelBranch {
	var required []ParallelBranch
	for _, branch := range ps.Branches {
		if branch.Required {
			required = append(required, branch)
		}
	}
	return required
}

// GetBranchesByTag returns branches that have any of the specified tags.
func (ps *ParallelSettings) GetBranchesByTag(tags ...string) []ParallelBranch {
	var matches []ParallelBranch
	tagMap := make(map[string]bool)
	for _, tag := range tags {
		tagMap[tag] = true
	}

	for _, branch := range ps.Branches {
		for _, branchTag := range branch.Tags {
			if tagMap[branchTag] {
				matches = append(matches, branch)
				break
			}
		}
	}
	return matches
}

// SetResourceConstraints sets resource constraints for the parallel execution.
func (ps *ParallelSettings) SetResourceConstraints(constraints *ResourceConstraints) {
	ps.ResourceConstraints = constraints
}

// SetRetryPolicy sets the default retry policy for all branches.
func (ps *ParallelSettings) SetRetryPolicy(policy *ParallelRetryPolicy) {
	ps.RetryPolicy = policy
}

// Validate validates the parallel settings configuration.
func (ps *ParallelSettings) Validate() error {
	if ps.MaxConcurrency < 0 {
		return ErrInvalidConfiguration // You'll need to define this error
	}

	if ps.BatchSize < 0 {
		return ErrInvalidConfiguration
	}

	if ps.ThrottleRate < 0 {
		return ErrInvalidConfiguration
	}

	if ps.MinSuccessCount < 0 {
		return ErrInvalidConfiguration
	}

	if ps.MaxFailureCount < 0 {
		return ErrInvalidConfiguration
	}

	// Validate branches
	branchIDs := make(map[string]bool)
	for _, branch := range ps.Branches {
		if branch.ID == "" {
			return ErrInvalidConfiguration
		}
		if branchIDs[branch.ID] {
			return ErrInvalidConfiguration // Duplicate branch ID
		}
		branchIDs[branch.ID] = true
	}

	return nil
}
