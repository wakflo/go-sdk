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

package context

import (
	"errors"

	"golang.org/x/oauth2"
)

// AuthContext provides authentication data for API calls
type AuthContext struct {
	// Access token for OAuth2 and token-based authentication
	AccessToken string `json:"accessToken,omitempty"`

	// Full OAuth2 token data
	Token *oauth2.Token `json:"token,omitempty"`

	// Token source for refreshing OAuth2 tokens
	TokenSource *oauth2.TokenSource `json:"tokenSource,omitempty"`

	// Token type (Bearer, MAC, etc.)
	TokenType string `json:"tokenType,omitempty"`

	// Basic auth credentials
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	// API key and secret
	Secret string `json:"secret,omitempty"`
	Key    string `json:"key,omitempty"`

	// OAuth2 scopes
	Scopes []string `json:"scopes,omitempty"`

	// Additional parameters for custom auth types
	Extra map[string]string `json:"extra,omitempty"`
}

// GetExtra retrieves additional parameters from the auth context
func (c *AuthContext) GetExtra() (map[string]string, error) {
	if c.Extra == nil {
		return nil, errors.New("extra info not found in context")
	}

	return c.Extra, nil
}

// GetCustomAuth retrieves custom authentication data
func (c *AuthContext) GetCustomAuth() (map[string]string, error) {
	if c.Extra == nil {
		return nil, errors.New("custom auth data not provided")
	}

	return c.Extra, nil
}

// NewAuthContext creates a basic AuthContext with the given token
func NewAuthContext(accessToken string) *AuthContext {
	return &AuthContext{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		Extra:       make(map[string]string),
	}
}

// WithUsernamePassword adds basic auth credentials to the context
func (c *AuthContext) WithUsernamePassword(username, password string) *AuthContext {
	c.Username = username
	c.Password = password
	return c
}

// WithKey adds an API key to the context
func (c *AuthContext) WithKey(key string) *AuthContext {
	c.Key = key
	return c
}

// WithSecret adds a secret to the context
func (c *AuthContext) WithSecret(secret string) *AuthContext {
	c.Secret = secret
	return c
}

// WithScopes adds OAuth2 scopes to the context
func (c *AuthContext) WithScopes(scopes []string) *AuthContext {
	c.Scopes = scopes
	return c
}

// WithExtraParam adds a parameter to the context
func (c *AuthContext) WithExtraParam(key, value string) *AuthContext {
	if c.Extra == nil {
		c.Extra = make(map[string]string)
	}
	c.Extra[key] = value
	return c
}
