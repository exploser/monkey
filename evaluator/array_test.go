package evaluator

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func TestArray(t *testing.T) {
	tests := []struct {
		input    string
		expected []int64
	}{
		{"[1, 2 * 2, 3 + 3]", []int64{1, 4, 6}},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		test.IntegerArray(t, tt.expected, evaluated, tt)
	}
}
