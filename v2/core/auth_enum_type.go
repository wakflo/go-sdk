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

// AuthType defines the types of authentication supported by the system
type AuthType string

const (
	// None indicates no authentication is required
	None AuthType = "none"

	// Basic indicates username and password authentication
	Basic AuthType = "basic"

	// Secret indicates a shared secret or token authentication
	Secret AuthType = "secret"

	// APIKey indicates API key authentication
	APIKey AuthType = "api_key"

	// OAuth2 indicates OAuth 2.0 authentication
	OAuth2 AuthType = "oauth2"

	// Custom indicates a custom authentication method
	Custom AuthType = "custom"

	// JWT indicates JSON Web Token authentication
	JWT AuthType = "jwt"

	// ApiKeyHeader indicates an API key in the header
	ApiKeyHeader AuthType = "api_key_header"

	// ApiKeyQuery indicates an API key in the query string
	ApiKeyQuery AuthType = "api_key_query"

	// BearerToken indicates a bearer token authentication
	BearerToken AuthType = "bearer_token"

	// ClientCert indicates client certificate authentication
	ClientCert AuthType = "client_cert"
)

// SQLTypeName returns the SQL type name for serialization
func (AuthType) SQLTypeName() string {
	return "auth_type"
}

// Values returns a slice of all string values for the enum
func (AuthType) Values() []string {
	return []string{
		"basic",
		"secret",
		"api_key",
		"oauth2",
		"none",
		"custom",
		"jwt",
		"api_key_header",
		"api_key_query",
		"bearer_token",
		"client_cert",
	}
}

// IsValid tests whether the value is a valid enum value
func (_j AuthType) IsValid() bool {
	return slices.Contains(_j.Values(), string(_j))
}

// Validate whether the value is within the range of enum values
func (_j AuthType) Validate() error {
	if !_j.IsValid() {
		return fmt.Errorf("AuthType(%v) is %w", _j, ErrNoValidEnum)
	}
	return nil
}

// String returns the string representation of the enum value
func (_j AuthType) String() string {
	if !_j.IsValid() {
		return fmt.Sprintf("AuthType(%v)", string(_j))
	}
	return string(_j)
}

// MarshalBinary implements encoding.BinaryMarshaler
func (_j AuthType) MarshalBinary() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (_j *AuthType) UnmarshalBinary(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler
func (_j AuthType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(_j.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler
func (_j *AuthType) UnmarshalGQL(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AuthType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// MarshalJSON implements json.Marshaler
func (_j AuthType) MarshalJSON() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthType. %w", _j, err)
	}
	return json.Marshal(_j.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (_j *AuthType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("AuthType should be a string, got %q", data)
	}
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// Scan implements sql/driver.Scanner
func (_j *AuthType) Scan(value interface{}) error {
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AuthType: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// MarshalText implements encoding.TextMarshaler
func (_j AuthType) MarshalText() ([]byte, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthType. %w", _j, err)
	}
	return []byte(_j.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (_j *AuthType) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// MarshalYAML implements a YAML Marshaler
func (_j AuthType) MarshalYAML() (interface{}, error) {
	if err := _j.Validate(); err != nil {
		return nil, fmt.Errorf("cannot marshal value %q as AuthType. %w", _j, err)
	}
	return _j.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler
func (_j *AuthType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if len(str) == 0 {
		return errors.New("AuthType cannot be derived from empty string")
	}

	var ok bool
	*_j, ok = AuthTypeFromString(str)
	if !ok {
		return fmt.Errorf("value %q does not represent a AuthType", str)
	}
	return nil
}

// AuthTypeFromString determines the enum value with an exact case match
func AuthTypeFromString(raw string) (AuthType, bool) {
	v, ok := _AuthTypeStringToValueMap[raw]
	if !ok {
		return None, false
	}
	return v, true
}

// AuthTypeFromStringIgnoreCase determines the enum value with a case-insensitive match
func AuthTypeFromStringIgnoreCase(raw string) (AuthType, bool) {
	v, ok := AuthTypeFromString(raw)
	if ok {
		return v, ok
	}
	v, ok = _AuthTypeLowerStringToValueMap[raw]
	if !ok {
		return "", false
	}
	return v, true
}

// Maps for looking up enum values from strings
var (
	_AuthTypeStringToValueMap = map[string]AuthType{
		"basic":          Basic,
		"secret":         Secret,
		"api_key":        APIKey,
		"oauth2":         OAuth2,
		"none":           None,
		"custom":         Custom,
		"jwt":            JWT,
		"api_key_header": ApiKeyHeader,
		"api_key_query":  ApiKeyQuery,
		"bearer_token":   BearerToken,
		"client_cert":    ClientCert,
	}
	_AuthTypeLowerStringToValueMap = map[string]AuthType{
		"basic":          Basic,
		"secret":         Secret,
		"api_key":        APIKey,
		"oauth2":         OAuth2,
		"none":           None,
		"custom":         Custom,
		"jwt":            JWT,
		"api_key_header": ApiKeyHeader,
		"api_key_query":  ApiKeyQuery,
		"bearer_token":   BearerToken,
		"client_cert":    ClientCert,
	}
)

// GetAuthIcon returns an appropriate icon for an auth type
func GetAuthIcon(authType AuthType) string {
	switch authType {
	case Basic:
		return "key"
	case Secret:
		return "lock"
	case APIKey, ApiKeyHeader, ApiKeyQuery:
		return "key"
	case OAuth2:
		return "shield"
	case JWT, BearerToken:
		return "badge"
	case ClientCert:
		return "file-certificate"
	case Custom:
		return "settings"
	default:
		return "circle"
	}
}

// GetAuthDisplayName returns a human-readable name for an auth type
func GetAuthDisplayName(authType AuthType) string {
	switch authType {
	case Basic:
		return "Basic Authentication"
	case Secret:
		return "Secret Key"
	case APIKey:
		return "API Key"
	case OAuth2:
		return "OAuth 2.0"
	case JWT:
		return "JWT Token"
	case ApiKeyHeader:
		return "API Key (Header)"
	case ApiKeyQuery:
		return "API Key (Query)"
	case BearerToken:
		return "Bearer Token"
	case ClientCert:
		return "Client Certificate"
	case Custom:
		return "Custom Authentication"
	case None:
		return "No Authentication"
	default:
		return "Authentication"
	}
}
