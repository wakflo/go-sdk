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

type MarkdownField struct {
	*BaseTextField
}

func NewMarkdownField() *MarkdownField {
	c := &MarkdownField{
		BaseTextField: newBaseTextField(),
	}
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeMarkdown)

	return c
}

func (b *MarkdownField) SetPlaceholder(placeholder string) *MarkdownField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *MarkdownField) SetLabel(label string) *MarkdownField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *MarkdownField) SetHint(hint string) *MarkdownField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *MarkdownField) SetHidden(hidden bool) *MarkdownField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
