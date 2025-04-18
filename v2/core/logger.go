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
	"context"
	"time"
)

// LogLevel represents the severity of a log.
type LogLevel string

const (
	LevelInfo    LogLevel = "INFO"
	LevelWarning LogLevel = "WARNING"
	LevelError   LogLevel = "ERROR"
	LevelDebug   LogLevel = "DEBUG" // Optional: for verbose debugging messages
)

// LogEntry represents a single log message with a timestamp and level.
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"` // Timestamp of the log
	Level     LogLevel  `json:"level"`     // Severity level of the log
	Message   string    `json:"message"`   // Log message itself
}

// Logger provides a centralized interface for managing logs.
type Logger interface {
	AddLog(ctx context.Context, level LogLevel, message string)
	SetPrefix(prefix string)
	GetLogs() []LogEntry
	ClearLogs()
	LogInfo(ctx context.Context, message string)
	LogWarning(ctx context.Context, message string)
	LogError(ctx context.Context, err error)
	LogDebug(ctx context.Context, message string)
}

// LogSink defines an interface for persisting logs to external systems.
type LogSink interface {
	Write(ctx context.Context, logEntry LogEntry) error // Write a log entry to the sink
}
