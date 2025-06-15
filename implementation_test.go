package lab2

import (
	"fmt"
	"testing"
)

func TestPostfixToInfix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:     "Simple addition",
			input:    "2 3 +",
			expected: "(2 + 3)",
			hasError: false,
		},
		{
			name:     "Simple subtraction",
			input:    "5 2 -",
			expected: "(5 - 2)",
			hasError: false,
		},
		{
			name:     "Simple multiplication",
			input:    "4 3 *",
			expected: "(4 * 3)",
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PostfixToInfix(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %s, got %s", tt.expected, result)
				}
			}
		})
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 3 +")
	fmt.Println(res)
}
