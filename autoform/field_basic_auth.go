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

// AuthBasicField type is a composite structure for building basic authentication fields in a form schema.
type AuthBasicField struct {
	*BaseComponentField
	props map[string]sdkcore.AutoFormSchema

	usernameField *BaseTextField
	passwordField *BaseTextField
}

// NewAuthBasicField initializes and returns a new AuthBasicField with default properties for basic authentication.
func NewAuthBasicField() *AuthBasicField {
	c := &AuthBasicField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.Object)
	c.builder.WithTitle("Basic Auth")
	c.builder.WithDescription("Basic Auth")
	c.builder.WithFieldRequired(true)

	return c.initProps()
}

// Build constructs and returns the final AutoFormSchema for an AuthBasicField, setting its properties and order.
func (b *AuthBasicField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.setProps().builder.Build()
	return b.schema
}

// SetDescription sets the description for the AuthBasicField. Returns the updated AuthBasicField.
func (b *AuthBasicField) SetDescription(desc string) *AuthBasicField {
	b.builder.WithDescription(desc)
	return b
}

// setProps configures the authentication fields and their properties and order for the AuthBasicField instance.
func (b *AuthBasicField) setProps() *AuthBasicField {
	b.builder.WithProperties(map[string]*sdkcore.AutoFormSchema{
		"username": b.usernameField.Build(),
		"password": b.passwordField.Build(),
	})

	b.builder.WithOrder([]string{"username", "password"})

	return b
}

// initProps initializes the properties for AuthBasicField, setting default values for username and password fields.
func (b *AuthBasicField) initProps() *AuthBasicField {
	b.passwordField = NewShortTextField().
		SetLabel("Username").
		SetPlaceholder("The auth username").
		SetRequired(true)

	b.usernameField = NewShortTextField().
		SetLabel("Password").
		SetPlaceholder("The auth password").
		SetRequired(true)

	return b
}

// SetUsernamePlaceholder sets the placeholder text for the username input field and returns the updated AuthBasicField instance.
func (b *AuthBasicField) SetUsernamePlaceholder(placeholder string) *AuthBasicField {
	b.usernameField.SetPlaceholder(placeholder)
	return b
}

// SetPasswordPlaceholder sets the placeholder text for the password field in the AuthBasicField component.
func (b *AuthBasicField) SetPasswordPlaceholder(placeholder string) *AuthBasicField {
	b.passwordField.SetPlaceholder(placeholder)
	return b
}

// SetUsernameLabel sets the label for the username field in the authentication form.
func (b *AuthBasicField) SetUsernameLabel(label string) *AuthBasicField {
	b.usernameField.SetLabel(label)
	return b
}

// SetPasswordLabel sets the label for the password field in the AuthBasicField and returns the updated AuthBasicField.
func (b *AuthBasicField) SetPasswordLabel(label string) *AuthBasicField {
	b.passwordField.SetLabel(label)
	return b
}

// SetUsernameHint sets a hint for the username field and returns the updated AuthBasicField object.
func (b *AuthBasicField) SetUsernameHint(hint string) *AuthBasicField {
	b.usernameField.SetHint(hint)
	return b
}

// SetPasswordHint sets a hint for the password field and returns the AuthBasicField object for method chaining.
func (b *AuthBasicField) SetPasswordHint(hint string) *AuthBasicField {
	b.passwordField.SetHint(hint)
	return b
}
