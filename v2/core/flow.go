package core

import (
	"database/sql/driver"
	"fmt"
)

// FlowComponentType defines the specific type of flow component
type FlowComponentType string

const (
	// FlowComponentTypeBranch represents a branching flow component
	FlowComponentTypeBranch FlowComponentType = "BRANCH"

	// FlowComponentTypeLoop represents a loop flow component
	FlowComponentTypeLoop FlowComponentType = "LOOP"

	// FlowComponentTypeCondition represents a condition flow component
	FlowComponentTypeCondition FlowComponentType = "CONDITION"

	// FlowComponentTypeRouter represents a router flow component
	FlowComponentTypeRouter FlowComponentType = "ROUTER"

	// FlowComponentTypeApproval represents an approval flow component
	FlowComponentTypeApproval FlowComponentType = "APPROVAL"

	// FlowComponentTypeDelay represents a delay flow component
	FlowComponentTypeDelay FlowComponentType = "DELAY"

	// FlowComponentTypeSubflow represents a subflow flow component
	FlowComponentTypeSubflow FlowComponentType = "SUBFLOW"

	// FlowComponentTypeParallel represents a subflow flow component
	FlowComponentTypeParallel FlowComponentType = "PARALLEL"

	FlowComponentTypeSequential FlowComponentType = "SEQUENTIAL"
)

// Values provides list valid values for Enum.
func (FlowComponentType) Values() []string {
	return []string{
		string(FlowComponentTypeBranch),
		string(FlowComponentTypeLoop),
		string(FlowComponentTypeCondition),
		string(FlowComponentTypeRouter),
		string(FlowComponentTypeApproval),
		string(FlowComponentTypeDelay),
		string(FlowComponentTypeSubflow),
		string(FlowComponentTypeParallel),
		string(FlowComponentTypeSequential),
	}
}

// Value implements the driver.Valuer interface.
func (s FlowComponentType) Value() (driver.Value, error) {
	return string(s), nil
}

// Scan implements the sql.Scanner interface.
func (s *FlowComponentType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	var st string
	switch v := src.(type) {
	case string:
		st = v
	case []byte:
		st = string(v)
	default:
		return fmt.Errorf("invalid type for FlowComponentType: %T", src)
	}
	*s = FlowComponentType(st)
	return nil
}
