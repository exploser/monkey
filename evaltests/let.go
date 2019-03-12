package evaltests

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func testLet(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let x = 5; x;", 5},
		{"let x = 5; let b = x * 2; b * 2;", 20},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		test.Integer(t, tt.expected, evaluated, tt)
	}
}
