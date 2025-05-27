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
	"log"
	"sync"
	"time"
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
	AddLog(ctx context.Context, level LogLevel, message string)
	SetPrefix(prefix string)
	GetLogs() []LogEntry
	ClearLogs()
	LogInfo(ctx context.Context, message string)
	LogWarning(ctx context.Context, message string)
	LogError(ctx context.Context, err error)
	LogDebug(ctx context.Context, message string)
}

// InternalLogger provides a centralized interface for managing logs.
type InternalLogger struct {
	mu     sync.RWMutex // Guards access to the logs slice
	logs   []LogEntry   // In-memory storage of logs
	sink   LogSink      // Optional: Sink to which logs can be written (e.g., database, file)
	level  LogLevel     // Minimum severity level to log (default: INFO)
	prefix string       // Prefix to enrich log messages
}

// LogSink defines an interface for persisting logs to external systems.
type LogSink interface {
	Write(ctx context.Context, logEntry LogEntry) error // Write a log entry to the sink
}

// NewLogger initializes a new Logger instance.
func NewLogger(sink LogSink, level LogLevel, prefix string) Logger {
	return &InternalLogger{
		logs:   []LogEntry{},
		sink:   sink,
		level:  level,
		prefix: prefix,
	}
}

// AddLog adds a log entry to the internal log storage and writes it to the sink (if present).
func (l *InternalLogger) AddLog(ctx context.Context, level LogLevel, message string) {
	// Check log level
	if !l.shouldLog(level) {
		return
	}

	// Add prefix to the log message
	if l.prefix != "" {
		if l.prefix[0] == '[' && l.prefix[len(l.prefix)-1] == ']' {
			message = l.prefix + " " + message
		} else {
			message = "[" + l.prefix + "] " + message
		}
	}

	// Create log entry
	logEntry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
	}

	// Store in memory for retrieval
	l.mu.Lock()
	l.logs = append(l.logs, logEntry)
	l.mu.Unlock()

	// Print to standard log output (optional)
	log.Printf("[%s] %s\n", logEntry.Level, logEntry.Message)

	// Persist to sink, if available
	if l.sink != nil {
		err := l.sink.Write(ctx, logEntry)
		if err != nil {
			log.Printf("Failed to write log to sink: %v\n", err)
		}
	}
}

// SetPrefix sets or updates the prefix for all subsequent logs.
func (l *InternalLogger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

// GetLogs retrieves all stored logs.
func (l *InternalLogger) GetLogs() []LogEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return append([]LogEntry{}, l.logs...) // Return a copy to prevent modifications
}

// ClearLogs clears the in-memory logs.
func (l *InternalLogger) ClearLogs() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logs = []LogEntry{}
}

// shouldLog determines if a log should be recorded based on the current log level.
func (l *InternalLogger) shouldLog(level LogLevel) bool {
	levelOrder := map[LogLevel]int{
		LevelDebug:   0,
		LevelInfo:    1,
		LevelWarning: 2,
		LevelError:   3,
	}
	return levelOrder[level] >= levelOrder[l.level]
}

// LogInfo Convenience wrappers for specific log levels
func (l *InternalLogger) LogInfo(ctx context.Context, message string) {
	l.AddLog(ctx, LevelInfo, message)
}

func (l *InternalLogger) LogWarning(ctx context.Context, message string) {
	l.AddLog(ctx, LevelWarning, message)
}

func (l *InternalLogger) LogError(ctx context.Context, err error) {
	if err != nil {
		l.AddLog(ctx, LevelError, err.Error())
	}
}

func (l *InternalLogger) LogDebug(ctx context.Context, message string) {
	l.AddLog(ctx, LevelDebug, message)
}
