package sdk

import (
	"github.com/wakflo/go-sdk/v2/context"
	"github.com/wakflo/go-sdk/v2/core"
)

type SmartFormDynFieldFn = func(args map[string]interface{}, formState map[string]interface{}) (interface{}, error)

type DynamicOptionsFn = func(ctx context.DynamicFieldContext) (*core.DynamicOptionsResponse, error)

func WithDynamicFunctionCalling(fn *DynamicOptionsFn) SmartFormDynFieldFn {
	return func(args map[string]interface{}, formState map[string]interface{}) (interface{}, error) {
		if fn == nil {
			return nil, nil
		}

		ctxRaw, ok := args["context"]
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
