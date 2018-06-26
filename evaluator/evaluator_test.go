package evaluator_test

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/evaluator"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"git.exsdev.ru/ExS/monkey/types"
	"github.com/stretchr/testify/require"
)

func testEval(input string) types.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := evaluator.GetBaseEnvironment()

	return evaluator.Eval(program, env)
}

func testBooleanObject(t *testing.T, expected bool, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Boolean), obj, "tc: %s, result: %s", context, obj)
	result := obj.(*types.Boolean)
	require.Equal(t, expected, result.Value, "tc: %s, result: %s", context, obj)
}

func testIntegerObject(t *testing.T, expected int64, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Integer), obj, "tc: %s, result: %s", context, obj)
	result := obj.(*types.Integer)
	require.Equal(t, expected, result.Value, "tc: %s, result: %s", context, obj)
}

func testNullObject(t *testing.T, obj types.Object, context interface{}) {
	require.Equal(t, evaluator.NilValue, obj, "tc: %s, result: %s", context, obj)
}

func testError(t *testing.T, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Error), obj, "tc: %s, result: %s", context, obj)
}

func testStringObject(t *testing.T, expected string, obj types.Object, context interface{}) {
	require.IsType(t, new(types.String), obj, "tc: %s, result: %s", context, obj)
	result := obj.(*types.String)
	require.Equal(t, expected, result.Value, "tc: %s, result: %s", context, obj)
}
