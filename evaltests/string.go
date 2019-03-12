package evaltests

import (
	"testing"

	"github.com/vasilevp/monkey/test"
)

func testString(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`"hello"`, "hello"},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		switch expected := tt.expected.(type) {
		case string:
			test.String(t, expected, evaluated, tt)
		default:
			test.Null(t, evaluated, tt)
		}
	}
}

func testStringConcat(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`"hello" + " " + "world!"`, "hello world!"},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		switch expected := tt.expected.(type) {
		case string:
			test.String(t, expected, evaluated, tt)
		default:
			test.Null(t, evaluated, tt)
		}
	}
}
