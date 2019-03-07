package evaluator

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func TestDeclare(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"x := 5; x;", 5},
		{"x := 5; b := x * 2; b*2", 20},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		test.Integer(t, tt.expected, evaluated, tt)
	}
}
