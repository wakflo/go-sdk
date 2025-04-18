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
	sdkcore "github.com/wakflo/go-sdk/oldcore"
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
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeDynamic)
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
	b.builder = b.builder.WithFieldRequired(required)
	return b
}

func (b *DynamicField) SetDisabled(disabled bool) *DynamicField {
	b.builder = b.builder.WithFieldDisabled(disabled)
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

func (b *DynamicField) SetPlaceholder(placeholder string) *DynamicField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *DynamicField) SetLabel(label string) *DynamicField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *DynamicField) SetHint(hint string) *DynamicField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *DynamicField) SeMultiSelect(enable bool) *DynamicField {
	b.builder.schema.UIProps.Multiple = enable
	return b
}

func (b *DynamicField) SetHidden(hidden bool) *DynamicField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
