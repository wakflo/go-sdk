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
	"github.com/gookit/goutil/arrutil"
	sdkcore "github.com/wakflo/go-sdk/core"
)

type OAuthField struct {
	*BaseComponentField
	props  map[string]sdkcore.AutoFormSchema
	extras map[string]*sdkcore.AutoFormSchema
}

func NewOAuthField(authURL string, tokenURL *string, scopes []string) *OAuthField {
	c := &OAuthField{
		BaseComponentField: NewBaseComponentField(),
		props:              map[string]sdkcore.AutoFormSchema{},
	}
	c.builder.WithType(sdkcore.Object)
	c.builder.WithFieldType(sdkcore.Oauth2Type)
	c.builder.WithDescription("Oauth2 Connection")
	c.builder.WithTitle("Oauth2 Connection")

	required := false
	c.Required = required
	c.builder.schema.Presentation.Required = required
	c.builder.schema.IsRequired = required

	c.builder.schema.AuthURL = &authURL
	c.builder.schema.TokenURL = tokenURL
	c.builder.schema.Scope = scopes

	return c
}

func (b *OAuthField) Build() *sdkcore.AutoFormSchema {
	b.schema = b.setProps().builder.Build()
	return b.schema
}

func (b *OAuthField) SetDescription(desc string) *OAuthField {
	b.builder.WithDescription(desc)
	return b
}

func (b *OAuthField) SetRequired(required bool) *OAuthField {
	b.Required = required
	b.builder.schema.Presentation.Required = required
	b.builder.schema.IsRequired = required
	return b
}

func (b *OAuthField) SetExtraFields(fields map[string]*sdkcore.AutoFormSchema) *OAuthField {
	b.builder.WithProperties(fields)
	return b
}

func (b *OAuthField) setProps() *OAuthField {
	tr := map[string]*sdkcore.AutoFormSchema{
		"appUrl": NewShortTextField().
			SetDisplayName("app url").
			SetDescription("oauth app url").
			SetDefaultValue(*b.builder.schema.AuthURL).
			SetRequired(true).Build(),

		"tokenUrl": NewShortTextField().
			SetDisplayName("token url").
			SetDescription("oauth token url").
			SetDefaultValue(*b.builder.schema.TokenURL).
			SetRequired(false).Build(),

		"scope": NewShortTextField().
			SetDisplayName("scope").
			SetDefaultValue(arrutil.JoinSlice(",", b.builder.schema.Scope)).
			SetDescription("oauth scope url").
			SetRequired(false).Build(),
	}

	f := tr
	if tr != nil {
		f = MergeMaps(tr, b.extras)
	}

	b.builder.WithProperties(f)
	return b
}

func (b *OAuthField) SetDisplayName(title string) *OAuthField {
	b.builder.WithTitle(title)
	return b
}

func MergeMaps[T any](map1, map2 map[string]T) map[string]T {
	UniqueMap := make(map[string]T)

	// for loop for the first map
	for key, val := range map1 {
		UniqueMap[key] = val
	}

	// for loop for the second map
	for key, val := range map2 {
		UniqueMap[key] = val
	}
	// return merged result
	return UniqueMap
}
