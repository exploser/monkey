package evaluator_test

import "testing"

func TestBoolean(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!(5-5)", true},
		{"!!10", true},
		{"!true", false},
		{"!!true", true},
		{"1 < 2", true},
		{"1 == 2", false},
		{"1 == (17 - 16)", true},
		{"2 != (17 - 16)", true},
		{"-1 == (16 - 17)", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, tt.expected, evaluated, tt)
	}
}
