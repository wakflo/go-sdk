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

// ActionType defines the type of action in a workflow.
type ActionType string

const (
	// ActionTypeAction represents a standard action.
	ActionTypeAction ActionType = "ACTION"

	// ActionTypeBranch represents a branch action that can conditionally execute different paths.
	ActionTypeBranch ActionType = "BRANCH"

	// ActionTypeBoolean represents a boolean-conditional action.
	ActionTypeBoolean ActionType = "BOOLEAN"

	// ActionTypeLoop represents a loop action.
	ActionTypeLoop ActionType = "LOOP"

	// ActionTypeRouter represents a router action that can direct flow based on conditions.
	ActionTypeRouter ActionType = "ROUTER"
)

func (ActionType) SQLTypeName() string {
	return "action_type"
}

// Values returns a slice of all String values of the enum.
func (ActionType) Values() []string {
	return []string{
		"ACTION",
		"BRANCH",
		"BOOLEAN",
		"LOOP",
		"ROUTER",
		"FLOW",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j ActionType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j ActionType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("ActionType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern ActionType(%d) instead.
func (_j ActionType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("ActionType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for ActionType.
func (_j ActionType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ActionType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for ActionType.
func (_j *ActionType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ActionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ActionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ActionType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for ActionType.
func (_j ActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for ActionType.
func (_j *ActionType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ActionType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ActionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ActionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ActionType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for ActionType.
func (_j ActionType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ActionType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ActionType.
func (_j *ActionType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("ActionType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("ActionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ActionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ActionType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for ActionType.
func (_j *ActionType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ActionType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("ActionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ActionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ActionType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for ActionType.
func (_j ActionType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ActionType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ActionType.
func (_j *ActionType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("ActionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ActionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ActionType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for ActionType.
func (_j ActionType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as ActionType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ActionType.
func (_j *ActionType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("ActionType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = ActionTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a ActionType", str)
	}
	return nil
}

// ActionTypeFromString determines the enum value with an exact case match.
func ActionTypeFromString(raw string) (ActionType, bool) {
	v, ok := _ActionTypeStringToValueMap[raw]
	if !ok {
		return ActionTypeAction, false
	}
	return v, true
}

// ActionTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func ActionTypeFromStringIgnoreCase(raw string) (ActionType, bool) {
	v, ok := ActionTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _ActionTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_ActionTypeStringToValueMap = map[string]ActionType{
		"ACTION":  ActionTypeAction,
		"BRANCH":  ActionTypeBranch,
		"BOOLEAN": ActionTypeBoolean,
		"LOOP":    ActionTypeLoop,
		"ROUTER":  ActionTypeRouter,
	}
	_ActionTypeLowerStringToValueMap = map[string]ActionType{
		"action":  ActionTypeAction,
		"branch":  ActionTypeBranch,
		"boolean": ActionTypeBoolean,
		"loop":    ActionTypeLoop,
		"router":  ActionTypeRouter,
	}
)

// // ActionType defines the types of actions in a workflow
// type ActionType string
//
// const (
// 	// ActionTypeAction represents a standard action
// 	ActionTypeAction ActionType = "ACTION"
//
// 	// ActionTypeBranch represents a branch/router action
// 	ActionTypeBranch ActionType = "BRANCH"
//
// 	// ActionTypeBoolean represents a boolean condition action
// 	ActionTypeBoolean ActionType = "BOOLEAN"
//
// 	// ActionTypeFlow represents a subflow action
// 	ActionTypeFlow ActionType = "FLOW"
//
// 	// ActionTypeStep represents a composite step action
// 	ActionTypeStep ActionType = "STEP"
//
// 	// ActionTypeTransform represents a data transformation action
// 	ActionTypeTransform ActionType = "TRANSFORM"
//
// 	// ActionTypeApproval represents an approval action
// 	ActionTypeApproval ActionType = "APPROVAL"
//
// 	// ActionTypeCustom represents a custom action
// 	ActionTypeCustom ActionType = "CUSTOM"
// )
//
// // IsValid checks if the action type is valid
// func (t ActionType) IsValid() bool {
// 	switch t {
// 	case ActionTypeAction, ActionTypeBranch, ActionTypeBoolean,
// 		ActionTypeFlow, ActionTypeStep, ActionTypeTransform,
// 		ActionTypeApproval, ActionTypeCustom:
// 		return true
// 	default:
// 		return false
// 	}
// }
//
// // String returns the string representation of the action type
// func (t ActionType) String() string {
// 	return string(t)
// }
