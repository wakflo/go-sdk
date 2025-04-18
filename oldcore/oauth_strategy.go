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

package oldcore

type OauthStrategy struct {
	ClientID          string            `json:"client_id"`
	ClientSecret      string            `json:"client_secret"`
	RedirectURL       string            `json:"redirect_url"`
	Scopes            []string          `json:"scopes"`
	Endpoint          AuthEndpoint      `json:"endpoint"`
	ResponseType      string            `json:"response_type"`
	GrantType         string            `json:"grant_type"`
	State             string            `json:"state"`
	AuthMethod        string            `json:"auth_method"`
	CodeChallenge     string            `json:"code_challenge"`
	CodeChallengeMath string            `json:"code_challenge_math"`
	PrivateKey        string            `json:"private_key"`
	ExtraParams       map[string]string `json:"extra_params"`
}

type AuthEndpoint struct {
	AuthURL          string `json:"auth_url,omitempty"`
	TokenURL         string `json:"token_url,omitempty"`
	RevocationURL    string `json:"revocation_url,omitempty"`
	IntrospectionURL string `json:"introspection_url,omitempty"`
	JWKSetURL        string `json:"jwk_set_url,omitempty"`
}
type OauthTokenResponse struct {
	AccessToken  string            `json:"access_token,omitempty"`
	TokenType    string            `json:"token_type,omitempty"`
	ExpiresIn    int64             `json:"expires_in,omitempty"`
	RefreshToken string            `json:"refresh_token,omitempty"`
	Scope        string            `json:"scope,omitempty"`
	Extra        map[string]string `json:"extra,omitempty"`
}
