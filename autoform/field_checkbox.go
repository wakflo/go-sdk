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

type CheckboxField struct {
	*BaseComponentField
}

func NewCheckboxField() *CheckboxField {
	c := &CheckboxField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Boolean)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeCheckbox)

	return c
}

func (b *CheckboxField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *CheckboxField) SetDefaultValue(defaultValue bool) *CheckboxField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *CheckboxField) SetDescription(desc string) *CheckboxField {
	b.builder.WithDescription(desc)
	return b
}

func (b *CheckboxField) SetDisplayName(title string) *CheckboxField {
	b.builder.WithTitle(title)
	return b
}

func (b *CheckboxField) SetRequired(required bool) *CheckboxField {
	b.Required = required
	b.builder.schema.UIProps.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *CheckboxField) SetDisabled(disabled bool) *CheckboxField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.UIProps.Disabled = disabled
	return b
}

func (b *CheckboxField) SetPlaceholder(placeholder string) *CheckboxField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *CheckboxField) SetLabel(label string) *CheckboxField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *CheckboxField) SetHint(hint string) *CheckboxField {
	b.builder.schema.UIProps.Hint = hint
	return b
}
