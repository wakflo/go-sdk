package utils

import (
	"testing"
)

func TestGenerateUniqueSlug(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Blank Input",
			input:    "",
			expected: "",
		},
		{
			name:     "Basic String",
			input:    "Testing",
			expected: "testing",
		},
		{
			name:     "String with special characters",
			input:    "Te$ting N@me!",
			expected: "te-ting-natme",
		},
		{
			name:     "String with spaces",
			input:    "Testing Slug Name",
			expected: "testing-slug-name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateUniqueSlug(tt.input)
			if result != tt.expected {
				t.Errorf("GenerateUniqueSlug() got = %v, want = %v", result, tt.expected)
			}
		})
	}
}
