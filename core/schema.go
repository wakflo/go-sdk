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
	String                 = "string"
	Number                 = "number"
	Object                 = "object"
	Array                  = "array"
	Boolean                = "boolean"
	Nullable               = "null"
	Integer                = "integer"
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
	Default              interface{}                `json:"default,omitempty"`
	Format               string                     `json:"format,omitempty"`
	Definitions          map[string]*AutoFormSchema `json:"definitions,omitempty"`
	Description          string                     `json:"description,omitempty"`
	AdditionalProperties interface{}                `json:"additionalProperties,omitempty"`
	MinProperties        *int                       `json:"minProperties,omitempty"`
	MaxProperties        *int                       `json:"maxProperties,omitempty"`
	PatternProperties    map[string]*AutoFormSchema `json:"patternProperties,omitempty"`
	Dependencies         map[string]interface{}     `json:"dependencies,omitempty"`
	Enum                 []interface{}              `json:"enum,omitempty"`
	AllOf                []*AutoFormSchema          `json:"allOf,omitempty"`
	AnyOf                []*AutoFormSchema          `json:"anyOf,omitempty"`
	OneOf                []*AutoFormSchema          `json:"oneOf,omitempty"`
	Not                  *AutoFormSchema            `json:"not,omitempty"`
	Minimum              interface{}                `json:"minimum,omitempty"`
	Maximum              interface{}                `json:"maximum,omitempty"`
	ExclusiveMinimum     interface{}                `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum     interface{}                `json:"exclusiveMaximum,omitempty"`
	MinLength            *int                       `json:"minLength,omitempty"`
	MaxLength            *int                       `json:"maxLength,omitempty"`
	Pattern              string                     `json:"pattern,omitempty"`
	MinItems             *int                       `json:"minItems,omitempty"`
	MaxItems             *int                       `json:"maxItems,omitempty"`
	UniqueItems          bool                       `json:"uniqueItems,omitempty"`
	Contains             *AutoFormSchema            `json:"contains,omitempty"`
	MinContains          *int                       `json:"minContains,omitempty"`
	MaxContains          *int                       `json:"maxContains,omitempty"`
	Const                interface{}                `json:"const,omitempty"`
	Disabled             bool                       `json:"disabled"`

	Presentation *AutoFormFieldPresentation      `json:"x-jsf-presentation,omitempty"`
	Validations  []string                        `json:"x-jsf-logic-validations,omitempty"`
	Logic        *AutoFormFieldLogic             `json:"x-jsf-logic,omitempty"`
	ErrorMessage *AutoFormFieldPresentationError `json:"x-jsf-errorMessage,omitempty"`
	Order        []string                        `json:"x-jsf-order,omitempty"`

	If   *AutoFormSchema `json:"if,omitempty"`
	Else *AutoFormSchema `json:"else,omitempty"`
	Then *AutoFormSchema `json:"then,omitempty"`

	Scope    []string `json:"scope,omitempty"`
	TokenUrl *string  `json:"tokenUrl,omitempty"`
	AuthUrl  *string  `json:"authUrl,omitempty"`

	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Secret   *string `json:"secret,omitempty"`

	DependsOn []string `json:"dependsOn,omitempty"`
	IsDynamic bool     `json:"isDynamic,omitempty"`

	dynamicOptionsFn *DynamicOptionsFn
}

func (af *AutoFormSchema) GetDynamicOptionsFn(ctx *DynamicOptionsContext) (interface{}, error) {
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

type AutoFormFieldPresentation struct {
	InputType   AutoFormFieldType             `json:"inputType,omitempty"`
	Required    bool                          `json:"required"`
	Disabled    bool                          `json:"disabled"`
	Options     []*AutoFormFieldSelectOptions `json:"options,omitempty"`
	MaskSecret  *int                          `json:"maskSecret,omitempty"`
	Statement   *AutoFormFieldPresentation    `json:"statement,omitempty"`
	Severity    *string                       `json:"severity,omitempty"`
	Placeholder *string                       `json:"placeholder,omitempty"`
	ReadOnly    bool                          `json:"readOnly"`
	MinDate     *time.Time                    `json:"minDate,omitempty"`
	MaxDate     *time.Time                    `json:"maxDate,omitempty"`

	Extras map[string]interface{} `json:"extras,omitempty"`
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

type DynamicOptionsLinker struct {
	dynamicOptionsFn DynamicOptionsFn
}

type DynamicOptionsFn = func(ctx *DynamicOptionsContext) (interface{}, error)
