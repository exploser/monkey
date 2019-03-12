package evaltests

import (
	"testing"

	"github.com/vasilevp/monkey/test"
)

func testBoolean(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!(5-5)", true},
		{"!!10", true},
		{"!true", false},
		{"!!true", true},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 == 2", false},
		{"1 == (17 - 16)", true},
		{"2 != (17 - 16)", true},
		{"-1 == (16 - 17)", true},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		test.Boolean(t, tt.expected, evaluated, tt)
	}
}
