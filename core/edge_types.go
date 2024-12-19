package core

// MarkerType represents the type of a marker.
type MarkerType string

const (
	MarkerArrow       MarkerType = "arrow"
	MarkerArrowClosed MarkerType = "arrowclosed"
)

type EdgeAnimationShape string

const (
	EdgeAnimationShapeCircle  MarkerType = "circle"
	EdgeAnimationShapePackage MarkerType = "package"
)

type EdgePathType string

const (
	EdgePathTypeBezier     MarkerType = "bezier"
	EdgePathTypeSmoothstep MarkerType = "smoothstep"
	EdgePathTypeStep       MarkerType = "step"
	EdgePathTypeStraight   MarkerType = "straight"
)

// EdgeMarker represents the properties of an edge marker.
type EdgeMarker struct {
	Type        MarkerType `json:"type"`
	Color       *string    `json:"color,omitempty"`
	Width       *float64   `json:"width,omitempty"`
	Height      *float64   `json:"height,omitempty"`
	MarkerUnits *string    `json:"markerUnits,omitempty"`
	Orient      *string    `json:"orient,omitempty"`
	StrokeWidth *float64   `json:"strokeWidth,omitempty"`
}

// EdgeMarkerType is a union type that can either be a string or an EdgeMarker.
type EdgeMarkerType interface{}

// StepEdgeData holds additional data for an edge.
type StepEdgeData struct {
	Animate   *bool               `json:"animate"`
	Shape     *EdgeAnimationShape `json:"shape"`
	Path      *EdgePathType       `json:"path"`
	Direction *string             `json:"direction"`
	Repeat    bool                `json:"repeat,omitempty"`
	Duration  *int                `json:"duration,omitempty"`
	Points    []map[string]any    `json:"points,omitempty"`
	Algorithm *string             `json:"algorithm,omitempty"`
}

// StepEdge represents an edge with various properties.
type StepEdge struct {
	SharedNodesProps
	ID               string         `json:"id"`
	Type             EdgeType       `json:"type,omitempty"`
	Source           string         `json:"source"`
	Target           string         `json:"target"`
	SourceHandle     *string        `json:"sourceHandle,omitempty"`
	TargetHandle     *string        `json:"targetHandle,omitempty"`
	Animated         *bool          `json:"animated,omitempty"`
	Hidden           *bool          `json:"hidden,omitempty"`
	Deletable        *bool          `json:"deletable,omitempty"`
	Selectable       *bool          `json:"selectable,omitempty"`
	Data             *StepEdgeData  `json:"data,omitempty"`
	Selected         *bool          `json:"selected,omitempty"`
	MarkerStart      EdgeMarkerType `json:"markerStart,omitempty"`
	MarkerEnd        EdgeMarkerType `json:"markerEnd,omitempty"`
	ZIndex           *int           `json:"zIndex,omitempty"`
	AriaLabel        *string        `json:"ariaLabel,omitempty"`
	InteractionWidth *float64       `json:"interactionWidth,omitempty"`
}
