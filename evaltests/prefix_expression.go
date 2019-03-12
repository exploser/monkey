package evaltests

import (
	"testing"

	"github.com/vasilevp/monkey/test"
)

func testBang(t *testing.T, e Evaluator) {
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
		evaluated := e(t, tt.input)
		test.Boolean(t, tt.expected, evaluated, tt)
	}
}
