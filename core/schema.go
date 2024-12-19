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
	"time"
)

type AutoFormType string

const (
	Undefined AutoFormType = "undefined"
	String    AutoFormType = "string"
	Number    AutoFormType = "number"
	Object    AutoFormType = "object"
	Array     AutoFormType = "array"
	Boolean   AutoFormType = "boolean"
	Nullable  AutoFormType = "null"
	Integer   AutoFormType = "integer"
)

func (s AutoFormType) String() string {
	return string(s)
}

// UnmarshalJSON function to convert json into enum value
func (s *AutoFormType) UnmarshalJSON(data []byte) error {
	var schema string
	if err := json.Unmarshal(data, &schema); err != nil {
		return err
	}

	switch schema {
	case "string":
		*s = String
	case "number":
		*s = Number
	case "object":
		*s = Object
	case "array":
		*s = Array
	case "boolean":
		*s = Boolean
	case "null":
		*s = Nullable
	case "integer":
		*s = Integer
	default:
		*s = Undefined
	}

	return nil
}

type AutoFormSchema struct {
	Schema               string                     `json:"$schema,omitempty"`
	ID                   string                     `json:"$id,omitempty"`
	Comment              *string                    `json:"$comment,omitempty"`
	Title                string                     `json:"title,omitempty"`
	Type                 AutoFormType               `json:"type,omitempty"`
	Properties           map[string]*AutoFormSchema `json:"properties,omitempty"`
	Items                *AutoFormSchema            `json:"items,omitempty"`
	AdditionalItems      *AutoFormSchema            `json:"additionalItems,omitempty"`
	Required             []string                   `json:"required,omitempty"`
	IsRequired           bool                       `json:"isRequired,omitempty"`
	Default              any                        `json:"default,omitempty"`
	Format               string                     `json:"format,omitempty"`
	Definitions          map[string]*AutoFormSchema `json:"definitions,omitempty"`
	Description          string                     `json:"description,omitempty"`
	AdditionalProperties any                        `json:"additionalProperties,omitempty"`
	MinProperties        *int                       `json:"minProperties,omitempty"`
	MaxProperties        *int                       `json:"maxProperties,omitempty"`
	PatternProperties    map[string]*AutoFormSchema `json:"patternProperties,omitempty"`
	Dependencies         map[string]any             `json:"dependencies,omitempty"`
	Enum                 []any                      `json:"enum,omitempty"`
	AllOf                []*AutoFormSchema          `json:"allOf,omitempty"`
	AnyOf                []*AutoFormSchema          `json:"anyOf,omitempty"`
	OneOf                []*AutoFormSchema          `json:"oneOf,omitempty"`
	Not                  *AutoFormSchema            `json:"not,omitempty"`
	Minimum              any                        `json:"minimum,omitempty"`
	Maximum              any                        `json:"maximum,omitempty"`
	ExclusiveMinimum     any                        `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum     any                        `json:"exclusiveMaximum,omitempty"`
	MinLength            *int                       `json:"minLength,omitempty"`
	MaxLength            *int                       `json:"maxLength,omitempty"`
	Pattern              string                     `json:"pattern,omitempty"`
	MinItems             *int                       `json:"minItems,omitempty"`
	MaxItems             *int                       `json:"maxItems,omitempty"`
	UniqueItems          bool                       `json:"uniqueItems,omitempty"`
	Contains             *AutoFormSchema            `json:"contains,omitempty"`
	MinContains          *int                       `json:"minContains,omitempty"`
	MaxContains          *int                       `json:"maxContains,omitempty"`
	Const                any                        `json:"const,omitempty"`
	Disabled             bool                       `json:"disabled"`
	Order                []string                   `json:"order,omitempty"`

	If   *AutoFormSchema `json:"if,omitempty"`
	Else *AutoFormSchema `json:"else,omitempty"`
	Then *AutoFormSchema `json:"then,omitempty"`

	UIControl AutoFormFieldType `json:"ui:control,omitempty"`
	// UIData            *AutoFormDataProps  `json:"ui:data,omitempty"`
	UIProps           *AutoFormFieldProps `json:"ui:props,omitempty"`
	UIComponentRemove *AutoFormSchema     `json:"ui:component:remove,omitempty"`
	UIBefore          []AutoFormSchema    `json:"ui:before,omitempty"`

	DependsOn []string `json:"dependsOn,omitempty"`
	IsDynamic bool     `json:"isDynamic,omitempty"`

	dynamicOptionsFn *DynamicOptionsFn
}

func (af *AutoFormSchema) GetDynamicOptionsFn(ctx *DynamicFieldContext) (interface{}, error) {
	fn := *af.dynamicOptionsFn
	return fn(ctx)
}

func (af *AutoFormSchema) SetDynamicOptionsFn(dynamicOptionsFn *DynamicOptionsFn) {
	af.dynamicOptionsFn = dynamicOptionsFn
}

type AutoFormFieldPresentationError struct {
	Required *string `json:"required,omitempty"`
	Minimum  *string `json:"minimum,omitempty"`
	Maximum  *string `json:"maximum,omitempty"`
}

type AutoFormFieldProps struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Type        string            `json:"type,omitempty"`
	ControlType AutoFormFieldType `json:"controlType,omitempty"`
	Required    bool              `json:"required"`
	Language    string            `json:"language,omitempty"`
	Hint        string            `json:"hint,omitempty"`
	Label       string            `json:"label,omitempty"`
	Disabled    bool              `json:"disabled"`
	Hidden      bool              `json:"hidden"`
	Placeholder string            `json:"placeholder,omitempty"`
	ReadOnly    bool              `json:"readOnly"`
	Multiple    bool              `json:"multiple"`
	MinDate     *time.Time        `json:"minDate,omitempty"`
	MaxDate     *time.Time        `json:"maxDate,omitempty"`
	Min         *int              `json:"min,omitempty"`
	Max         *int              `json:"max,omitempty"`

	Auth *AuthSchemaProps `json:"auth,omitempty"`

	// special fields //

	// Notify is used to notify other fields after changes
	Notify []string `json:"notify,omitempty"`

	// Example Usage in JS
	// relevant: ({ formApi, scope }) => {
	//    return formApi.getValue(`${scope}.married`) == 'yes';
	// }
	// The value should be a stringified JS function
	Relevant     *string `json:"relevant,omitempty"`
	InitialValue any     `json:"initialValue,omitempty"`
	KeepState    bool    `json:"keepState,omitempty"`
}

type AutoFormFieldLogic struct {
	Validations map[string]AutoFormFieldLogicValidationDefinition `json:"validations,omitempty"`
}

type AutoFormFieldLogicValidationDefinition struct {
	ErrorMessage string                                              `json:"errorMessage,omitempty"`
	Rule         map[FieldLogicOperators]*AutoFormFieldLogicRuleData `json:"rule,omitempty"`
}

type AutoFormFieldLogicRuleData struct {
	Var         *string `json:"var,omitempty"`
	MissingSome *string `json:"missing_some,omitempty"`
	Missing     *string `json:"missing,omitempty"`
}

// AutoFormFieldSelectOptions represents an enumerated type with a value and label.
type AutoFormFieldSelectOptions struct {
	Value    string `json:"value"`
	Label    string `json:"label"`
	Metadata any    `json:"metadata,omitempty"`
}

type DynamicOptionsFn = func(ctx *DynamicFieldContext) (*DynamicOptionsResponse, error)

type AutoFormDataProps struct {
	Auth *AuthSchemaProps `json:"auth,omitempty"`
}

type AuthSchemaProps struct {
	Scope    []string `json:"scope,omitempty"`
	TokenURL *string  `json:"tokenUrl,omitempty"`
	AuthURL  *string  `json:"authUrl,omitempty"`

	Username       *string  `json:"username,omitempty"`
	Password       *string  `json:"password,omitempty"`
	Secret         *string  `json:"secret,omitempty"`
	ExcludedParams []string `json:"excludedParams,omitempty"`
}
