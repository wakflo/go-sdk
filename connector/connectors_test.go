package sdk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Let's create a dummy implementation of IConnectorPlugin to simulate expected behavior
type MockPlugin struct{}

func (mp *MockPlugin) Authenticate(ctx context.Context) {}

func (mp *MockPlugin) Test(ctx context.Context) {}

func TestConnectorPlugin_Authenticate(t *testing.T) {
	connectorPlugin := ConnectorPlugin{}
	ctx := context.Background()

	connectorPlugin.Authenticate(ctx) // In this case, function does not return anything and does not alter any state
}

func TestConnectorPlugin_Test(t *testing.T) {
	connectorPlugin := ConnectorPlugin{}
	ctx := context.Background()

	connectorPlugin.Test(ctx) // In this case, function does not return anything and does not alter any state
}

func TestNewConnectorPlugin(t *testing.T) {
	// create initial data
	validArgs := CreateConnectorArgs{
		Name:        "validname",
		Description: "validdescription",
		Version:     "validversion",
	}

	invalidArgs := CreateConnectorArgs{
		// Let's leave Name as empty string.
		Description: "validdescription",
		Version:     "validversion",
	}

	tests := map[string]struct {
		inputArgs *CreateConnectorArgs
		wantErr   bool
	}{
		"valid inputs":   {inputArgs: &validArgs, wantErr: false},
		"invalid inputs": {inputArgs: &invalidArgs, wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewConnectorPlugin(tc.inputArgs)
			if (err != nil) != tc.wantErr {
				t.Fatalf("NewConnectorPlugin() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if tc.wantErr == false && got == nil {
				t.Errorf("NewConnectorPlugin() = %v, want non-nil", got)
			}
		})
	}
}

func TestValidateConnectorInfo(t *testing.T) {
	// create initial data
	validInfo := CreateConnectorArgs{
		Name:        "validname",
		Description: "validdescription",
		Version:     "validversion",
	}

	invalidInfo := CreateConnectorArgs{
		// Leave Name as empty string.
		Description: "validdescription",
		Version:     "validversion",
	}

	tests := map[string]struct {
		inputInfo *CreateConnectorArgs
		wantErr   bool
	}{
		"valid inputs":   {inputInfo: &validInfo, wantErr: false},
		"invalid inputs": {inputInfo: &invalidInfo, wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := ValidateConnectorInfo(tc.inputInfo)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ValidateConnectorInfo() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestValidateOperationInfo(t *testing.T) {
	t.Run("info is nil", func(t *testing.T) {
		err := ValidateOperationInfo(nil)
		require.Error(t, err)
		assert.Equal(t, "operation info missing", err.Error())
	})

	t.Run("info name is missing", func(t *testing.T) {
		info := &OperationInfo{
			Description: "Some description",
		}

		err := ValidateOperationInfo(info)
		require.Error(t, err)
		assert.Equal(t, "operation info name missing", err.Error())
	})

	t.Run("info description is missing", func(t *testing.T) {
		info := &OperationInfo{
			Name: "Some name",
		}

		err := ValidateOperationInfo(info)
		require.Error(t, err)
		assert.Equal(t, "operation info description missing", err.Error())
	})

	t.Run("info is valid", func(t *testing.T) {
		info := &OperationInfo{
			Name:        "Some name",
			Description: "Some description",
		}

		err := ValidateOperationInfo(info)
		require.NoError(t, err)
	})
}
