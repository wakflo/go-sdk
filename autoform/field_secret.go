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

type AuthSecretField struct {
	*BaseComponentField
	props map[string]sdkcore.AutoFormSchema
}

func NewAuthSecretField() *AuthSecretField {
	c := &AuthSecretField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.String)
	c.builder.WithFieldType(sdkcore.Secret)
	c.builder.WithDescription("Secret Connection")
	c.builder.WithTitle("Secret Connection")

	required := false
	c.Required = required
	c.builder.schema.Presentation.Required = required
	c.builder.schema.IsRequired = required

	return c
}

func (b *AuthSecretField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.setProps().builder.Build()
	return b.schema
}

func (b *AuthSecretField) SetDescription(desc string) *AuthSecretField {
	b.builder.WithDescription(desc)
	return b
}

func (b *AuthSecretField) setProps() *AuthSecretField {
	b.builder.WithProperties(map[string]*sdkcore.AutoFormSchema{
		"secret": NewShortTextField().
			SetDisplayName("Secret").
			SetDescription("auth secret key").
			SetDefaultValue(*b.builder.schema.AuthUrl).
			SetRequired(true).Build(),
	})
	return b
}

func (b *AuthSecretField) SetDisplayName(title string) *AuthSecretField {
	b.builder.WithTitle(title)
	return b
}
