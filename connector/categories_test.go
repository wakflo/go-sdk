package sdk

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		cc   ConnectorGroup
		want bool
	}{
		{
			name: "Valid String",
			cc:   "scripts",
			want: true,
		},
		{
			name: "Invalid String",
			cc:   "invalid",
			want: false,
		},
		{
			name: "Test Case Sensitivity",
			cc:   "SCRIPTS",
			want: false,
		},
		{
			name: "Empty String",
			cc:   "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cc.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		cc      ConnectorGroup
		wantErr bool
	}{
		{
			name:    "Valid String",
			cc:      "ai",
			wantErr: false,
		},
		{
			name:    "Invalid String",
			cc:      "invalid",
			wantErr: true,
		},
		{
			name:    "Empty String",
			cc:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cc.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFromStringMethods(t *testing.T) {
	tests := []struct {
		name     string
		cc       string
		want     ConnectorGroup
		wantBool bool
	}{
		{
			name:     "Valid String",
			cc:       "core",
			want:     "core",
			wantBool: true,
		},
		{
			name:     "Invalid String",
			cc:       "invalid",
			want:     "apps",
			wantBool: false,
		},
		{
			name:     "Case Sensitive",
			cc:       "COre",
			want:     "apps",
			wantBool: false,
		},
		{
			name:     "Empty String",
			cc:       "",
			want:     "apps",
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotBool := ConnectorCategoryFromString(tt.cc)
			if got != tt.want || gotBool != tt.wantBool {
				t.Errorf("FromString() = %v, %v, want %v, %v", got, gotBool, tt.want, tt.wantBool)
			}
		})
	}
}
