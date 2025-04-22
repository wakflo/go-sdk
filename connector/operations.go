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
	sdkcore "github.com/wakflo/go-sdk/oldcore"
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

	// HelpText holds the value of the "helpText" field.
	HelpText *string `json:"helpText,omitempty"`

	Settings sdkcore.ActionSettings `json:"settings" validate:"required"`

	Type sdkcore.ActionType `json:"type" validate:"required,oneof=ACTION"`
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
