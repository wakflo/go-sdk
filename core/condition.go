package core

import (
	"fmt"
)

import (
	"encoding/json"
	"io"
)

// Operator Enum
type Operator string

const (
	OperatorAND Operator = "AND"
	OperatorOR  Operator = "OR"
	OperatorXOR Operator = "XOR"
)

func (o *Operator) UnmarshalJSON(data []byte) error {
	str := string(data[1 : len(data)-1]) // Trim quotes
	switch str {
	case string(OperatorAND), string(OperatorOR), string(OperatorXOR):
		*o = Operator(str)
		return nil
	default:
		return fmt.Errorf("invalid Operator value: %s", str)
	}
}

func (o Operator) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(o))
}

func (o Operator) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, o)))
}

func (o *Operator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Operator must be a string")
	}
	switch str {
	case string(OperatorAND), string(OperatorOR), string(OperatorXOR):
		*o = Operator(str)
		return nil
	default:
		return fmt.Errorf("invalid Operator value: %s", str)
	}
}

// DataType Enum
type DataType string

const (
	DataTypeString  DataType = "string"
	DataTypeNumber  DataType = "number"
	DataTypeBoolean DataType = "boolean"
	DataTypeDate    DataType = "date"
	DataTypeList    DataType = "list"
	DataTypeObject  DataType = "object" // Newly added
)

func (dt *DataType) UnmarshalJSON(data []byte) error {
	str := string(data[1 : len(data)-1]) // Trim quotes
	switch str {
	case string(DataTypeString), string(DataTypeNumber), string(DataTypeBoolean), string(DataTypeDate), string(DataTypeList), string(DataTypeObject):
		*dt = DataType(str)
		return nil
	default:
		return fmt.Errorf("invalid DataType value: %s", str)
	}
}

func (dt DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(dt))
}

func (dt DataType) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, dt)))
}

func (dt *DataType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("DataType must be a string")
	}
	switch str {
	case string(DataTypeString), string(DataTypeNumber), string(DataTypeBoolean), string(DataTypeDate), string(DataTypeList), string(DataTypeObject):
		*dt = DataType(str)
		return nil
	default:
		return fmt.Errorf("invalid DataType value: %s", str)
	}
}

// ConditionOrLogical is an interface implemented by Condition and LogicalGroup
type ConditionOrLogical interface {
	IsConditionOrLogical() // Marker interface method
}

// Condition Struct
type Condition struct {
	ID            string          `json:"id"`
	Field         string          `json:"field"`
	Operator      LogicalOperator `json:"operator"`
	Value         any             `json:"value"`
	Type          DataType        `json:"type"`
	CaseSensitive bool            `json:"caseSensitive"`
}

// LogicalGroup Struct
type LogicalGroup struct {
	ID         string               `json:"id"`
	Operator   Operator             `json:"operator"`
	Conditions []ConditionOrLogical `json:"conditions"`
}

// IsConditionOrLogical Implement IsConditionOrLogical for Condition
func (c Condition) IsConditionOrLogical() {}

// IsConditionOrLogical Implement IsConditionOrLogical for LogicalGroup
func (lg LogicalGroup) IsConditionOrLogical() {}
