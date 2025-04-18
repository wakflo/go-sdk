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

type AuthSecretField struct {
	*BaseComponentField
	props       map[string]sdkcore.AutoFormSchema
	keyField    *sdkcore.AutoFormSchema
	secretField *BaseTextField
}

func NewAuthSecretField() *AuthSecretField {
	c := &AuthSecretField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.String)
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeSecretAuth)
	c.builder.WithDescription("Secret Connection")
	c.builder.WithTitle("Secret Connection")
	c.builder = c.builder.WithFieldRequired(true)

	return c.initProps()
}

func (b *AuthSecretField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.setProps().builder.Build()
	return b.schema
}

func (b *AuthSecretField) SetDescription(desc string) *AuthSecretField {
	b.builder.WithDescription(desc)
	return b
}

func (b *AuthSecretField) WithKey(field *sdkcore.AutoFormSchema) *AuthSecretField {
	b.keyField = field
	return b
}

func (b *AuthSecretField) setProps() *AuthSecretField {
	var order []string
	o := map[string]*sdkcore.AutoFormSchema{}

	if b.keyField != nil {
		o["key"] = b.keyField
		order = append(order, "key")
	}

	b.builder.WithProperties(map[string]*sdkcore.AutoFormSchema{
		"secret": b.secretField.Build(),
	})

	order = append(order, "secret")
	b.builder.WithOrder(order)

	return b
}

func (b *AuthSecretField) initProps() *AuthSecretField {
	name := "Secret"
	desc := "Auth secret key"

	b.secretField = NewShortTextField().
		SetDisplayName(name).
		SetLabel(name).
		SetDescription(desc).
		SetPlaceholder(desc).
		SetRequired(true)
	return b
}

func (b *AuthSecretField) SetDisplayName(title string) *AuthSecretField {
	return b.SetLabel(title)
}

func (b *AuthSecretField) SetPlaceholder(placeholder string) *AuthSecretField {
	b.secretField.SetPlaceholder(placeholder)
	return b
}

func (b *AuthSecretField) SetLabel(label string) *AuthSecretField {
	b.secretField.SetLabel(label)
	return b
}

func (b *AuthSecretField) SetHint(hint string) *AuthSecretField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *AuthSecretField) SetHidden(hidden bool) *AuthSecretField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}
