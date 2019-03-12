package evaltests

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func testEvalIntegerExpression(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"50 / 10", 5},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		test.Integer(t, tt.expected, evaluated, tt)
	}
}
