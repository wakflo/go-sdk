package core

import (
	"github.com/google/uuid"
)

// Position represents a specific location designation, typically used to define edges such as top, bottom, left, or right.
type Position string

// Different predefined positions for a layout or orientation setting.
const (
	PositionLeft   Position = "left"
	PositionTop    Position = "top"
	PositionRight  Position = "right"
	PositionBottom Position = "bottom"
)

// XYPosition represents a coordinate with X and Y floating-point values.
type XYPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// XYZPosition represents a position in 3D space, extending the 2D XYPosition with an additional Z coordinate.
type XYZPosition struct {
	XYPosition
	Z float64 `json:"z"`
}

// Dimensions represents the size of an object with a width and a height in float64 format.
type Dimensions struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Rect represents a rectangle with dimensions and a position in a 2D space.
type Rect struct {
	Dimensions
	XYPosition
}

// Box represents a rectangular area, defined by two corner points. It embeds the XYPosition struct for its first point.
type Box struct {
	XYPosition
	X2 float64 `json:"x2"`
	Y2 float64 `json:"y2"`
}

// Transform represents a 3-dimensional vector used for geometric transformations.
type Transform [3]float64

// CoordinateExtent defines a 2x2 matrix representation of coordinates, typically used for bounding rectangles in 2D space.
type CoordinateExtent [2][2]float64

// NodeOrigin represents the origin point of a node in a two-dimensional space, with coordinates stored as an array of two float64 values.
type NodeOrigin [2]float64

// StepNodeNodeAuthInput represents the authentication input for a node in a step, including a connection ID as a UUID.
type StepNodeNodeAuthInput struct {
	ConnectionID *uuid.UUID `json:"connectionId"`
}

// StepNodeFormInput represents the input structure for a form node, containing form state and input data.
type StepNodeFormInput struct {
	State FormState  `json:"state"`
	Input JSONObject `json:"input"`
}

// StepNodeTestResult represents the result of a test at a specific step in a process, including its timestamp, status, and output.
type StepNodeTestResult struct {
	Timestamp *int        `json:"timestamp"`
	Status    *TestStatus `json:"status"`
	Output    any         `json:"output"`
}

// StepNodeTest represents the testing data for a step node. It includes results, last test time, and last test status.
type StepNodeTest struct {
	Results        []StepNodeTestResult `json:"results"`
	LastTestTime   *int                 `json:"lastTestTime"`
	LastTestStatus *TestStatus          `json:"lastTestStatus"`
}

type StepNodeMeta struct {
	ConnectorName    string       `json:"connectorName,omitempty"`
	ConnectorVersion string       `json:"connectorVersion,omitempty"`
	TriggerType      *TriggerType `json:"triggerType,omitempty"`
}

type StepNodeConnector struct {
	Name        string       `json:"name,omitempty"`
	Version     string       `json:"version,omitempty"`
	TriggerType *TriggerType `json:"triggerType,omitempty"`
}

type StepNodeSettings struct {
	Trigger *StepTriggerSettings `json:"trigger,omitempty"`
	Error   StepErrorSettings    `json:"errorSettings,omitempty"`
}

type NodeRelationshipHandle struct {
	ID string `json:"id,omitempty"`
}

// StepNodeData represents the data associated with a step node, including metadata, operation IDs, authentication,
// form inputs, outputs, test results, and hierarchy information within the structure.
type StepNodeData struct {
	Meta          map[string]any           `json:"meta"`
	Connector     StepNodeConnector        `json:"connector"`
	Settings      StepNodeSettings         `json:"settings"`
	OperationID   *string                  `json:"operationId"`
	Auth          *StepNodeNodeAuthInput   `json:"auth"`
	Form          StepNodeFormInput        `json:"form"`
	Output        any                      `json:"output"`
	Tests         StepNodeTest             `json:"tests"`
	ParentID      *string                  `json:"parentId"`
	Label         string                   `json:"label"`
	Icon          string                   `json:"icon,omitempty" validate:"required"`
	IsTrigger     bool                     `json:"isTrigger" validate:"required"`
	IsFolded      bool                     `json:"isFolded,omitempty"`
	SourceHandles []NodeRelationshipHandle `json:"sourceHandles,omitempty"`
	TargetHandles []NodeRelationshipHandle `json:"targetHandles,omitempty"`
}

// NodeHandle represents a handle for a node, typically used to manage connections or references between nodes.
type NodeHandle struct {
	HandleID string `json:"handleId"` // Example field, please adjust as per actual use case
}

// SharedNodesProps defines properties that can be shared among different node elements in a graph or similar structures.
type SharedNodesProps struct {
	Style     map[string]any `json:"style,omitempty"`
	ClassName *string        `json:"className,omitempty"`
	Resizing  *bool          `json:"resizing,omitempty"`
	Focusable *bool          `json:"focusable,omitempty"`
}

// StepNode represents a node in a graph with metadata and positioning information.
type StepNode struct {
	SharedNodesProps
	ID             string       `json:"id"`
	Position       XYPosition   `json:"position"`
	Data           StepNodeData `json:"data"`
	Type           NodeType     `json:"type,omitempty"`
	SourcePosition *Position    `json:"sourcePosition,omitempty"`
	TargetPosition *Position    `json:"targetPosition,omitempty"`
	Hidden         *bool        `json:"hidden,omitempty"`
	Selected       *bool        `json:"selected,omitempty"`
	Dragging       *bool        `json:"dragging,omitempty"`
	Draggable      *bool        `json:"draggable,omitempty"`
	Selectable     *bool        `json:"selectable,omitempty"`
	Deletable      *bool        `json:"deletable,omitempty"`
	Connectable    *bool        `json:"connectable,omitempty"`
	DragHandle     *string      `json:"dragHandle,omitempty"`
	Width          *float64     `json:"width,omitempty"`
	Height         *float64     `json:"height,omitempty"`
	InitialWidth   *float64     `json:"initialWidth,omitempty"`
	InitialHeight  *float64     `json:"initialHeight,omitempty"`
	ParentID       *string      `json:"parentId,omitempty"`
	ZIndex         *int         `json:"zIndex,omitempty"`
	Extent         any          `json:"extent,omitempty"` // Support for both 'parent' string and CoordinateExtent
	ExpandParent   *bool        `json:"expandParent,omitempty"`
	AriaLabel      *string      `json:"ariaLabel,omitempty"`
	Origin         *NodeOrigin  `json:"origin,omitempty"`
	Handles        []NodeHandle `json:"handles,omitempty"`
	Measured       *Dimensions  `json:"measured,omitempty"`
	Path           []string     `json:"path,omitempty"`
	NodeIndex      int          `json:"nodeIndex,omitempty"`
}

// StepNodeWithChildren extends StepNode, allowing for hierarchical representation with children nodes.
type StepNodeWithChildren struct {
	StepNode
	Children []StepNodeWithChildren `json:"children,omitempty"`
}

// ToTree constructs a hierarchical tree from a slice of StepEdge, creating StepNodeWithChildren as tree nodes.
func (n *StepNode) ToTree(edges []StepEdge) []StepNodeWithChildren {
	nodeMap := make(map[string]*StepNodeWithChildren)
	var buildTree func(string) *StepNodeWithChildren
	buildTree = func(nodeID string) *StepNodeWithChildren {
		node := nodeMap[nodeID]
		if node == nil {
			node = &StepNodeWithChildren{StepNode: StepNode{ID: nodeID}}
			nodeMap[nodeID] = node
		}
		for _, edge := range edges {
			if edge.Source == nodeID {
				childNode := buildTree(edge.Target)
				node.Children = append(node.Children, *childNode)
			}
		}
		return node
	}

	rootNodes := []StepNodeWithChildren{}
	for _, edge := range edges {
		if _, exists := nodeMap[edge.Source]; !exists {
			root := buildTree(edge.Source)
			rootNodes = append(rootNodes, *root)
		}
	}
	return rootNodes
}

// StepNodesToTree constructs a hierarchical tree from a slice of StepEdge, creating StepNodeWithChildren as tree nodes.
func StepNodesToTree(nodes []StepNode, edges []StepEdge) []StepNodeWithChildren {
	nodeMap := make(map[string]*StepNodeWithChildren)

	// Populate the node map with StepNodeWithChildren
	for _, node := range nodes {
		nodeWithChildren := StepNodeWithChildren{StepNode: node}
		nodeMap[node.ID] = &nodeWithChildren
	}

	rootNodes := []StepNodeWithChildren{}

	// Build the tree by associating children with their parents
	for _, node := range nodes {
		if node.Data.ParentID != nil {
			parentID := *node.Data.ParentID
			if parent, exists := nodeMap[parentID]; exists {
				parent.Children = append(parent.Children, *nodeMap[node.ID])
			}
		} else {
			// Nodes without a ParentID are considered root nodes
			rootNodes = append(rootNodes, *nodeMap[node.ID])
		}
	}

	return rootNodes
}
