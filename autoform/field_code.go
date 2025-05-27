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

type CodeField struct {
	*BaseComponentField
	props map[string]sdkcore.AutoFormSchema
}

func NewCodeEditorField() *CodeField {
	c := &CodeField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.String)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeCode)
	c.builder.schema.UIProps.Language = sdkcore.CodeLanguageJavascript

	return c
}

func (b *CodeField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *CodeField) SetLanguage(language sdkcore.CodeLanguage) *CodeField {
	b.builder.schema.UIProps.Language = language
	return b
}

func (b *CodeField) SetDescription(desc string) *CodeField {
	b.builder.WithDescription(desc)
	return b
}

func (b *CodeField) SetDisplayName(title string) *CodeField {
	b.builder.WithTitle(title)
	return b
}

func (b *CodeField) SetRequired(required bool) *CodeField {
	b.Required = required
	b.builder.WithFieldRequired(required)
	return b
}

func (b *CodeField) SetDisabled(disabled bool) *CodeField {
	b.builder = b.builder.WithFieldDisabled(disabled)
	return b
}

func (b *CodeField) SetDefaultValue(defaultValue string) *CodeField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *CodeField) SetPlaceholder(placeholder string) *CodeField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *CodeField) SetLabel(label string) *CodeField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *CodeField) SetHint(hint string) *CodeField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *CodeField) SetHidden(hidden bool) *CodeField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
