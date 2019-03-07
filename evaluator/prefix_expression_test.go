package evaluator

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

func TestBang(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!5", false},
		{"!!10", true},
		{"!true", false},
		{"!!true", true},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		test.Boolean(t, tt.expected, evaluated, tt)
	}
}
