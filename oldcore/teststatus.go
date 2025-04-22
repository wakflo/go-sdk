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

type TestStatus string

const (
	Fail TestStatus = "fail"
	Pass TestStatus = "pass"
)

func (TestStatus) SQLTypeName() string {
	return "test_status"
}

// Values returns a slice of all String values of the enum.
func (TestStatus) Values() []string {
	return []string{
		"fail",
		"pass",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j TestStatus) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j TestStatus) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("TestStatus(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern TestStatus(%d) instead.
func (_j TestStatus) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("TestStatus(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for TestStatus.
func (_j TestStatus) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TestStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for TestStatus.
func (_j *TestStatus) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("TestStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TestStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TestStatus", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for TestStatus.
func (_j TestStatus) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for TestStatus.
func (_j *TestStatus) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of TestStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("TestStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TestStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TestStatus", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for TestStatus.
func (_j TestStatus) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TestStatus. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TestStatus.
func (_j *TestStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("TestStatus should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("TestStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TestStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TestStatus", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for TestStatus.
func (_j *TestStatus) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of TestStatus: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("TestStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TestStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TestStatus", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for TestStatus.
func (_j TestStatus) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TestStatus. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for TestStatus.
func (_j *TestStatus) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("TestStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TestStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TestStatus", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for TestStatus.
func (_j TestStatus) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as TestStatus. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for TestStatus.
func (_j *TestStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("TestStatus cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = TestStatusFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a TestStatus", str)
	}
	return nil
}

// TestStatusFromString determines the enum value with an exact case match.
func TestStatusFromString(raw string) (TestStatus, bool) {
	v, ok := _TestStatusStringToValueMap[raw]
	if !ok {
		return Fail, false
	}
	return v, true
}

// TestStatusFromStringIgnoreCase determines the enum value with a case-insensitive match.
func TestStatusFromStringIgnoreCase(raw string) (TestStatus, bool) {
	v, ok := TestStatusFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _TestStatusLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_TestStatusStringToValueMap = map[string]TestStatus{
		"fail": Fail,
		"pass": Pass,
	}
	_TestStatusLowerStringToValueMap = map[string]TestStatus{
		"fail": Fail,
		"pass": Pass,
	}
)
