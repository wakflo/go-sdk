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

package sdk

import (
	"context"

	sdkcore "github.com/wakflo/go-sdk/core"
)

type OperationInfo struct {
	// Key holds the value of the "key" field.
	Name string `json:"name,omitempty" validate:"required"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty" validate:"required"`
	// Input holds the value of the "input" field.
	Input map[string]*sdkcore.AutoFormSchema `json:"input,omitempty"`
	// Auth holds the value of the "auth" field.
	Auth          *sdkcore.AutoFormSchema   `json:"auth"`
	SampleOutput  sdkcore.JSONObject        `json:"sample_output"`
	ErrorSettings sdkcore.StepErrorSettings `json:"error_settings" validate:"required"`
	RequireAuth   bool                      `json:"require_auth"`

	// Documentation represents the field used to store the connector's documentation in markdown.
	Documentation *string `json:"documentation,omitempty"`
}

// IOperation is an interface that represents an operation within a connector. It extends the IRunnable interface and provides a method to get information about the operation.
// GetInfo returns the OperationInfo metadata for the operation.
//
//	type IOperation interface {
//	    IRunnable
//	    GetInfo() *OperationInfo
type IOperation interface {
	IRunnable
	GetInfo() *OperationInfo
}

// CreateIOperationArgs represents the arguments for creating an IOperation.
// It contains the following fields:
// - Name: the name of the IOperation
// - Description: the description of the IOperation
// - Input
type CreateIOperationArgs struct {
	Name        string
	Description string
	Input       *map[string]interface{}
	Output      *map[string]interface{}

	Run                *func(ctx context.Context) (JSON, error)
	Auth               *func(ctx context.Context) (JSON, error)
	Test               *func(ctx context.Context) (JSON, error)
	GenerateSampleData *func(ctx context.Context) (JSON, error)
}
