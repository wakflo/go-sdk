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

type DynamicField struct {
	*BaseComponentField
	GetDynamicOptions *sdkcore.DynamicOptionsFn `json:"_,omitempty"`
}

func NewDynamicField(schemaType sdkcore.AutoFormType) *DynamicField {
	c := &DynamicField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(schemaType)
	c.builder.WithFieldType(sdkcore.DynamicType)
	c.builder.schema.IsDynamic = true

	return c
}

func (b *DynamicField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *DynamicField) SetDefaultValue(defaultValue interface{}) *DynamicField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *DynamicField) SetMinLength(len int) *DynamicField {
	b.builder.WithMinLength(&len)
	return b
}

func (b *DynamicField) SetMaxLength(len int) *DynamicField {
	b.builder.WithMaxLength(&len)
	return b
}

// rest
func (b *DynamicField) SetDescription(desc string) *DynamicField {
	b.builder.WithDescription(desc)
	return b
}

func (b *DynamicField) SetDisplayName(title string) *DynamicField {
	b.builder.WithTitle(title)
	return b
}

func (b *DynamicField) SetRequired(required bool) *DynamicField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *DynamicField) SetDisabled(disabled bool) *DynamicField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.Presentation.Disabled = disabled
	return b
}

func (b *DynamicField) SetAnyOf(schemas []*sdkcore.AutoFormSchema) *DynamicField {
	b.builder.WithAnyOf(schemas)
	return b
}

func (b *DynamicField) SetOneOf(schemas []*sdkcore.AutoFormSchema) *DynamicField {
	b.builder.WithOneOf(schemas)
	return b
}

func (b *DynamicField) SetAllOf(schemas []*sdkcore.AutoFormSchema) *DynamicField {
	b.builder.WithAllOf(schemas)
	return b
}

func (b *DynamicField) SetDynamicOptions(fn *sdkcore.DynamicOptionsFn) *DynamicField {
	b.builder.schema.SetDynamicOptionsFn(fn)
	b.GetDynamicOptions = fn
	return b
}

func (b *DynamicField) SetDependsOn(deps []string) *DynamicField {
	b.builder.schema.DependsOn = deps
	return b
}
