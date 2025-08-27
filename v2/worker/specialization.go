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

	"github.com/rs/xid"
	"github.com/wakflo/go-sdk/v2/core"
)

// SpecializationType represents a category of worker specialization.
type SpecializationType string

const (
	// SpecializationTypeIntegration indicates specialization in specific integrations.
	SpecializationTypeIntegration SpecializationType = "integration"

	// SpecializationTypeProject indicates specialization in specific projects.
	SpecializationTypeProject SpecializationType = "project"

	// SpecializationTypeWorkflow indicates specialization in specific workflows.
	SpecializationTypeWorkflow SpecializationType = "workflow"

	// SpecializationTypeResource indicates specialization in resource-intensive operations.
	SpecializationTypeResource SpecializationType = "resource"

	// SpecializationTypeRegion indicates specialization in a geographic region.
	SpecializationTypeRegion SpecializationType = "region"
)

func (s SpecializationType) Values() []string {
	return []string{
		string(SpecializationTypeIntegration),
		string(SpecializationTypeProject),
		string(SpecializationTypeWorkflow),
		string(SpecializationTypeResource),
		string(SpecializationTypeRegion),
	}
}

func (s SpecializationType) String() string {
	return string(s)
}

// SpecializationRequirement defines requirements for a task to be processed.
type SpecializationRequirement struct {
	// Type is the category of specialization required
	Type SpecializationType `json:"type"`

	// Value is the specific value required (e.g., integration name)
	Value string `json:"value"`

	// Priority indicates the importance of this requirement (higher = more important)
	Priority int `json:"priority"`

	// IsRequired indicates if this specialization is mandatory
	IsRequired bool `json:"isRequired"`

	// Description provides additional context about this requirement
	Description string `json:"description,omitempty"`
}

// SpecializationCapability represents a worker's capability to handle specific tasks.
type SpecializationCapability struct {
	// Type is the category of specialization offered
	Type SpecializationType `json:"type"`

	// Value is the specific capability offered (e.g., integration name)
	Value string `json:"value"`

	// Confidence indicates the worker's proficiency in this capability (0.0-1.0)
	Confidence float64 `json:"confidence"`

	// MaxConcurrent is the maximum number of concurrent tasks of this type
	MaxConcurrent int `json:"maxConcurrent"`

	// Description provides details about this capability
	Description string `json:"description,omitempty"`
}

// SpecializationMatch represents how well a worker matches task requirements.
type SpecializationMatch struct {
	// WorkerID is the worker that matches the requirements
	WorkerID xid.ID `json:"workerId"`

	// MatchScore indicates how well the worker matches (higher is better)
	MatchScore float64 `json:"matchScore"`

	// MissingRequirements contains any required specializations the worker lacks
	MissingRequirements []SpecializationRequirement `json:"missingRequirements,omitempty"`

	// MatchDetails provides details about specific requirement matches
	MatchDetails map[string]float64 `json:"matchDetails,omitempty"`
}

// Specialization defines the interface for worker specialization.
type Specialization interface {
	// RegisterCapabilities registers a worker's specialization capabilities
	RegisterCapabilities(ctx context.Context, workerID xid.ID, capabilities []SpecializationCapability) error

	// UpdateCapabilities updates a worker's specialization capabilities
	UpdateCapabilities(ctx context.Context, workerID xid.ID, capabilities []SpecializationCapability) error

	// GetCapabilities retrieves a worker's specialization capabilities
	GetCapabilities(ctx context.Context, workerID xid.ID) ([]SpecializationCapability, error)

	// FindMatchingWorkers finds workers that match specific requirements
	FindMatchingWorkers(ctx context.Context, requirements []SpecializationRequirement) ([]SpecializationMatch, error)

	// CalculateMatchScore determines how well a worker matches requirements
	CalculateMatchScore(capabilities []SpecializationCapability, requirements []SpecializationRequirement) (*SpecializationMatch, error)

	// Logger returns a structured logger for specialization operations
	Logger() core.Logger
}
