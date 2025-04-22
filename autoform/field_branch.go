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

type BranchField struct {
	*BaseComponentField
	props map[string]sdkcore.AutoFormSchema
}

func NewBranchField() *BranchField {
	c := &BranchField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.Array)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeBranch)

	return c
}

func (b *BranchField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.builder.Build()
	return b.schema
}

func (b *BranchField) SetDescription(desc string) *BranchField {
	b.builder.WithDescription(desc)
	return b
}

func (b *BranchField) SetDisplayName(title string) *BranchField {
	b.builder.WithTitle(title)
	return b
}

func (b *BranchField) SetRequired(required bool) *BranchField {
	b.Required = required
	b.builder.WithFieldRequired(required)
	return b
}

func (b *BranchField) SetDisabled(disabled bool) *BranchField {
	b.builder = b.builder.WithFieldDisabled(disabled)
	return b
}

func (b *BranchField) SetPlaceholder(placeholder string) *BranchField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *BranchField) SetLabel(label string) *BranchField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *BranchField) SetHint(hint string) *BranchField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *BranchField) SetHidden(hidden bool) *BranchField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
