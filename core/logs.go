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
	"time"

	"github.com/google/uuid"
)

type LogLevel string

const (
	Debug   LogLevel = "debug"
	Error   LogLevel = "error"
	Warning LogLevel = "warning"
	Info    LogLevel = "info"
)

type SystemActivityLog struct {
	Level     LogLevel  `json:"level"`
	Scope     string    `json:"scope"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type SystemActivityLogs = []SystemActivityLog

type WriteLogLineOpts struct {
	// The step run id
	TeamID string `json:"team_id" validate:"required,uuid"`

	// The step run id
	StepRunID *string `json:"step_run_id" validate:"uuid"`

	// The workflow run id
	WorkflowID string `json:"workflow_id" validate:"uuid"`

	// (optional) The time when the log line was created.
	CreatedAt *time.Time

	// (required) The message of the log line.
	Message string `json:"message" validate:"required,min=1,max=10000"`

	// (optional) The level of the log line.
	Level *string `json:"level" validate:"omitnil,oneof=INFO ERROR WARN DEBUG"`

	// (optional) The metadata of the log line.
	Metadata map[string]any `json:"metadata"`
}

// LogLine is the model entity for the LogLine schema.
type LogLine struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID uuid.UUID `json:"team_id,omitempty"`
	// StepRunID holds the value of the "step_run_id" field.
	StepRunID *string `json:"step_run_id,omitempty"`
	// WorkflowID holds the value of the "step_run_id" field.
	WorkflowID *string `json:"workflow_id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Level holds the value of the "level" field.
	Level LogLineLevel `json:"level,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Log struct {
	ops *WriteLogLineOpts
}

func NewLog(teamID string, workflowID string, stepRunID *string) *Log {
	return &Log{ops: &WriteLogLineOpts{
		StepRunID:  stepRunID,
		WorkflowID: workflowID,
		TeamID:     teamID,
		Message:    "",
		Metadata:   nil,
	}}
}

func (b *Log) Error() *LogBuilder {
	lvl := LogLineLevelError.String()
	b.ops.Level = &lvl
	return NewLogBuilder(b.ops)
}

func (b *Log) Info() *LogBuilder {
	lvl := LogLineLevelInfo.String()
	b.ops.Level = &lvl
	return NewLogBuilder(b.ops)
}

func (b *Log) Warn() *LogBuilder {
	lvl := LogLineLevelWarn.String()
	b.ops.Level = &lvl
	return NewLogBuilder(b.ops)
}

func (b *Log) Debug() *LogBuilder {
	lvl := LogLineLevelDebug.String()
	b.ops.Level = &lvl
	return NewLogBuilder(b.ops)
}

type LogBuilder struct {
	ops *WriteLogLineOpts
}

func NewLogBuilder(
	ops *WriteLogLineOpts,
) *LogBuilder {
	return &LogBuilder{ops: ops}
}

func (b *LogBuilder) Meta(meta map[string]interface{}) *LogBuilder {
	b.ops.Metadata = meta
	return b
}

func (b *LogBuilder) Msg(message string) {
	b.ops.Message = message
	t := time.Now()
	b.ops.CreatedAt = &t
}
