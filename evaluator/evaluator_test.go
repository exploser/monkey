package evaluator_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/monkey/evaluator"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"git.exsdev.ru/ExS/monkey/types"
)

func testEval(input string) types.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := evaluator.GetBaseEnvironment()

	return evaluator.Eval(program, env)
}

func testBooleanObject(t *testing.T, expected bool, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Boolean), obj, "tc: %v, result: %v", context, obj)
	result := obj.(*types.Boolean)
	require.Equal(t, expected, result.Value, "tc: %v, result: %v", context, obj)
}

func testIntegerObject(t *testing.T, expected int64, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Integer), obj, "tc: %v, result: %v", context, obj)
	result := obj.(*types.Integer)
	require.Equal(t, expected, result.Value, "tc: %v, result: %v", context, obj)
}

func testNullObject(t *testing.T, obj types.Object, context interface{}) {
	require.Equal(t, evaluator.NilValue, obj, "tc: %v, result: %v", context, obj)
}

func testError(t *testing.T, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Error), obj, "tc: %v, result: %v", context, obj)
}

func testStringObject(t *testing.T, expected string, obj types.Object, context interface{}) {
	require.IsType(t, new(types.String), obj, "tc: %v, result: %v", context, obj)
	result := obj.(*types.String)
	require.Equal(t, expected, result.Value, "tc: %v, result: %v", context, obj)
}

func testIntegerArrayObject(t *testing.T, expected []int64, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Array), obj, "tc: %v, result: %v", context, obj)

	array := obj.(*types.Array)

	for k, v := range array.Elements {
		require.IsType(t, new(types.Integer), v, "tc: %v, result: %v", context, v)
		result := v.(*types.Integer)
		require.Equal(t, expected[k], result.Value, "tc: %v, result: %v", context, v)
	}
}
