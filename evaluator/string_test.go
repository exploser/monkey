package evaluator

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`"hello"`, "hello"},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		switch expected := tt.expected.(type) {
		case string:
			testStringObject(t, expected, evaluated, tt)
		default:
			testNullObject(t, evaluated, tt)
		}
	}
}

func TestStringConcat(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`"hello" + " " + "world!"`, "hello world!"},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		switch expected := tt.expected.(type) {
		case string:
			testStringObject(t, expected, evaluated, tt)
		default:
			testNullObject(t, evaluated, tt)
		}
	}
}
