package evaltests

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func testArray(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected []int64
	}{
		{"[1, 2 * 2, 3 + 3]", []int64{1, 4, 6}},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		test.IntegerArray(t, tt.expected, evaluated, tt)
	}
}
