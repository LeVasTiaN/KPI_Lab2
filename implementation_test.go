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
		{
			name:     "Simple division",
			input:    "8 2 /",
			expected: "(8 / 2)",
			hasError: false,
		},
		{
			name:     "Simple power",
			input:    "2 3 ^",
			expected: "(2 ^ 3)",
			hasError: false,
		},
		{
			name:     "Three operand expression",
			input:    "2 3 + 4 *",
			expected: "((2 + 3) * 4)",
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

func TestPostfixToInfixComplex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:     "Complex expression 1",
			input:    "4 2 - 3 2 ^ * 5 +",
			expected: "(((4 - 2) * (3 ^ 2)) + 5)",
			hasError: false,
		},
		{
			name:     "Complex expression 2",
			input:    "15 7 1 1 + - / 3 * 2 1 1 + + -",
			expected: "(((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1)))",
			hasError: false,
		},
		{
			name:     "Many operations",
			input:    "1 2 + 3 4 + * 5 6 + 7 8 + * +",
			expected: "(((1 + 2) * (3 + 4)) + ((5 + 6) * (7 + 8)))",
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

func TestPostfixToInfixErrors(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		hasError bool
	}{
		{
			name:     "Empty string",
			input:    "",
			hasError: true,
		},
		{
			name:     "Whitespace only",
			input:    "   ",
			hasError: true,
		},
		{
			name:     "Invalid character",
			input:    "2 3 #",
			hasError: true,
		},
		{
			name:     "Insufficient operands",
			input:    "2 +",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := PostfixToInfix(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 3 +")
	fmt.Println(res)
	// Output:
	// 5
}
