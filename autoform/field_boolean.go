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

type BooleanField struct {
	*BaseComponentField
}

func NewBooleanField() *BooleanField {
	c := &BooleanField{
		BaseComponentField: NewBaseComponentField(),
	}
	c.builder.WithType(sdkcore.Boolean)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeBoolean)

	return c
}

func (b *BooleanField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *BooleanField) SetDefaultValue(defaultValue bool) *BooleanField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *BooleanField) SetDescription(desc string) *BooleanField {
	b.builder.WithDescription(desc)
	return b
}

func (b *BooleanField) SetDisplayName(title string) *BooleanField {
	b.builder.WithTitle(title)
	return b
}

func (b *BooleanField) SetRequired(required bool) *BooleanField {
	b.Required = required
	b.builder.schema.UIProps.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *BooleanField) SetDisabled(disabled bool) *BooleanField {
	b.builder.schema.Disabled = disabled
	b.builder.schema.UIProps.Disabled = disabled
	return b
}

func (b *BooleanField) SetPlaceholder(placeholder string) *BooleanField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *BooleanField) SetLabel(label string) *BooleanField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *BooleanField) SetHint(hint string) *BooleanField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *BooleanField) SetHidden(hidden bool) *BooleanField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
