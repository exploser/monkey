package evaluator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vasilevp/monkey/types"
)

func TestFunction(t *testing.T) {
	input := "fn(x) { return x + 2;}"

	evaluated := testEval(t, input)
	require.IsType(t, new(types.Function), evaluated)
	fn := evaluated.(*types.Function)
	require.Len(t, fn.Parameters, 1)
	require.Equal(t, "x", fn.Parameters[0].String())
	require.Equal(t, "return (x + 2); ", fn.Body.String())
}

func TestFunctionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let identity = fn (x) { x; }; identity(10)", 10},
		{"let double = fn (x) { x*2; }; double(5)", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(t, tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, int64(expected), evaluated, tt)
		default:
			testNullObject(t, evaluated, tt)
		}
	}
}
