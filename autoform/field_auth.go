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

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (b *Auth) NewOauth2Auth(authURL string, tokenURL *string, scopes []string) *OAuthField {
	return NewOAuthField(authURL, tokenURL, scopes)
}

func (b *Auth) NewSecretAuth() *AuthSecretField {
	return NewAuthSecretField()
}

func (b *Auth) NewBasicAuth() *AuthSecretField {
	return NewAuthSecretField()
}

func (b *Auth) NewCustomAuth() *CustomAuthField {
	return NewCustomAuthField()
}
