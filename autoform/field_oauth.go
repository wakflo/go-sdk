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
	"strings"

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
	c.builder.WithFieldType(sdkcore.AutoFormFieldTypeOauth2)
	c.builder.WithDescription("Oauth2 Connection")
	c.builder.WithTitle("Oauth2 Connection")

	c.builder = c.builder.WithFieldRequired(true)
	c.builder.schema.UIProps.Auth = &sdkcore.AuthSchemaProps{
		AuthURL:          &authURL,
		TokenURL:         tokenURL,
		Scope:            scopes,
		RedirectParamKey: "redirect_uri", 
	}

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
	b.builder = b.builder.WithFieldRequired(true)
	return b
}

func (b *OAuthField) SetExtraFields(fields map[string]*sdkcore.AutoFormSchema) *OAuthField {
	b.builder.WithProperties(fields)
	return b
}

func (b *OAuthField) SetPlaceholder(placeholder string) *OAuthField {
	b.builder.schema.UIProps.Placeholder = placeholder
	return b
}

func (b *OAuthField) SetLabel(label string) *OAuthField {
	b.builder.WithTitle(label)
	b.builder.schema.UIProps.Label = label
	return b
}

func (b *OAuthField) SetHint(hint string) *OAuthField {
	b.builder.schema.UIProps.Hint = hint
	return b
}

func (b *OAuthField) setProps() *OAuthField {
	scope := strings.Join(b.builder.schema.UIProps.Auth.Scope, ",")
	tr := map[string]*sdkcore.AutoFormSchema{
		"authUrl": NewShortTextField().
			SetDisplayName("Auth URL").
			SetDescription("oauth auth url").
			SetDefaultValue(*b.builder.schema.UIProps.Auth.AuthURL).
			SetRequired(true).Build(),
		"tokenUrl": NewShortTextField().
			SetDisplayName("Token URL").
			SetDescription("oauth token url").
			SetDefaultValue(*b.builder.schema.UIProps.Auth.TokenURL).
			SetRequired(false).Build(),
		"scopes": NewShortTextField().
			SetDisplayName("Scopes").
			SetDefaultValue(scope).
			SetDescription("oauth scope url").
			SetRequired(false).Build(),
		"clientId": NewShortTextField().
			SetDisplayName("Client ID").
			SetDescription("oauth client id").
			SetRequired(true).Build(),
		"clientSecret": NewShortTextField().
			SetDisplayName("Client Secret").
			SetDescription("Oauth client secret").
			SetRequired(true).Build(),
	}

	f := tr
	if b.extras != nil {
		f = MergeMaps(tr, b.extras)
	}

	b.builder.WithProperties(f)
	b.builder.WithOrder([]string{
		"appUrl",
		"tokenUrl",
		"scope",
		"clientId",
		"clientSecret",
	})
	return b
}

func (b *OAuthField) SetDisplayName(title string) *OAuthField {
	b.builder.WithTitle(title)
	return b
}

func (b *OAuthField) SetHidden(hidden bool) *OAuthField {
	b.builder.schema.UIProps.Hidden = hidden
	return b
}

func (b *OAuthField) SetExcludedQueryParams(params []string) *OAuthField {
	if b.builder.schema.UIProps.Auth == nil {
		b.builder.schema.UIProps.Auth = &sdkcore.AuthSchemaProps{}
	}
	b.builder.schema.UIProps.Auth.ExcludedParams = params
	return b
}

func (b *OAuthField) SetRedirectQueryParamKey(key string) *OAuthField {
	if b.builder.schema.UIProps.Auth == nil {
		b.builder.schema.UIProps.Auth = &sdkcore.AuthSchemaProps{}
	}

	b.builder.schema.UIProps.Auth.RedirectParamKey = key
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
