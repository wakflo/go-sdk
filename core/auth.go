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

package core

import (
	"time"

	"github.com/grokify/goauth/authutil"
	"github.com/grokify/goauth/google"
	"golang.org/x/oauth2"
)

type BasicAuthSrategy struct{}

type WakfloAuthStrategy struct{}

type AuthStrategy struct {
	google.Credentials
	authutil.AuthorizationType
}

type AuthState struct {
	ClientID     *string             `json:"clientId,omitempty"`
	Code         *string             `json:"code,omitempty"`
	ClientSecret *string             `json:"clientSecret,omitempty"`
	AuthURL      *string             `json:"authUrl,omitempty"`
	TokenURL     *string             `json:"tokenUrl,omitempty"`
	Username     *string             `json:"username,omitempty"`
	Password     *string             `json:"password,omitempty"`
	Secret       *string             `json:"secret,omitempty"`
	Key          *string             `json:"key,omitempty"`
	RedirectURL  *string             `json:"redirectURL,omitempty"`
	GrantType    *string             `json:"grantType,omitempty"`
	AccessToken  *string             `json:"accessToken,omitempty"`
	RefreshToken *string             `json:"refreshToken,omitempty"`
	Scopes       []string            `json:"scopes,omitempty"`
	TokenType    *string             `json:"tokenType,omitempty"`
	Endpoint     *oauth2.Endpoint    `json:"endpoint,omitempty"`
	Expiry       *time.Time          `json:"expiry,omitempty"`
	Extra        map[string][]string `json:"extra,omitempty"`
}

type AuthContext struct {
	AccessToken string              `json:"accessToken,omitempty"`
	Token       *oauth2.Token       `json:"token,omitempty"`
	TokenSource *oauth2.TokenSource `json:"tokenSource,omitempty"`
	TokenType   string              `json:"tokenType,omitempty"`
	Username    string              `json:"username,omitempty"`
	Password    string              `json:"password,omitempty"`
	Secret      string              `json:"secret,omitempty"`
	Key         string              `json:"key,omitempty"`
	Extra       map[string][]string `json:"extra,omitempty"`
}

type ConnectorAuthMetadata struct {
	ClientID     *string             `json:"clientId,omitempty"`
	ClientSecret *string             `json:"clientSecret,omitempty"`
	AuthURL      *string             `json:"authUrl,omitempty"`
	TokenURL     *string             `json:"tokenUrl,omitempty"`
	Username     *string             `json:"username,omitempty"`
	Password     *string             `json:"password,omitempty"`
	Secret       *string             `json:"secret,omitempty"`
	RedirectURL  *string             `json:"redirectUrl,omitempty"`
	Key          *string             `json:"key,omitempty"`
	Extra        map[string][]string `json:"extra,omitempty"`
}
