package evaluator

import (
	"testing"
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
		testBooleanObject(t, tt.expected, evaluated, tt)
	}
}
