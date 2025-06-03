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

// TriggerType defines how a workflow can be triggered
type TriggerType string

const (
	// TriggerTypeScheduled indicates a workflow triggered by a schedule
	TriggerTypeScheduled TriggerType = "SCHEDULED"

	// TriggerTypeEvent indicates a workflow triggered by an event
	TriggerTypeEvent TriggerType = "EVENT"

	TriggerTypePubsub TriggerType = "PUBSUB"

	// TriggerTypePolling indicates a workflow triggered by polling for changes
	TriggerTypePolling TriggerType = "POLLING"

	// TriggerTypeWebhook indicates a workflow triggered by a webhook
	TriggerTypeWebhook TriggerType = "WEBHOOK"

	// TriggerTypeManual indicates a workflow triggered manually by a user
	TriggerTypeManual TriggerType = "MANUAL"

	// TriggerTypeAPI indicates a workflow triggered via the API
	TriggerTypeAPI TriggerType = "API"

	// TriggerTypeWorkflow indicates a workflow triggered by another workflow
	TriggerTypeWorkflow TriggerType = "WORKFLOW"

	// TriggerTypeMessage indicates a workflow triggered by a message
	TriggerTypeMessage TriggerType = "MESSAGE"

	// TriggerTypeButton indicates a workflow triggered by a button click
	TriggerTypeButton TriggerType = "BUTTON"
)

// SQLTypeName returns the SQL type name for serialization
func (TriggerType) SQLTypeName() string {
	return "trigger_type"
}

// Values returns a slice of all String values of the enum
func (TriggerType) Values() []string {
	return []string{
		"SCHEDULED",
		"EVENT",
		"PUBSUB",
		"POLLING",
		"WEBHOOK",
		"MANUAL",
		"API",
		"WORKFLOW",
		"MESSAGE",
		"BUTTON",
	}
}

// IsValid tests whether the value is a valid enum value
func (_j TriggerType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values
func (_j TriggerType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("TriggerType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value
func (_j TriggerType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("TriggerType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements encoding.BinaryMarshaler
func (_j TriggerType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (_j *TriggerType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler
func (_j TriggerType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler
func (_j *TriggerType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of TriggerType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// MarshalJSON implements json.Marshaler
func (_j TriggerType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (_j *TriggerType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("TriggerType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// Scan implements sql/driver.Scanner
func (_j *TriggerType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of TriggerType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// MarshalText implements encoding.TextMarshaler
func (_j TriggerType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (_j *TriggerType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler
func (_j TriggerType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TriggerType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler
func (_j *TriggerType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("TriggerType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TriggerTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TriggerType", str)
	}
	return nil
}

// TriggerTypeFromString determines the enum value with an exact case match
func TriggerTypeFromString(raw string) (TriggerType, bool) {
	v, ok := _TriggerTypeStringToValueMap[raw]
	if !ok {
		return TriggerTypeScheduled, false
	}
	return v, true
}

// TriggerTypeFromStringIgnoreCase determines the enum value with a case-insensitive match
func TriggerTypeFromStringIgnoreCase(raw string) (TriggerType, bool) {
	v, ok := TriggerTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _TriggerTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

// Maps for looking up enum values from strings
var (
	_TriggerTypeStringToValueMap = map[string]TriggerType{
		"SCHEDULED": TriggerTypeScheduled,
		"EVENT":     TriggerTypeEvent,
		"PUBSUB":    TriggerTypePubsub,
		"POLLING":   TriggerTypePolling,
		"WEBHOOK":   TriggerTypeWebhook,
		"MANUAL":    TriggerTypeManual,
		"API":       TriggerTypeAPI,
		"WORKFLOW":  TriggerTypeWorkflow,
		"MESSAGE":   TriggerTypeMessage,
		"BUTTON":    TriggerTypeButton,
	}
	_TriggerTypeLowerStringToValueMap = map[string]TriggerType{
		"scheduled": TriggerTypeScheduled,
		"event":     TriggerTypeEvent,
		"pubsub":    TriggerTypePubsub,
		"polling":   TriggerTypePolling,
		"webhook":   TriggerTypeWebhook,
		"manual":    TriggerTypeManual,
		"api":       TriggerTypeAPI,
		"workflow":  TriggerTypeWorkflow,
		"message":   TriggerTypeMessage,
		"button":    TriggerTypeButton,
	}
)

// GetTriggerIcon returns an icon name for a trigger type
func GetTriggerIcon(triggerType TriggerType) string {
	switch triggerType {
	case TriggerTypeScheduled:
		return "calendar"
	case TriggerTypeEvent:
		return "zap"
	case TriggerTypePolling:
		return "refresh-cw"
	case TriggerTypeWebhook:
		return "link"
	case TriggerTypeManual:
		return "user"
	case TriggerTypeAPI:
		return "code"
	case TriggerTypeWorkflow:
		return "git-merge"
	case TriggerTypeMessage:
		return "message-circle"
	case TriggerTypeButton:
		return "mouse-pointer"
	default:
		return "play-circle"
	}
}

// GetTriggerDisplayName returns a human-readable name for a trigger type
func GetTriggerDisplayName(triggerType TriggerType) string {
	switch triggerType {
	case TriggerTypeScheduled:
		return "Scheduled"
	case TriggerTypeEvent:
		return "Event"
	case TriggerTypePolling:
		return "Polling"
	case TriggerTypeWebhook:
		return "Webhook"
	case TriggerTypeManual:
		return "Manual"
	case TriggerTypeAPI:
		return "API"
	case TriggerTypeWorkflow:
		return "Workflow"
	case TriggerTypeMessage:
		return "Message"
	case TriggerTypeButton:
		return "Button"
	default:
		return "Trigger"
	}
}

// RequiresConfiguration returns true if the trigger type needs additional configuration
func (t TriggerType) RequiresConfiguration() bool {
	switch t {
	case TriggerTypeScheduled, TriggerTypePolling, TriggerTypeWebhook, TriggerTypeEvent:
		return true
	default:
		return false
	}
}

// SupportsPolling returns true if the trigger type supports polling
func (t TriggerType) SupportsPolling() bool {
	return t == TriggerTypePolling
}

// SupportsScheduling returns true if the trigger type supports scheduling
func (t TriggerType) SupportsScheduling() bool {
	return t == TriggerTypeScheduled
}

// SupportsWebhook returns true if the trigger type supports webhooks
func (t TriggerType) SupportsWebhook() bool {
	return t == TriggerTypeWebhook
}

// SupportsManualTrigger returns true if the trigger type supports manual triggering
func (t TriggerType) SupportsManualTrigger() bool {
	return t == TriggerTypeManual || t == TriggerTypeButton
}
