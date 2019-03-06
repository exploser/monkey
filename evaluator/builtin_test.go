package evaluator

import "testing"

func TestLen(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`len("")`, 0},
		{`len("hello")`, 5},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		testIntegerObject(t, tt.expected, evaluated, tt)
	}
}
