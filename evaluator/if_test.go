package evaluator

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func TestIf(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1<2) { 10 }", 10},
		{"if (nil) { 10 }", nil},
		{"if (nil) { 10 } else {5}", 5},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		switch expected := tt.expected.(type) {
		case int:
			test.Integer(t, int64(expected), evaluated, tt)
		default:
			test.Null(t, evaluated, tt)
		}
	}
}
