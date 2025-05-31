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
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/wakflo/go-sdk/v2/core"
)

// LifecycleContext defines the interface for trigger lifecycle management.
// It provides methods for handling the setup, teardown, and lifecycle state
// of triggers within a workflow.
type LifecycleContext interface {
	// Context returns the underlying Go context for the lifecycle operation.
	Context() context.Context

	// WorkflowID returns the unique identifier of the workflow.
	WorkflowID() xid.ID

	// ProjectID returns the unique identifier of the project.
	ProjectID() xid.ID

	// TriggerID returns the unique identifier of the trigger being managed.
	TriggerID() string

	// Logger returns a structured logger for the trigger.
	Logger() core.Logger

	// Config returns the configuration for the trigger.
	Config() map[string]interface{}

	// Input returns the Input for the trigger.
	Input() map[string]interface{}

	// GetLastRunTime returns the timestamp of the last successful run.
	GetLastRunTime() (*time.Time, error)

	// SetLastRunTime updates the timestamp of the last successful run.
	SetLastRunTime(time.Time) error

	// GetState retrieves stored state for stateful triggers.
	GetState() (map[string]interface{}, error)

	// SetState stores state for stateful triggers.
	SetState(map[string]interface{}) error

	// TriggerCriteria returns the configured criteria for the trigger.
	TriggerCriteria() (*core.TriggerCriteria, error)

	// EmitEvent allows triggers to emit events that can be processed by the workflow runtime.
	EmitEvent(payload core.JSON) error

	// StoreMetadata allows triggers to store metadata about their execution.
	StoreMetadata(key string, value interface{}) error

	// GetMetadata retrieves stored metadata.
	GetMetadata(key string) (interface{}, error)

	// Cancel signals that the trigger should be canceled.
	Cancel() error

	// IsCanceled checks if the trigger has been canceled.
	IsCanceled() bool
}
