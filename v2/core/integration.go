package core

// IntegrationBuildMetadata is a build metadata for integrations.
type IntegrationBuildMetadata struct {
	Platform IntegrationPlatform `json:"compiler"`
	Language IntegrationLanguage `json:"language,omitempty"`
}
