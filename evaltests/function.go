package evaltests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/monkey/test"
	"git.exsdev.ru/ExS/monkey/types"
)

func testFunction(t *testing.T, e Evaluator) {
	input := "fn(x) { return x + 2;}"

	evaluated := e(t, input)
	require.IsType(t, new(types.Function), evaluated)
	fn := evaluated.(*types.Function)
	require.Len(t, fn.Parameters, 1)
	require.Equal(t, "x", fn.Parameters[0].String())
	require.Equal(t, "return (x + 2); ", fn.Body.String())
}

func testFunctionCall(t *testing.T, e Evaluator) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let identity = fn (x) { x; }; identity(10)", 10},
		{"let double = fn (x) { x*2; }; double(5)", 10},
	}

	for _, tt := range tests {
		evaluated := e(t, tt.input)
		switch expected := tt.expected.(type) {
		case int:
			test.Integer(t, int64(expected), evaluated, tt)
		default:
			test.Null(t, evaluated, tt)
		}
	}
}
