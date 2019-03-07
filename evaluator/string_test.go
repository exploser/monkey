package evaluator

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/test"
)

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
			test.String(t, expected, evaluated, tt)
		default:
			test.Null(t, evaluated, tt)
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
			test.String(t, expected, evaluated, tt)
		default:
			test.Null(t, evaluated, tt)
		}
	}
}
