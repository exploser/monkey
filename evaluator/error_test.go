package evaluator_test

import "testing"

func TestError(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"5 + true;"},
		{"5 + true; 5"},
		{"if (true) { if (true) { return 10+true; } } return 1;"},
		{"asdfg"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testError(t, evaluated, tt)
	}
}
