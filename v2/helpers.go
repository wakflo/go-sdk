package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	"github.com/juicycleff/smartform/v1"
	"github.com/pelletier/go-toml/v2"
	"github.com/wakflo/go-sdk/autoform"
	sdkcore "github.com/wakflo/go-sdk/core"
	sdkcontext "github.com/wakflo/go-sdk/v2/context"
	"github.com/wakflo/go-sdk/v2/core"
	"github.com/wakflo/go-sdk/validator"
)

// InputToType returns a pointer to a value of type T by marshaling and unmarshaling the ResolvedInput field of the provided RunContext struct.
// If there is an error during the marshaling or unmarshaling process, nil is returned.
// The function signature is as follows:
func InputToType[T any](ctx sdkcontext.BaseContext) *T {
	b, err := json.Marshal(ctx.Input())
	if err != nil {
		return nil
	}

	var rsp T
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil
	}

	return &rsp
}

// InputToTypeSafely returns a pointer to a value of type T by marshaling and unmarshaling the ResolvedInput field of the provided RunContext struct.
// If there is an error during the marshaling or unmarshaling process, nil is returned.
// The function signature is as follows:
func InputToTypeSafely[T any](ctx sdkcontext.BaseContext) (*T, error) {
	b, err := json.Marshal(ctx.Input())
	if err != nil {
		return nil, err
	}

	var rsp T
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}

// InputPropsToType returns a pointer to a value of type T by marshaling and unmarshaling the ResolvedInput field of the provided RunContext struct.
// If there is an error during the marshaling or unmarshaling process, nil is returned.
// The function signature is as follows:
func InputPropsToType[T any](input sdkcore.JSON) (*T, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var rsp T
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}

// DynamicInputToType converts the resolved input of type `sdkcore.DynamicOptionsContext` to the desired type T.
// It uses JSON marshaling and unmarshalling to perform the conversion.
// If any error occurs during marshaling or unmarshaling, it returns nil.
// The function returns a pointer to the converted value of type T.
func DynamicInputToType[T any](ctx sdkcontext.BaseContext) *T {
	b, err := json.Marshal(ctx.Input())
	if err != nil {
		return nil
	}

	var rsp T
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return nil
	}

	return &rsp
}

// StringToFile converts a file string to a *autoform.File object.
//
// The function checks if the file string is a base64-encoded data or a URL. If the file string is base64-encoded data, it decodes the data and assigns it to the `Data` field of the
func StringToFile(fileStr string) (*autoform.File, error) {
	// file := &autoform.File{}
	//
	// if valid.IsBase64(fileStr) {
	// 	data, err := base64.StdEncoding.DecodeString(fileStr)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	mime := mimetype.Detect(data)
	// 	file.Data = bytes.NewReader(data)
	// 	file.Extension = mime.Extension()
	// 	file.Mime = mime.String()
	//
	// 	return file, nil
	// }
	//
	// if valid.IsURL(fileStr) {
	// 	data, err := DownloadFile(fileStr)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	bt, err := data.Bytes()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	mime := mimetype.Detect(bt)
	// 	file.Data = bytes.NewReader(bt)
	// 	file.Extension = mime.Extension()
	// 	file.Size = data.Size()
	// 	file.Name = data.Filename
	// 	file.Mime = mime.String()
	//
	// 	return file, nil
	// }

	return nil, errors.New("invalid file string")
}

// DownloadFile downloads a file from the specified URL using the grab package.
// It returns the grab.Response object and an error if any.
func DownloadFile(url string) (*grab.Response, error) {
	resp, err := grab.Get(".", url)
	if err != nil {
		return nil, err
	}

	resp.Wait()

	return resp, nil
}

type RegistrationMap struct {
	Versions map[string]Integration
}

type IntegrationsRegistrar = map[string]RegistrationMap

func Register(integration Integration) Integration {
	err := validator.NewDefaultValidator().Validate(integration.Metadata())
	if err != nil {
		log.Panicf("invalid integration: %s", err)
	}
	return integration
}

func registerIntegration(flow string, readme string) IntegrationMetadata {
	reg, err := registerInternal(flow, readme)
	if err != nil {
		panic(err)
	}
	return *reg
}

func registerInternal(flow string, readme string) (*IntegrationMetadata, error) {
	readme, err := ReadREADME(readme)
	if err != nil {
		return nil, fmt.Errorf("failed to read README.md: %w", err)
	}

	info, err := ReadFloFile(flow)
	if err != nil {
		return nil, err
	}

	// name := slug.Make(info.Name)
	// info.Name = name
	info.Documentation = readme

	return info, nil
}

type IntegrationSchemaModel struct {
	Name        string   `json:"name" toml:"name" yaml:"name" validate:"required"`
	Description string   `json:"description" toml:"description"  yaml:"description" validate:"required"`
	Version     string   `json:"version" toml:"version" yaml:"version" validate:"required"`
	Authors     []string `json:"authors" toml:"authors" yaml:"authors" validate:"required"`
	Website     *string  `json:"website" toml:"website" yaml:"website"`
	Categories  []string `json:"categories" toml:"categories" yaml:"categories" validate:"required"`
	Icon        string   `json:"icon" toml:"icon" yaml:"icon" validate:"required"`
}

type SchemaConfig struct {
	Integration IntegrationMetadata `json:"integration" toml:"integration" yaml:"integration" validate:"required"`
}

func ReadFloFile(content string) (*IntegrationMetadata, error) {
	// Deserialize the file content into SchemaConfig
	var config SchemaConfig
	err := toml.Unmarshal([]byte(content), &config)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize TOML: %w", err)
	}

	v := validator.NewDefaultValidator()

	err = v.Validate(&config)
	if err != nil {
		return nil, err
	}

	return &config.Integration, nil
}

// ReadREADME extracts the content of README.md from the current directory.
func ReadREADME(content string) (string, error) {
	// fileName := "README.md"
	// // Resolve the full path relative to the caller's directory
	// fullPath, err := GetRelativePathWithDepth(fileName, 4)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to resolve full path: %w", err)
	// }
	//
	// // Define the path to README.md
	// readmePath := filepath.Join(fullPath)
	//
	// // Check if the file exists
	// if _, err := os.Stat(readmePath); os.IsNotExist(err) {
	// 	return "", fmt.Errorf("README.md not found in the directory: %s", fullPath)
	// }
	//
	// // Read the file content
	// content, err := os.ReadFile(readmePath)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to read README.md: %w", err)
	// }

	// Trim any leading or trailing whitespace
	return strings.TrimSpace(string(content)), nil
}

// GetRelativePathWithDepth resolves a file path relative to the location of the caller, with adjustable depth.
func GetRelativePathWithDepth(relativePath string, depth int) (string, error) {
	// Get the caller's frame based on the specified depth
	_, callerFile, _, ok := runtime.Caller(depth)
	if !ok {
		return "", fmt.Errorf("unable to retrieve caller information")
	}

	// Get the directory of the caller
	callerDir := filepath.Dir(callerFile)

	// Resolve the full path relative to the caller's directory
	fullPath := filepath.Join(callerDir, relativePath)
	return fullPath, nil
}

// DynamicOptionsFn defines a function type that processes a DynamicFieldContext and returns a DynamicOptionsResponse or an error.
type DynamicOptionsFn = func(ctx sdkcontext.DynamicFieldContext) (*core.DynamicOptionsResponse, error)

// WithDynamicFunctionCalling wraps a DynamicOptionsFn into a smartform.DynamicFunction to execute dynamic field actions.
// If the provided function is nil or required arguments are missing in the call context, it returns nil.
func WithDynamicFunctionCalling(fn *DynamicOptionsFn) smartform.DynamicFunction {
	return func(args map[string]interface{}, formState map[string]interface{}) (interface{}, error) {
		if fn == nil {
			return nil, nil
		}

		ctxRaw, ok := args["ctx"]
		if !ok {
			return nil, nil
		}

		ctx, ok := ctxRaw.(sdkcontext.DynamicFieldContext)
		if !ok {
			return nil, nil
		}

		return (*fn)(ctx)
	}
}

// FieldType represents a string-based type used for categorizing various kinds of fields in the system.
type FieldType string

// FieldTypeCondition represents a condition type field.
// FieldTypeEnhancedCondition represents an enhanced condition type field.
// FieldTypeMap represents a map type field.
// FieldTypeKeyValue represents a key-value type field.
// FieldTypeKeyCode represents a code type field.
// FieldTypeKeyIDECode represents an IDE code type field.
// FieldTypeBranch represents a branch type field.
// FieldTypeRouter represents a router type field.
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

// String converts the FieldType value to its string representation.
func (t FieldType) String() string { return string(t) }
