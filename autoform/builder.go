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

package autoform

import (
	sdkcore "github.com/wakflo/go-sdk/core"
)

type SchemaBuilder struct {
	schema *sdkcore.AutoFormSchema
}

func NewSchemaBuilder() *SchemaBuilder {
	s := &SchemaBuilder{schema: &sdkcore.AutoFormSchema{}}
	s.schema.UIProps = &sdkcore.AutoFormFieldProps{}

	return s
}

func (b *SchemaBuilder) Build() *sdkcore.AutoFormSchema {
	return b.schema
}

func (b *SchemaBuilder) WithTitle(title string) *SchemaBuilder {
	b.schema.Title = title
	return b
}

func (b *SchemaBuilder) WithType(schemaType sdkcore.AutoFormType) *SchemaBuilder {
	b.schema.Type = schemaType
	return b
}

func (b *SchemaBuilder) WithFieldType(fieldType sdkcore.AutoFormFieldType) *SchemaBuilder {
	b.schema.UIControl = fieldType
	b.schema.UIProps.ControlType = fieldType
	return b
}

func (b *SchemaBuilder) WithFieldRequired(required bool) *SchemaBuilder {
	b.schema.UIProps.Required = required
	b.schema.IsRequired = required
	return b
}

func (b *SchemaBuilder) WithFieldDisabled(disabled bool) *SchemaBuilder {
	b.schema.Disabled = disabled
	b.schema.UIProps.Disabled = disabled
	return b
}

func (b *SchemaBuilder) WithDescription(description string) *SchemaBuilder {
	b.schema.Description = description
	return b
}

func (b *SchemaBuilder) WithDefault(defaultValue interface{}) *SchemaBuilder {
	b.schema.Default = defaultValue
	return b
}

func (b *SchemaBuilder) WithProperties(properties map[string]*sdkcore.AutoFormSchema) *SchemaBuilder {
	b.schema.Properties = properties
	return b
}

func (b *SchemaBuilder) WithRequired(required []string) *SchemaBuilder {
	b.schema.Required = required
	return b
}

func (b *SchemaBuilder) WithOrder(order []string) *SchemaBuilder {
	b.schema.Order = order
	return b
}

func (b *SchemaBuilder) WithDependencies(dependencies map[string]interface{}) *SchemaBuilder {
	b.schema.Dependencies = dependencies
	return b
}

func (b *SchemaBuilder) WithEnum(enum []interface{}) *SchemaBuilder {
	b.schema.Enum = enum
	return b
}

func (b *SchemaBuilder) WithItems(items *sdkcore.AutoFormSchema) *SchemaBuilder {
	b.schema.Items = items
	return b
}

func (b *SchemaBuilder) WithAdditionalItems(additionalItems *sdkcore.AutoFormSchema) *SchemaBuilder {
	b.schema.AdditionalItems = additionalItems
	return b
}

func (b *SchemaBuilder) WithSchema(schema string) *SchemaBuilder {
	b.schema.Schema = schema
	return b
}

func (b *SchemaBuilder) WithID(id string) *SchemaBuilder {
	b.schema.ID = id
	return b
}

func (b *SchemaBuilder) WithPattern(id string) *SchemaBuilder {
	b.schema.Pattern = id
	return b
}

func (b *SchemaBuilder) WithFormat(format string) *SchemaBuilder {
	b.schema.Format = format
	return b
}

func (b *SchemaBuilder) WithMinimum(min interface{}) *SchemaBuilder {
	b.schema.Minimum = min
	return b
}

func (b *SchemaBuilder) WithMaximum(max interface{}) *SchemaBuilder {
	b.schema.Maximum = max
	return b
}

func (b *SchemaBuilder) WithExclusiveMinimum(min interface{}) *SchemaBuilder {
	b.schema.ExclusiveMinimum = min
	return b
}

func (b *SchemaBuilder) WithExclusiveMaximum(max interface{}) *SchemaBuilder {
	b.schema.ExclusiveMaximum = max
	return b
}

func (b *SchemaBuilder) WithMinLength(minLength *int) *SchemaBuilder {
	b.schema.MinLength = minLength
	return b
}

func (b *SchemaBuilder) WithMaxLength(maxLength *int) *SchemaBuilder {
	b.schema.MaxLength = maxLength
	return b
}

func (b *SchemaBuilder) WithMaxContains(maxContains *int) *SchemaBuilder {
	b.schema.MaxContains = maxContains
	return b
}

func (b *SchemaBuilder) WithMinContains(minContains *int) *SchemaBuilder {
	b.schema.MinContains = minContains
	return b
}

func (b *SchemaBuilder) WithOneOf(oneOf []*sdkcore.AutoFormSchema) *SchemaBuilder {
	b.schema.OneOf = oneOf
	return b
}

func (b *SchemaBuilder) WithAnyOf(anyOf []*sdkcore.AutoFormSchema) *SchemaBuilder {
	b.schema.AnyOf = anyOf
	return b
}

func (b *SchemaBuilder) WithAllOf(allOf []*sdkcore.AutoFormSchema) *SchemaBuilder {
	b.schema.AllOf = allOf
	return b
}

func (b *SchemaBuilder) WithUniqueItems(uniqueItems bool) *SchemaBuilder {
	b.schema.UniqueItems = uniqueItems
	return b
}

func (b *SchemaBuilder) WithNot(not *sdkcore.AutoFormSchema) *SchemaBuilder {
	b.schema.Not = not
	return b
}

func (b *SchemaBuilder) WithPatternProperties(patternProperties map[string]*sdkcore.AutoFormSchema) *SchemaBuilder {
	b.schema.PatternProperties = patternProperties
	return b
}

func (b *SchemaBuilder) WithConst(constant interface{}) *SchemaBuilder {
	b.schema.Const = constant
	return b
}
