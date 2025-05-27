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
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

// LogLevel represents the severity of a log.
type LogLevel string

const (
	LevelInfo    LogLevel = "INFO"
	LevelWarning LogLevel = "WARN"
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
	AddLog(level LogLevel, message string, a ...any)
	SetPrefix(prefix string)
	GetLogs() []LogEntry
	ClearLogs()
	Info(message string)
	Infof(message string, a ...any)
	Warn(message string)
	Warnf(message string, a ...any)
	Error(err error)
	Errorf(err error, message string, a ...any)
	Debug(message string)
	Debugf(message string, a ...any)
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	Clone() Logger
}

// LogSink defines an interface for persisting logs to external systems.
type LogSink interface {
	Write(ctx context.Context, logEntry LogEntry) error // Write a log entry to the sink
}

type NoopLogger struct {
	logs   []LogEntry
	prefix string
}

func (n *NoopLogger) AddLog(level LogLevel, message string, a ...any) {
	log.Info().Msg(message)

	var formattedMessage string
	if len(a) > 0 {
		formattedMessage = fmt.Sprintf(message, a...)
	} else {
		formattedMessage = message
	}

	n.logs = append(n.logs, LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   formattedMessage,
	})
}

func (n *NoopLogger) AddLogMessage(level LogLevel, message string) {
	log.Info().Msg(message)

	var formattedMessage string
	formattedMessage = message

	n.logs = append(n.logs, LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   formattedMessage,
	})
}

func (n *NoopLogger) SetPrefix(prefix string) {
	n.prefix = prefix
}

func (n *NoopLogger) GetLogs() []LogEntry {
	return n.logs
}

func (n *NoopLogger) Clone() Logger {
	return &NoopLogger{
		logs:   n.logs,
		prefix: n.prefix,
	}
}

func (n *NoopLogger) ClearLogs() {
	n.logs = []LogEntry{}
}

func (n *NoopLogger) Info(message string) {
	n.AddLogMessage(LevelInfo, message)
}

func (n *NoopLogger) Infof(message string, a ...any) {
	n.AddLog(LevelInfo, message, a...)
}

func (n *NoopLogger) Warn(message string) {
	//nolint:govet,printf // Disable multiple linters
	n.AddLogMessage(LevelWarning, message)
}

func (n *NoopLogger) Warnf(message string, a ...any) {
	n.AddLog(LevelWarning, message, a...)
}

func (n *NoopLogger) Error(err error) {
	n.AddLogMessage(LevelError, err.Error())
}

func (n *NoopLogger) Errorf(err error, message string, a ...any) {
	// Format the message first, then pass as single argument
	formattedMessage := fmt.Sprintf(message, a...)
	n.AddLog(LevelError, "%s: %v", formattedMessage, err)
}

func (n *NoopLogger) Debug(message string) {
	n.AddLogMessage(LevelDebug, message)
}

func (n *NoopLogger) Debugf(message string, a ...any) {
	n.AddLog(LevelDebug, message, a...)
}

func (n *NoopLogger) WithField(key string, value interface{}) Logger {
	return n
}

func (n *NoopLogger) WithFields(fields map[string]interface{}) Logger {
	return n
}

func NewNoopLogger() Logger {
	return &NoopLogger{}
}
