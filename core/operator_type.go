package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
)

type LogicalOperator string

const (
	LogicalOperatorEqual        LogicalOperator = "==="
	LogicalOperatorNotEqual     LogicalOperator = "!=="
	LogicalOperatorGreaterThan  LogicalOperator = ">"
	LogicalOperatorLessThan     LogicalOperator = "<"
	LogicalOperatorGreaterEqual LogicalOperator = ">="
	LogicalOperatorLessEqual    LogicalOperator = "<="
	LogicalOperatorAnd          LogicalOperator = "&&"
	LogicalOperatorOr           LogicalOperator = "||"
	LogicalOperatorNot          LogicalOperator = "!"

	LogicalOperatorDateBefore       LogicalOperator = "date_before"
	LogicalOperatorDateAfter        LogicalOperator = "date_after"
	LogicalOperatorDateEquals       LogicalOperator = "date_equals"
	LogicalOperatorBooleanIsTrue    LogicalOperator = "is_true"
	LogicalOperatorBooleanIsFalse   LogicalOperator = "is_false"
	LogicalOperatorStringStartsWith LogicalOperator = "starts_with"
	LogicalOperatorStringEndsWith   LogicalOperator = "ends_with"
	LogicalOperatorStringContains   LogicalOperator = "contains"
	LogicalOperatorArrayContains    LogicalOperator = "array_contains"
	LogicalOperatorArrayNotContains LogicalOperator = "array_not_contains"
)

func (LogicalOperator) SQLTypeName() string {
	return "logical_operator_type"
}

// Values returns a slice of all String values of the enum.
func (LogicalOperator) Values() []string {
	return []string{
		"===",
		"!==",
		">",
		"<",
		">=",
		"<=",
		"&&",
		"||",
		"!",
		"date_before",
		"date_after",
		"date_equals",
		"is_true",
		"is_false",
		"starts_with",
		"ends_with",
		"contains",
		"array_contains",
		"array_not_contains",
	}
}

// IsValid tests whether the value is a valid enum value.
func (_j LogicalOperator) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values.
func (_j LogicalOperator) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("LogicalOperator(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string of the enum value.
// If the enum value is invalid, it will produce a string
// of the following pattern LogicalOperator(%d) instead.
func (_j LogicalOperator) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("LogicalOperator(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for LogicalOperator.
func (_j LogicalOperator) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as LogicalOperator. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for LogicalOperator.
func (_j *LogicalOperator) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("LogicalOperator cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogicalOperatorFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a LogicalOperator", str)
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for LogicalOperator.
func (_j LogicalOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for LogicalOperator.
func (_j *LogicalOperator) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of LogicalOperator: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("LogicalOperator cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogicalOperatorFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a LogicalOperator", str)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for LogicalOperator.
func (_j LogicalOperator) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as LogicalOperator. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for LogicalOperator.
func (_j *LogicalOperator) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("LogicalOperator should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("LogicalOperator cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogicalOperatorFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a LogicalOperator", str)
	}
	return nil
}

// Scan implements the sql/driver.Scanner interface for LogicalOperator.
func (_j *LogicalOperator) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of LogicalOperator: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("LogicalOperator cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogicalOperatorFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a LogicalOperator", str)
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for LogicalOperator.
func (_j LogicalOperator) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as LogicalOperator. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for LogicalOperator.
func (_j *LogicalOperator) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("LogicalOperator cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogicalOperatorFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a LogicalOperator", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler for LogicalOperator.
func (_j LogicalOperator) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as LogicalOperator. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for LogicalOperator.
func (_j *LogicalOperator) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("LogicalOperator cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = LogicalOperatorFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a LogicalOperator", str)
	}
	return nil
}

// LogicalOperatorFromString determines the enum value with an exact case match.
func LogicalOperatorFromString(raw string) (LogicalOperator, bool) {
	v, ok := _LogicalOperatorStringToValueMap[raw]
	if !ok {
		return LogicalOperatorEqual, false
	}
	return v, true
}

// LogicalOperatorFromStringIgnoreCase determines the enum value with a case-insensitive match.
func LogicalOperatorFromStringIgnoreCase(raw string) (LogicalOperator, bool) {
	v, ok := LogicalOperatorFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _LogicalOperatorLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

var (
	_LogicalOperatorStringToValueMap = map[string]LogicalOperator{
		"===":                LogicalOperatorEqual,
		"!==":                LogicalOperatorNotEqual,
		">":                  LogicalOperatorGreaterThan,
		"<":                  LogicalOperatorLessThan,
		">=":                 LogicalOperatorGreaterEqual,
		"<=":                 LogicalOperatorLessEqual,
		"&&":                 LogicalOperatorAnd,
		"||":                 LogicalOperatorOr,
		"!":                  LogicalOperatorNot,
		"date_before":        LogicalOperatorDateBefore,
		"date_after":         LogicalOperatorDateAfter,
		"date_equals":        LogicalOperatorDateEquals,
		"is_true":            LogicalOperatorBooleanIsTrue,
		"is_false":           LogicalOperatorBooleanIsFalse,
		"starts_with":        LogicalOperatorStringStartsWith,
		"ends_with":          LogicalOperatorStringEndsWith,
		"contains":           LogicalOperatorStringContains,
		"array_contains":     LogicalOperatorArrayContains,
		"array_not_contains": LogicalOperatorArrayNotContains,
	}
	_LogicalOperatorLowerStringToValueMap = map[string]LogicalOperator{
		"===":                LogicalOperatorEqual,
		"!==":                LogicalOperatorNotEqual,
		">":                  LogicalOperatorGreaterThan,
		"<":                  LogicalOperatorLessThan,
		">=":                 LogicalOperatorGreaterEqual,
		"<=":                 LogicalOperatorLessEqual,
		"&&":                 LogicalOperatorAnd,
		"||":                 LogicalOperatorOr,
		"!":                  LogicalOperatorNot,
		"date_before":        LogicalOperatorDateBefore,
		"date_after":         LogicalOperatorDateAfter,
		"date_equals":        LogicalOperatorDateEquals,
		"is_true":            LogicalOperatorBooleanIsTrue,
		"is_false":           LogicalOperatorBooleanIsFalse,
		"starts_with":        LogicalOperatorStringStartsWith,
		"ends_with":          LogicalOperatorStringEndsWith,
		"contains":           LogicalOperatorStringContains,
		"array_contains":     LogicalOperatorArrayContains,
		"array_not_contains": LogicalOperatorArrayNotContains,
	}
)
