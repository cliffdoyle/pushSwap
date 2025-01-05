package validation

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []int
		hasError bool
	}{
		{
			name:     "Valid input with unique numbers",
			input:    []string{"1", "2", "3", "4", "5"},
			expected: []int{1, 2, 3, 4, 5},
			hasError: false,
		},
		{
			name:     "Invalid number in input",
			input:    []string{"1", "2", "three", "4"},
			expected: nil,
			hasError: true,
		},
		{
			name:     "Duplicate numbers in input",
			input:    []string{"1", "2", "3", "2"},
			expected: nil,
			hasError: true,
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: []int{},
			hasError: false,
		},
		{
			name:     "Single valid number",
			input:    []string{"10"},
			expected: []int{10},
			hasError: false,
		},
		{
			name:     "Negative numbers",
			input:    []string{"-1", "-2", "0", "3"},
			expected: []int{-1, -2, 0, 3},
			hasError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseInput(tc.input)

			// Check for expected errors
			if (err != nil) != tc.hasError {
				t.Errorf("unexpected error state: got %v, want error? %v", err, tc.hasError)
			}

			// Check for expected results only if no error is expected
			if !tc.hasError && !equalSlices(result, tc.expected) {
				t.Errorf("unexpected result: got %v, want %v", result, tc.expected)
			}
		})
	}
}

// Helper function to compare two slices of integers
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

