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

// WorkerStatus represents the current operational status of a worker.
type WorkerStatus string

const (
	// WorkerStatusIdle indicates the worker is available but not processing any tasks.
	WorkerStatusIdle WorkerStatus = "idle"

	// WorkerStatusBusy indicates the worker is currently processing tasks.
	WorkerStatusBusy WorkerStatus = "busy"

	// WorkerStatusStarting indicates the worker is in the process of starting up.
	WorkerStatusStarting WorkerStatus = "starting"

	// WorkerStatusStopping indicates the worker is in the process of shutting down.
	WorkerStatusStopping WorkerStatus = "stopping"

	// WorkerStatusError indicates the worker has encountered an error.
	WorkerStatusError WorkerStatus = "error"

	// WorkerStatusMaintenance indicates the worker is in maintenance mode.
	WorkerStatusMaintenance WorkerStatus = "maintenance"

	// WorkerStatusOffline indicates the worker is currently not operational or unreachable.
	WorkerStatusOffline WorkerStatus = "offline"

	// WorkerStatusOnline indicates the worker is available and actively connected.
	WorkerStatusOnline WorkerStatus = "online"

	// WorkerStatusDraining indicates the worker is completing current tasks and not accepting new ones.
	WorkerStatusDraining WorkerStatus = "draining"
)

func (s WorkerStatus) String() string {
	return string(s)
}

func (s WorkerStatus) Values() []string {
	return []string{
		string(WorkerStatusIdle),
		string(WorkerStatusBusy),
		string(WorkerStatusStarting),
		string(WorkerStatusStopping),
		string(WorkerStatusError),
		string(WorkerStatusMaintenance),
		string(WorkerStatusOffline),
		string(WorkerStatusOnline),
		string(WorkerStatusDraining),
	}
}

func (WorkerStatus) SQLTypeName() string {
	return "worker_status"
}

// WorkerMetadata contains additional metadata about a worker.
type WorkerMetadata struct {
	// Hostname is the host where the worker is running
	Hostname string `json:"hostname"`

	// IPAddress is the IP address of the worker
	IPAddress string `json:"ipAddress,omitempty"`

	// Version is the worker software version
	Version string `json:"version"`

	// StartTime is when the worker started
	StartTime time.Time `json:"startTime"`

	// OperatingSystem contains information about the OS
	OperatingSystem string `json:"operatingSystem,omitempty"`

	// Tags contains arbitrary tags for worker classification
	Tags []string `json:"tags,omitempty"`

	// AdditionalInfo contains any other worker information
	AdditionalInfo map[string]interface{} `json:"additionalInfo,omitempty"`
}

// RegistrationInfo represents the information required to register a worker.
type RegistrationInfo struct {
	// ID is the unique identifier for this worker
	ID xid.ID `json:"id"`

	// DisplayName is a human-readable name for this worker
	DisplayName string `json:"displayName"`

	// Status is the current operational status of the worker
	Status WorkerStatus `json:"status"`

	// Capabilities defines what this worker can do
	Capabilities []string `json:"capabilities"`

	// Specializations defines what this worker specializes in
	Specializations []string `json:"specializations,omitempty"`

	// MaxConcurrentTasks is the maximum number of tasks this worker can process simultaneously
	MaxConcurrentTasks int `json:"maxConcurrentTasks"`

	// CurrentTasks is the number of tasks currently being processed
	CurrentTasks int `json:"currentTasks"`

	// Metadata contains additional worker information
	Metadata WorkerMetadata `json:"metadata"`
}

// Registration defines the interface for worker registration.
type Registration interface {
	// Register registers the worker with the workflow system
	Register(ctx context.Context, info RegistrationInfo) error

	// Unregister removes the worker from the workflow system
	Unregister(ctx context.Context, workerID xid.ID) error

	// UpdateStatus updates the worker's operational status
	UpdateStatus(ctx context.Context, workerID xid.ID, status WorkerStatus) error

	// UpdateCapabilities updates the worker's capabilities
	UpdateCapabilities(ctx context.Context, workerID xid.ID, capabilities []string) error

	// UpdateMetadata updates the worker's metadata
	UpdateMetadata(ctx context.Context, workerID xid.ID, metadata WorkerMetadata) error

	// GetWorkerInfo retrieves the current worker information
	GetWorkerInfo(ctx context.Context, workerID xid.ID) (*RegistrationInfo, error)

	// Logger returns a structured logger for registration operations
	Logger() core.Logger
}
