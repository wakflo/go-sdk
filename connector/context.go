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
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	valid "github.com/asaskevich/govalidator"
	"github.com/cavaliergopher/grab/v3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/wakflo/go-sdk/autoform"
	sdkcore "github.com/wakflo/go-sdk/core"
	"time"
)

// IRunnable is an interface that represents a runnable entity.
// Run executes the runnable entity with the provided context and run context, and returns the result as a JSON object or an error.
// Test runs a test of the runnable entity with the provided context, and returns the result as a JSON object or an error.
type IRunnable interface {
	// Run executes the runnable entity with the provided context and run context, and returns the result as a JSON object or an error.
	Run(ctx *RunContext) (Json, error)

	// Test runs a test of the runnable entity with the provided context, and returns the result as a JSON object or an error.
	Test(ctx *RunContext) (Json, error)
}

// PauseMetadata represents metadata for pausing execution.
//
// Type represents the type of pause, which can be either "DELAY" or "WEBHOOK".
// ResumeAt is an optional field that represents the time at which the execution should resume.
// RequestId is an optional field that represents the ID of the request associated with the pause.
// Response is an optional field that represents the response associated with the pause.
//
// Example usage:
//
//	metadata := PauseMetadata{
//	   Type:      sdkcore.PauseType("DELAY"),
//	   ResumeAt:  &time.Time{},
//	   RequestId: &"requestID",
//	   Response:  &any{},
//	}
//
// PauseMetadataFull is an example usage of PauseMetadata:
//
//	type PauseMetadataFull struct {
//	   PauseMetadata
//	}
//
// RunContext.PauseExecution is a method that accepts PauseMetadata as a parameter:
//
//	func (rctx *RunContext) PauseExecution(metadata PauseMetadata) (Json, error) {
//	   rctx.isPaused = true
//	   rctx.pausedTime = metadata.ResumeAt
//	   return &PauseMetadataFull{
//	       PauseMetadata: metadata,
//	   }, nil
//	}
//
// The RunContext structure has a method PauseExecution that accepts PauseMetadata and pauses the execution based on the metadata provided:
//
//	type RunContext struct {
//	   Step          *sdkcore.ConnectorStep       `json:"step"`
//	   Auth          *sdkcore.AuthContext         `json:"auth"`
//	   State         *sdkcore.StepsState          `json:"state"`
//	   Workflow      *sdkcore.WorkflowRunMetadata `json:"workflow"`
//	   Input         sdkcore.JsonObject           `json:"input"`
//	   ResolvedInput any                          `json:"resolvedInput"`
//	   LastRun       *time.Time                   `json:"lastRun"`
//	   Files         sdkcore.FileManager
//	   ctx           context.Context
//	   ExecutionType sdkcore.ExecutionType `json:"executionType"`
//	   TriggerType   sdkcore.TriggerHookType
//	   isPaused      bool
//	   pausedTime    *time.Time
//	}
//
// Json is an alias for any:
// type Json = any
//
// sdkcore.PauseType is a type that represents the pause type and has the following methods:
// - SqlTypeName() string
// - Values() []string
// - IsValid() bool
// - Validate() error
// - String() string
// - MarshalBinary() ([]byte, error)
// - UnmarshalBinary(text []byte) error
// - MarshalGQL(w io.Writer)
// - UnmarshalGQL(value interface{}) error
// - MarshalJSON() ([]byte, error)
// - UnmarshalJSON(data []byte) error
// - Scan(value interface{}) error
// - MarshalText() ([]byte, error)
// - UnmarshalText(text []byte) error
// - MarshalYAML() (interface{}, error)
// - UnmarshalYAML(unmarshal func(interface{}) error) error
type PauseMetadata struct {
	Type      sdkcore.PauseType `json:"type"`
	ResumeAt  *time.Time        `json:"resumeAt"`
	RequestId *string           `json:"requestId"`
	Response  *any              `json:"response"`
}

// PauseMetadataFull is a type that extends the PauseMetadata struct.
// PauseMetadata is a struct that contains information about a pause execution, including the pause type, resume time, request ID, and response.
//
//	type PauseMetadata struct {
//		Type      sdkcore.PauseType `json:"type"`
//		ResumeAt  *time.Time        `json:"resumeAt"`
//		RequestId *string           `json:"requestId"`
//		Response  *any              `json:"response"`
//	}
//
// PauseMetadataFull extends the PauseMetadata struct.
//
//	type PauseMetadataFull struct {
//		PauseMetadata
//	}
//
// This type is used in the RunContext.PauseExecution method to represent the metadata of a pause execution.
// Usage example:
//
//	func (rctx *RunContext) PauseExecution(metadata PauseMetadata) (Json, error) {
//		rctx.isPaused = true
//		rctx.pausedTime = metadata.ResumeAt
//		return &PauseMetadataFull{
//			PauseMetadata: metadata,
//		}, nil
//	}
type PauseMetadataFull struct {
	PauseMetadata
}

// TriggerRunContext is a struct that represents the context for triggering a run.
type TriggerRunContext struct {
	Metadata *sdkcore.WorkflowRunMetadata `json:"metadata"`
	Auth     *sdkcore.AuthContext         `json:"auth"`
	Trigger  *sdkcore.Trigger
}

// RunContext represents the context in which a workflow step is executed.
//
// It contains information about the workflow, the current step, the connector version,
// the workflow instance, authentication data, and the state of the workflow steps.
type RunContext struct {
	Step          *sdkcore.ConnectorStep       `json:"step"`
	Auth          *sdkcore.AuthContext         `json:"auth"`
	State         *sdkcore.StepsState          `json:"state"`
	Workflow      *sdkcore.WorkflowRunMetadata `json:"workflow"`
	Input         sdkcore.JsonObject           `json:"input"`
	ResolvedInput any                          `json:"resolvedInput"`
	LastRun       *time.Time                   `json:"lastRun"`
	Files         sdkcore.FileManager
	ctx           context.Context
	ExecutionType sdkcore.ExecutionType `json:"executionType"`
	TriggerType   sdkcore.TriggerHookType
	isPaused      bool
	pausedTime    *time.Time
}

func (rctx *RunContext) SetContext(ctx context.Context) {
	rctx.ctx = ctx
}

// IsPaused returns a boolean value indicating whether the execution is currently paused.
// It checks the value of the 'isPaused' field in the RunContext struct.
func (rctx *RunContext) IsPaused() bool {
	return rctx.isPaused
}

// PauseExecution pauses the execution of the RunContext.
// It sets the isPaused field of the RunContext to true and the pausedTime field to the provided resume time from the PauseMetadata.
// It returns a pointer to a PauseMetadataFull object and nil error.
func (rctx *RunContext) PauseExecution(metadata PauseMetadata) (Json, error) {
	rctx.isPaused = true
	rctx.pausedTime = metadata.ResumeAt
	return &PauseMetadataFull{
		PauseMetadata: metadata,
	}, nil
}

// GetPauseTime returns the paused time in the RunContext.
// It retrieves the value from the pausedTime field in the RunContext struct.
func (rctx *RunContext) GetPauseTime() *time.Time {
	return rctx.pausedTime
}

// InputToType returns a pointer to a value of type T by marshaling and unmarshaling the ResolvedInput field of the provided RunContext struct.
// If there is an error during the marshaling or unmarshaling process, nil is returned.
// The function signature is as follows:
func InputToType[T any](ctx *RunContext) *T {
	b, err := json.Marshal(ctx.ResolvedInput)
	if err != nil {
		return nil
	}

	var rsp T
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil
	}

	return &rsp
}

// DynamicInputToType converts the resolved input of type `sdkcore.DynamicOptionsContext` to the desired type T.
// It uses JSON marshaling and unmarshalling to perform the conversion.
// If any error occurs during marshaling or unmarshaling, it returns nil.
// The function returns a pointer to the converted value of type T.
func DynamicInputToType[T any](ctx *sdkcore.DynamicOptionsContext) *T {
	b, err := json.Marshal(ctx.ResolvedInput)
	if err != nil {
		return nil
	}

	var rsp T
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil
	}

	return &rsp
}

// StringToFile converts a file string to a *autoform.File object.
//
// The function checks if the file string is a base64-encoded data or a URL. If the file string is base64-encoded data, it decodes the data and assigns it to the `Data` field of the
func StringToFile(fileStr string) (*autoform.File, error) {
	file := &autoform.File{}

	if valid.IsBase64(fileStr) {
		data, err := base64.StdEncoding.DecodeString(fileStr)
		if err != nil {
			return nil, err
		}

		mime := mimetype.Detect(data)
		file.Data = bytes.NewReader(data)
		file.Extension = mime.Extension()
		file.Mime = mime.String()

		return file, nil
	}

	if valid.IsURL(fileStr) {
		data, err := DownloadFile(fileStr)
		if err != nil {
			return nil, err
		}

		bt, err := data.Bytes()
		if err != nil {
			return nil, err
		}

		mime := mimetype.Detect(bt)
		file.Data = bytes.NewReader(bt)
		file.Extension = mime.Extension()
		file.Size = data.Size()
		file.Name = data.Filename
		file.Mime = mime.String()

		return file, nil
	}

	return nil, errors.New("invalid file string")
}

// DownloadFile downloads a file from the specified URL using the grab package.
// It returns the grab.Response object and an error if any.
func DownloadFile(url string) (*grab.Response, error) {
	resp, err := grab.Get(".", url)
	if err != nil {
		return nil, err
	}

	resp.Wait()

	return resp, nil
}
