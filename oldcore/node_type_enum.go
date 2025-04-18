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

package oldcore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

type NodeType string

const (
	NodeTypeBranch      NodeType = "branch"
	NodeTypeBoolean     NodeType = "boolean"
	NodeTypeNormal      NodeType = "normal"
	NodeTypeLoop        NodeType = "loop"
	NodeTypeAction      NodeType = "action"
	NodeTypeConditon    NodeType = "condition"
	NodeTypePlaceholder NodeType = "placeholder"
)

func (NodeType) SQLTypeName() string {
	return "connector_type"
}

// Values returns a slice of all String values of the enum.
func (NodeType) Values() []string {
	return []string{
		"branch",
		"boolean",
		"normal",
		"loop",
		"action",
		"condition",
		"placeholder",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j NodeType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j NodeType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("NodeType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern NodeType(%d) instead.
func (_j NodeType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("NodeType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for NodeType.
func (_j NodeType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as NodeType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for NodeType.
func (_j *NodeType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("NodeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = NodeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a NodeType", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for NodeType.
func (_j NodeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for NodeType.
func (_j *NodeType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of NodeType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("NodeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = NodeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a NodeType", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for NodeType.
func (_j NodeType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as NodeType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for NodeType.
func (_j *NodeType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("NodeType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("NodeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = NodeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a NodeType", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for NodeType.
func (_j *NodeType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of NodeType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("NodeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = NodeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a NodeType", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for NodeType.
func (_j NodeType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as NodeType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for NodeType.
func (_j *NodeType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("NodeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = NodeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a NodeType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for NodeType.
func (_j NodeType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as NodeType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for NodeType.
func (_j *NodeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("NodeType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = NodeTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a NodeType", str)
	}
	return nil
}

// NodeTypeFromString determines the enum value with an exact case match.
func NodeTypeFromString(raw string) (NodeType, bool) {
	v, ok := _NodeTypeStringToValueMap[raw]
	if !ok {
		return NodeTypeNormal, false
	}
	return v, true
}

// NodeTypeFromStringIgnoreCase determines the enum value with a case-insensitive match.
func NodeTypeFromStringIgnoreCase(raw string) (NodeType, bool) {
	v, ok := NodeTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _NodeTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_NodeTypeStringToValueMap = map[string]NodeType{
		"branch":      NodeTypeBranch,
		"boolean":     NodeTypeBoolean,
		"normal":      NodeTypeNormal,
		"loop":        NodeTypeLoop,
		"action":      NodeTypeAction,
		"condition":   NodeTypeConditon,
		"placeholder": NodeTypePlaceholder,
	}
	_NodeTypeLowerStringToValueMap = map[string]NodeType{
		"branch":      NodeTypeBranch,
		"boolean":     NodeTypeBoolean,
		"normal":      NodeTypeNormal,
		"loop":        NodeTypeLoop,
		"action":      NodeTypeAction,
		"condition":   NodeTypeConditon,
		"placeholder": NodeTypePlaceholder,
	}
)
