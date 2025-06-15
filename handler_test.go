package lab2

import (
	"bytes"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "Simple postfix expression",
			input:          "2 3 +",
			expectedOutput: "(2 + 3)",
			expectError:    false,
		},
		{
			name:           "Complex postfix expression",
			input:          "4 2 - 3 2 ^ * 5 +",
			expectedOutput: "(((4 - 2) * (3 ^ 2)) + 5)",
			expectError:    false,
		},
		{
			name:           "Single number",
			input:          "42",
			expectedOutput: "42",
			expectError:    false,
		},
		{
			name:        "Empty input",
			input:       "",
			expectError: true,
		},
		{
			name:        "Invalid expression - insufficient operands",
			input:       "2 +",
			expectError: true,
		},
		{
			name:        "Invalid expression - too many operands",
			input:       "2 3 4 +",
			expectError: true,
		},
		{
			name:        "Invalid token",
			input:       "2 abc +",
			expectError: true,
		},
		{
			name:           "Expression with whitespace",
			input:          "  2 3 +  ",
			expectedOutput: "(2 + 3)",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)

			var output bytes.Buffer

			handler := &ComputeHandler{
				Input:  input,
				Output: &output,
			}

			err := handler.Compute()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				actualOutput := output.String()
				if actualOutput != tt.expectedOutput {
					t.Errorf("Expected output %q, got %q", tt.expectedOutput, actualOutput)
				}
			}
		})
	}
}
