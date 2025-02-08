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

package integration

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/cavaliergopher/grab/v3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/wakflo/go-sdk/autoform"
	sdkcore "github.com/wakflo/go-sdk/core"
)

type ExecutionMode = sdkcore.ExecutionMode

type BaseContext struct {
	ctx           context.Context
	Auth          *sdkcore.AuthContext `json:"auth"`
	rawInput      map[string]any
	input         any
	metadata      *ExecuteMetadata
	log           sdkcore.Logger
	Files         sdkcore.FileManager
	ExecutionMode ExecutionMode `json:"execution_mode"`
}

// Metadata returns a boolean value indicating whether the execution is currently paused.
// It checks the value of the 'isPaused' field in the RunContext struct.
func (r *BaseContext) Metadata() ExecuteMetadata {
	return *r.metadata
}

// GetRawInput returns a boolean value indicating whether the execution is currently paused.
// It checks the value of the 'isPaused' field in the RunContext struct.
func (r *BaseContext) GetRawInput() sdkcore.JSONObject {
	return r.rawInput
}

// Input returns a boolean value indicating whether the execution is currently paused.
// It checks the value of the 'isPaused' field in the RunContext struct.
func (r *BaseContext) Input() any {
	return r.input
}
func (r *BaseContext) RawInput() map[string]any {
	return r.rawInput
}

func (r *BaseContext) Log() sdkcore.Logger {
	return r.log
}

func NewBaseContext(
	ctx context.Context,
	files sdkcore.FileManager,
	meta *ExecuteMetadata,
	auth *sdkcore.AuthContext,
	resolvedInput sdkcore.JSON,
	input map[string]any,
	log sdkcore.Logger,
) *BaseContext {
	return &BaseContext{
		ctx:           ctx,
		Auth:          auth,
		Files:         files,
		input:         resolvedInput,
		ExecutionMode: sdkcore.ExecutionModeLive,
		metadata:      meta,
		log:           log,
		rawInput:      input,
	}
}

type LifecycleContext struct {
	BaseContext
	Ack  func(output sdkcore.JSON) error
	Nack func(err error) error
}

type ExecuteContext struct {
	BaseContext
}

type PerformContext struct {
	BaseContext
}

func NewExecuteContext(
	ctx context.Context,
	files sdkcore.FileManager,
	meta *ExecuteMetadata,
	auth *sdkcore.AuthContext,
	resolvedInput sdkcore.JSON,
	input map[string]any,
	log sdkcore.Logger,
) *ExecuteContext {
	return &ExecuteContext{
		BaseContext: *NewBaseContext(
			ctx,
			files,
			meta,
			auth,
			resolvedInput,
			input,
			log,
		),
	}
}

// IsPaused returns a boolean value indicating whether the execution is currently paused.
// It checks the value of the 'isPaused' field in the RunContext struct.
func (r *ExecuteContext) IsPaused() bool {
	return true
}

// PauseExecution pauses the execution of the RunContext.
// It sets the isPaused field of the RunContext to true and the pausedTime field to the provided resume time from the PauseMetadata.
// It returns a pointer to a PauseMetadataFull object and nil error.
func (r *ExecuteContext) PauseExecution() (sdkcore.JSON, error) {
	return nil, nil
}

// GetPauseTime returns the paused time in the RunContext.
// It retrieves the value from the pausedTime field in the RunContext struct.
func (r *ExecuteContext) GetPauseTime() *time.Time {
	return nil
}

// InputToType returns a pointer to a value of type T by marshaling and unmarshaling the ResolvedInput field of the provided RunContext struct.
// If there is an error during the marshaling or unmarshaling process, nil is returned.
// The function signature is as follows:
func InputToType[T any](ctx BaseContext) *T {
	b, err := json.Marshal(ctx.Input())
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

// InputToTypeSafely returns a pointer to a value of type T by marshaling and unmarshaling the ResolvedInput field of the provided RunContext struct.
// If there is an error during the marshaling or unmarshaling process, nil is returned.
// The function signature is as follows:
func InputToTypeSafely[T any](ctx BaseContext) (*T, error) {
	b, err := json.Marshal(ctx.Input())
	if err != nil {
		return nil, err
	}

	var rsp T
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}

// InputPropsToType returns a pointer to a value of type T by marshaling and unmarshaling the ResolvedInput field of the provided RunContext struct.
// If there is an error during the marshaling or unmarshaling process, nil is returned.
// The function signature is as follows:
func InputPropsToType[T any](input sdkcore.JSON) (*T, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var rsp T
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}

// DynamicInputToType converts the resolved input of type `sdkcore.DynamicOptionsContext` to the desired type T.
// It uses JSON marshaling and unmarshalling to perform the conversion.
// If any error occurs during marshaling or unmarshaling, it returns nil.
// The function returns a pointer to the converted value of type T.
func DynamicInputToType[T any](ctx *sdkcore.DynamicFieldContext) *T {
	b, err := json.Marshal(ctx.Input)
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
