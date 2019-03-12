package evaltests

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func testLen(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`len("")`, 0},
		{`len("hello")`, 5},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		test.Integer(t, tt.expected, evaluated, tt)
	}
}
