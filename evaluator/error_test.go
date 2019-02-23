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
		{`"asd" - "d"`},
		{"[5+nil]"},
		{`-("a"+nil)`},
		{`(0+"a")+("b"+0)`},
		{`(0+1)+("b"+0)`},
		{`let a = "a"+0`},
		{`a := "a"+0`},
		{`a(0)`},
		{`fn(x) { return x }(a)`},
		{`fn(x) { return x }()`},
		{`7(0)`},
		{`-"a"`},
		{"if (5+nil) { 7 }"},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		testError(t, evaluated, tt)
	}
}
