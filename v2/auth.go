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

package sdk

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/xid"
	wakcontext "github.com/wakflo/go-sdk/v2/context"
	"github.com/wakflo/go-sdk/v2/core"
	"golang.org/x/oauth2"
)

// AuthRequest represents a request to initialize an authentication flow.
type AuthRequest struct {
	// IntegrationID is the identifier for the integration requiring authentication
	IntegrationID string `json:"integrationId"`

	// ProjectID is the identifier for the project this authentication belongs to
	ProjectID xid.ID `json:"projectId"`

	// UserID is the identifier for the user initiating the authentication
	UserID xid.ID `json:"userId,omitempty"`

	// Type specifies the authentication method to use
	Type core.AuthType `json:"type"`

	// ConnectionName is a user-defined name for this connection
	ConnectionName string `json:"connectionName"`

	// Credentials contains authentication credentials (for non-OAuth flows)
	Credentials map[string]interface{} `json:"credentials,omitempty"`

	// Metadata contains additional information about this authentication
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// RedirectURL is where to redirect after OAuth authentication
	RedirectURL string `json:"redirectUrl,omitempty"`

	// Scopes contains the OAuth scopes to request
	Scopes []string `json:"scopes,omitempty"`
}

// AuthResponse represents the result of an authentication flow.
type AuthResponse struct {
	// ConnectionID is the unique identifier for this authentication connection
	ConnectionID xid.ID `json:"connectionId"`

	// Status indicates the current status of this connection
	Status AuthConnectionStatus `json:"status"`

	// Type indicates the authentication method used
	Type core.AuthType `json:"type"`

	// ExpiresAt indicates when this authentication expires (if applicable)
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`

	// Message contains additional information about the authentication result
	Message string `json:"message,omitempty"`

	// AuthorizationURL is the URL to redirect to for OAuth flows
	AuthorizationURL string `json:"authorizationUrl,omitempty"`

	// ConnectionDetails contains details about the connection (sanitized for display)
	ConnectionDetails map[string]interface{} `json:"connectionDetails,omitempty"`
}

// AuthConnectionStatus represents the status of an authentication connection.
type AuthConnectionStatus string

const (
	// AuthStatusPending indicates the authentication flow has been initiated but not completed.
	AuthStatusPending AuthConnectionStatus = "pending"

	// AuthStatusActive indicates the authentication is active and usable.
	AuthStatusActive AuthConnectionStatus = "active"

	// AuthStatusExpired indicates the authentication has expired and needs renewal.
	AuthStatusExpired AuthConnectionStatus = "expired"

	// AuthStatusRevoked indicates the authentication has been explicitly revoked.
	AuthStatusRevoked AuthConnectionStatus = "revoked"

	// AuthStatusFailed indicates the authentication attempt failed.
	AuthStatusFailed AuthConnectionStatus = "failed"
)

// Auth defines the interface for authentication operations.
type Auth interface {
	// Initialize starts an authentication flow
	Initialize(ctx context.Context, request AuthRequest) (*AuthResponse, error)

	// CompleteOAuth finishes an OAuth flow with a code from the OAuth provider
	CompleteOAuth(ctx context.Context, connectionID xid.ID, code string, state string) (*AuthResponse, error)

	// Refresh renews an authentication token
	Refresh(ctx context.Context, connectionID xid.ID) (*AuthResponse, error)

	// Revoke explicitly revokes an authentication connection
	Revoke(ctx context.Context, connectionID xid.ID) error

	// Validate checks if an authentication connection is valid
	Validate(ctx context.Context, connectionID xid.ID) (*AuthResponse, error)

	// GetConnection retrieves authentication connection details
	GetConnection(ctx context.Context, connectionID xid.ID) (*AuthResponse, error)

	// ListConnections retrieves all authentication connections for a project
	ListConnections(ctx context.Context, projectID xid.ID) ([]AuthResponse, error)

	// GetToken retrieves an OAuth token for use in API calls
	GetToken(ctx context.Context, connectionID xid.ID) (*oauth2.Token, error)

	// GetAuthContext retrieves the auth context for use in integration operations
	GetAuthContext(ctx context.Context, connectionID xid.ID) (*wakcontext.AuthContext, error)

	// CreateAuthenticatedClient creates an HTTP client with authentication credentials
	CreateAuthenticatedClient(ctx context.Context, connectionID xid.ID) (*http.Client, error)
}
