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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

// LogLineLevel represents the kind of job to be started.
type LogLineLevel string

// 'DEBUG', 'INFO', 'WARN', 'ERROR'

// Available job kinds.
const (
	LogLineLevelDebug LogLineLevel = "DEBUG"
	LogLineLevelInfo  LogLineLevel = "INFO"
	LogLineLevelWarn  LogLineLevel = "WARN"
	LogLineLevelError LogLineLevel = "ERROR"
)

func (LogLineLevel) SQLTypeName() string {
	return "log_line_level"
}

// Values returns a slice of all String values of the enum.
func (LogLineLevel) Values() []string {
	return []string{
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j LogLineLevel) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j LogLineLevel) Validate() error {
	if !_j.IsValid() {
		return errors.New(fmt.Sprintf("LogLineLevel(%v) is %v", _j, ErrNoValidEnum))
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern LogLineLevel(%d) instead.
func (_j LogLineLevel) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("LogLineLevel(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for LogLineLevel.
func (_j LogLineLevel) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as LogLineLevel. %v", _j, err))
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for LogLineLevel.
func (_j *LogLineLevel) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("LogLineLevel cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogLineLevelFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a LogLineLevel", str))
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for LogLineLevel.
func (_j LogLineLevel) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for LogLineLevel.
func (_j *LogLineLevel) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return errors.New(fmt.Sprintf("invalid value of LogLineLevel: %[1]T(%[1]v)", value))
	}
	if len(str) == 0 {
		return errors.New("LogLineLevel cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogLineLevelFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a LogLineLevel", str))
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for LogLineLevel.
func (_j LogLineLevel) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as LogLineLevel. %v", _j, err))
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for LogLineLevel.
func (_j *LogLineLevel) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return errors.New(fmt.Sprintf("LogLineLevel should be a string, got %q", data))
	}
	if len(str) == 0 {
		return errors.New("LogLineLevel cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogLineLevelFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a LogLineLevel", str))
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for LogLineLevel.
func (_j *LogLineLevel) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return errors.New(fmt.Sprintf("invalid value of LogLineLevel: %[1]T(%[1]v)", value))
	}
	if len(str) == 0 {
		return errors.New("LogLineLevel cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogLineLevelFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a LogLineLevel", str))
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for LogLineLevel.
func (_j LogLineLevel) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as LogLineLevel. %v", _j, err))
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for LogLineLevel.
func (_j *LogLineLevel) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("LogLineLevel cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogLineLevelFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a LogLineLevel", str))
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for LogLineLevel.
func (_j LogLineLevel) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot marshal value %q as LogLineLevel. %v", _j, err))
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for LogLineLevel.
func (_j *LogLineLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("LogLineLevel cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogLineLevelFromString(str)
	if !ok {
		return errors.New(fmt.Sprintf("value %q does not represent a LogLineLevel", str))
	}
	return nil
}

// LogLineLevelFromString determines the enum value with an exact case match.
func LogLineLevelFromString(raw string) (LogLineLevel, bool) {
	v, ok := _LogLineLevelStringToValueMap[raw]
	if !ok {
		return LogLineLevelDebug, false
	}
	return v, true
}

// LogLineLevelFromStringIgnoreCase determines the enum value with a case-insensitive match.
func LogLineLevelFromStringIgnoreCase(raw string) (LogLineLevel, bool) {
	v, ok := LogLineLevelFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _LogLineLevelLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_LogLineLevelStringToValueMap = map[string]LogLineLevel{
		"DEBUG": LogLineLevelDebug,
		"INFO":  LogLineLevelInfo,
		"WARN":  LogLineLevelWarn,
		"ERROR": LogLineLevelError,
	}
	_LogLineLevelLowerStringToValueMap = map[string]LogLineLevel{
		"DEBUG": LogLineLevelDebug,
		"INFO":  LogLineLevelInfo,
		"WARN":  LogLineLevelWarn,
		"ERROR": LogLineLevelError,
	}
)
