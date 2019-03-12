package evaltests

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func testInfix(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1-27+16", -10},
		{"1-1 * 10+10 * 1", 1},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		test.Integer(t, tt.expected, evaluated, tt)
	}
}
