package sdk

import (
	"github.com/juicycleff/smartform/v1"
	"github.com/wakflo/go-sdk/v2/context"
	"github.com/wakflo/go-sdk/v2/core"
)

type DynamicOptionsFn = func(ctx context.DynamicFieldContext) (*core.DynamicOptionsResponse, error)

func WithDynamicFunctionCalling(fn *DynamicOptionsFn) smartform.DynamicFunction {
	return func(args map[string]interface{}, formState map[string]interface{}) (interface{}, error) {
		if fn == nil {
			return nil, nil
		}

		ctxRaw, ok := args["ctx"]
		if !ok {
			return nil, nil
		}

		ctx, ok := ctxRaw.(context.DynamicFieldContext)
		if !ok {
			return nil, nil
		}

		return (*fn)(ctx)
	}
}

type FieldType string

const (
	FieldTypeCondition         FieldType = "condition"
	FieldTypeEnhancedCondition FieldType = "enhanced_condition"
	FieldTypeMap               FieldType = "map"
	FieldTypeKeyValue          FieldType = "keyvalue"
	FieldTypeKeyCode           FieldType = "code"
	FieldTypeKeyIDECode        FieldType = "ide_code"
	FieldTypeBranch            FieldType = "branch"
	FieldTypeRouter            FieldType = "router"
)

func (t FieldType) String() string { return string(t) }
