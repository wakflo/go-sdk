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

type CodeEditorField struct {
	*BaseComponentField
	props map[string]sdkcore.AutoFormSchema
}

func NewCodeEditorField(language sdkcore.CodeEditorLanguage) *CodeEditorField {
	c := &CodeEditorField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.String)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeCodeEditor)
	c.builder.schema.UIProps.Language = string(language)

	return c
}

func (b *CodeEditorField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *CodeEditorField) SetLanguage(language sdkcore.CodeEditorLanguage) *CodeEditorField {
	b.builder.schema.UIProps.Language = string(language)
	return b
}

func (b *CodeEditorField) SetDescription(desc string) *CodeEditorField {
	b.builder.WithDescription(desc)
	return b
}

func (b *CodeEditorField) SetDisplayName(title string) *CodeEditorField {
	b.builder.WithTitle(title)
	return b
}

func (b *CodeEditorField) SetRequired(required bool) *CodeEditorField {
	b.Required = required
	b.builder.WithFieldRequired(required)
	return b
}

func (b *CodeEditorField) SetDisabled(disabled bool) *CodeEditorField {
	b.builder = b.builder.WithFieldDisabled(disabled)
	return b
}

func (b *CodeEditorField) SetDefaultValue(defaultValue string) *CodeEditorField {
	b.builder.WithDefault(defaultValue)
	return b
}

func (b *CodeEditorField) SetPlaceholder(placeholder string) *CodeEditorField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *CodeEditorField) SetLabel(label string) *CodeEditorField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *CodeEditorField) SetHint(hint string) *CodeEditorField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *CodeEditorField) SetHidden(hidden bool) *CodeEditorField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
