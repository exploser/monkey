package evaltests

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func testReturn(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"return 10;", 10},
		{"return 10; 2;", 10},
		{"if (true) { if (true) { return 10; } } return 1;", 10},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		switch expected := tt.expected.(type) {
		case int:
			test.Integer(t, int64(expected), evaluated, tt)
		default:
			test.Null(t, evaluated, tt)
		}
	}
}
