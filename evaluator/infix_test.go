package evaluator

import (
	"testing"
)

func TestInfix(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1-27+16", -10},
		{"1-1 * 10+10 * 1", 1},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		testIntegerObject(t, tt.expected, evaluated, tt)
	}
}
