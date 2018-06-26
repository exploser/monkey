package evaluator_test

import (
	"fmt"
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
	env := types.NewEnvironment()

	return evaluator.Eval(program, env)
}

func testBooleanObject(t *testing.T, expected bool, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Boolean), obj, fmt.Sprint(context))
	result := obj.(*types.Boolean)
	require.Equal(t, expected, result.Value, fmt.Sprint(context))
}

func testIntegerObject(t *testing.T, expected int64, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Integer), obj, fmt.Sprint(context))
	result := obj.(*types.Integer)
	require.Equal(t, expected, result.Value, fmt.Sprint(context))
}

func testNullObject(t *testing.T, obj types.Object, context interface{}) {
	require.Equal(t, evaluator.NilValue, obj, fmt.Sprint(context))
}

func testError(t *testing.T, obj types.Object, context interface{}) {
	require.IsType(t, new(types.Error), obj, fmt.Sprint(context))
}
