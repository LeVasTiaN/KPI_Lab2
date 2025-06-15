package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func PostfixToInfix(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("empty expression")
	}

	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return "", fmt.Errorf("empty expression")
	}

	var stack []string

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return "", fmt.Errorf("insufficient operands for operator %s", token)
			}

			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result := fmt.Sprintf("(%s %s %s)", operand1, token, operand2)
			stack = append(stack, result)
		default:
			if _, err := strconv.ParseFloat(token, 64); err != nil {
				return "", fmt.Errorf("invalid token: %s", token)
			}
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression: too many operands")
	}

	return stack[0], nil
}
