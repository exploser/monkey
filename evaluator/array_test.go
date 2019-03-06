package evaluator

import "testing"

func TestArray(t *testing.T) {
	tests := []struct {
		input    string
		expected []int64
	}{
		{"[1, 2 * 2, 3 + 3]", []int64{1, 4, 6}},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		testIntegerArrayObject(t, tt.expected, evaluated, tt)
	}
}
