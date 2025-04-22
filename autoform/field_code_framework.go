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

type CodeFrameworkField struct {
	*BaseComponentField
	props map[string]sdkcore.AutoFormSchema
}

func NewCodeFrameworkField() *CodeFrameworkField {
	c := &CodeFrameworkField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.Object)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeCodeFramework)
	c.builder.schema.UIProps.Framework = sdkcore.CodeFrameworkNode

	return c
}

func (b *CodeFrameworkField) Build() *sdkcore.AutoFormSchema {
	b.builder.WithDefault(sdkcore.CodeFrameworkDefaults[b.builder.schema.UIProps.Framework])
	b.schema = b.builder.Build()
	return b.schema
}

func (b *CodeFrameworkField) SetFramework(framework sdkcore.CodeFramework) *CodeFrameworkField {
	b.builder.schema.UIProps.Framework = framework
	return b
}

func (b *CodeFrameworkField) SetDescription(desc string) *CodeFrameworkField {
	b.builder.WithDescription(desc)
	return b
}

func (b *CodeFrameworkField) SetDisplayName(title string) *CodeFrameworkField {
	b.builder.WithTitle(title)
	return b
}

func (b *CodeFrameworkField) SetRequired(required bool) *CodeFrameworkField {
	b.Required = required
	b.builder.WithFieldRequired(required)
	return b
}

func (b *CodeFrameworkField) SetDisabled(disabled bool) *CodeFrameworkField {
	b.builder = b.builder.WithFieldDisabled(disabled)
	return b
}

func (b *CodeFrameworkField) setDefaultValue(defaultValue sdkcore.CodeFrameworkProps) *CodeFrameworkField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *CodeFrameworkField) SetPlaceholder(placeholder string) *CodeFrameworkField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *CodeFrameworkField) SetLabel(label string) *CodeFrameworkField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *CodeFrameworkField) SetHint(hint string) *CodeFrameworkField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *CodeFrameworkField) SetHidden(hidden bool) *CodeFrameworkField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
