package sdk

import (
	"encoding/json"
	"errors"

	"github.com/cavaliergopher/grab/v3"
	autoform "github.com/wakflo/go-sdk/autoform"
	sdkcore "github.com/wakflo/go-sdk/core"
	sdkcontext "github.com/wakflo/go-sdk/v2/context"
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
