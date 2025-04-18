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

package connector

import (
	"context"

	sdkcore "github.com/wakflo/go-sdk/oldcore"
	"github.com/wakflo/go-sdk/sdk"
)

// TriggerInfo represents the information of a trigger.
// It includes the name, description, input schema, authentication schema,
// sample output, error settings, and whether authentication is required.
type TriggerInfo struct {
	// Key holds the value of the "key" field.
	Name string `json:"name,omitempty" validate:"required"`
	// Type holds the value of the "description" field.
	Type sdkcore.TriggerType `json:"type,omitempty" validate:"required,oneof=SCHEDULED EVENT POLLING WEBHOOK"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty" validate:"required"`
	// Input holds the value of the "input" field.
	Input map[string]*sdkcore.AutoFormSchema `json:"input,omitempty"`
	// Auth holds the value of the "auth" field.
	Auth *sdkcore.AutoFormSchema `json:"auth"`

	SampleOutput sdkcore.JSONObject `json:"sample_output"`

	Settings *sdkcore.OldTriggerSettings `json:"settings" validate:"required"`

	ErrorSettings *sdkcore.StepErrorSettings `json:"error_settings" validate:"required"`

	// Documentation represents the field used to store the connector's documentation in markdown.
	Documentation *string `json:"documentation,omitempty"`

	RequireAuth bool `json:"requireAuth"`

	// HelpText holds the value of the "helpText" field.
	HelpText *string `json:"helpText,omitempty"`
}

type ITrigger interface {
	IRunnable

	GetInfo() *TriggerInfo

	OnEnabled(ctx *RunContext) error

	OnDisabled(ctx *RunContext) error
}

type CreateTriggerArgs struct {
	Name        string
	Description string
	Input       *map[string]interface{}
	Output      *map[string]interface{}

	Run                *func(ctx context.Context) (sdk.JSON, error)
	Auth               *func(ctx context.Context) (sdk.JSON, error)
	Test               *func(ctx context.Context) (sdk.JSON, error)
	GenerateSampleData *func(ctx context.Context) (sdk.JSON, error)
}

func CreateTrigger(args CreateTriggerArgs) {
}
