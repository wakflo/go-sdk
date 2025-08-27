package flow

import (
	"github.com/juicycleff/smartform/v1"
	"github.com/wakflo/go-sdk/v2/core"
)

// FlowMetadata contains metadata about a flow component.
type FlowMetadata struct {
	// ID is the unique identifier for this flow component
	ID string `json:"id"`

	// DisplayName is the human-readable name of the flow component
	DisplayName string `json:"displayName"`

	// Description provides details about the flow component's purpose
	Description string `json:"description"`

	// HelpText provides additional guidance for configuring the flow component
	HelpText string `json:"helpText,omitempty"`

	// Icon is a URL or base64-encoded image for the component icon
	Icon string `json:"icon,omitempty"`

	// Type specifies the flow component type (e.g., BRANCH, LOOP, CONDITION)
	Type core.FlowComponentType `json:"type"`

	// Documentation provides comprehensive usage instructions
	Documentation string `json:"documentation,omitempty"`

	// Tags are searchable labels for the flow component
	Tags []string `json:"tags,omitempty"`

	// Category provides a categorization for this flow component
	Categories []string `json:"categories,omitempty"`

	// Version specifies the version of the flow component.
	Version string `json:"version"`

	// Settings contains the settings for the flow component
	Settings core.ActionSettings `json:"settings,omitempty"`

	Authors []string `json:"author"`

	Website string `json:"website"`

	SampleOutput map[string]interface{} `json:"sampleOutput,omitempty"`
}

// Flow represents a flow component that can be registered and executed
type Flow interface {
	// Metadata returns metadata about the flow component
	Metadata() FlowMetadata

	// Properties returns the schema for the flow component's input configuration
	Properties() *smartform.FormSchema
}
