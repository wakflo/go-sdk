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

package worker

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/wakflo/go-sdk/v2/core"
)

// HeartbeatStatus represents detailed status information for a worker heartbeat.
type HeartbeatStatus struct {
	// Status is the current operational status of the worker
	Status WorkerStatus `json:"status"`

	// CurrentTaskCount is the number of tasks currently being processed
	CurrentTaskCount int `json:"currentTaskCount"`

	// CPUUsage is the current CPU usage percentage (0-100)
	CPUUsage float64 `json:"cpuUsage"`

	// MemoryUsage is the current memory usage percentage (0-100)
	MemoryUsage float64 `json:"memoryUsage"`

	// DiskUsage is the current disk usage percentage (0-100)
	DiskUsage float64 `json:"diskUsage"`

	// NetworkLatency is the current network latency in milliseconds
	NetworkLatency int64 `json:"networkLatency"`

	// LastError contains details of the last error encountered, if any
	LastError string `json:"lastError,omitempty"`

	// TaskIDs contains the IDs of tasks currently being processed
	TaskIDs []string `json:"taskIds,omitempty"`

	// AdditionalMetrics contains any other metrics to report
	AdditionalMetrics map[string]interface{} `json:"additionalMetrics,omitempty"`
}

// HeartbeatOptions configures heartbeat behavior.
type HeartbeatOptions struct {
	// Interval is how often to send heartbeats
	Interval time.Duration `json:"interval"`

	// Timeout is how long before a worker is considered dead if no heartbeat is received
	Timeout time.Duration `json:"timeout"`

	// IncludeDetailedMetrics indicates whether to include detailed system metrics
	IncludeDetailedMetrics bool `json:"includeDetailedMetrics"`

	// RetryAttempts is the number of attempts to make when a heartbeat fails
	RetryAttempts int `json:"retryAttempts"`

	// RetryBackoff is the delay between retry attempts
	RetryBackoff time.Duration `json:"retryBackoff"`
}

// Heartbeat defines the interface for worker heartbeat management.
type Heartbeat interface {
	// Start begins sending regular heartbeats
	Start(ctx context.Context, workerID xid.ID, options HeartbeatOptions) error

	// Stop terminates heartbeat sending
	Stop(ctx context.Context, workerID xid.ID) error

	// SendHeartbeat sends a single heartbeat immediately
	SendHeartbeat(ctx context.Context, workerID xid.ID, status HeartbeatStatus) error

	// GetLastHeartbeat retrieves the last heartbeat information for a worker
	GetLastHeartbeat(ctx context.Context, workerID xid.ID) (*HeartbeatStatus, time.Time, error)

	// IsAlive checks if a worker is considered alive based on heartbeats
	IsAlive(ctx context.Context, workerID xid.ID) (bool, error)

	// GetDeadWorkers retrieves a list of workers that have failed to send heartbeats
	GetDeadWorkers(ctx context.Context) ([]xid.ID, error)

	// Logger returns a structured logger for heartbeat operations
	Logger() core.Logger
}
