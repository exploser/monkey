package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/monkey/globals"
	"git.exsdev.ru/ExS/monkey/types"
)

func Boolean(t testing.TB, expected bool, obj types.Object, context interface{}) {
	t.Helper()

	require.IsType(t, new(types.Boolean), obj, "tc: %+v, result: %v", context, obj)
	result := obj.(*types.Boolean)
	require.Equal(t, expected, result.Value, "tc: %+v, result: %v", context, obj)
}

func Integer(t testing.TB, expected int64, obj types.Object, context interface{}) {
	t.Helper()

	require.IsType(t, new(types.Integer), obj, "tc: %+v, result: %v", context, obj)
	result := obj.(*types.Integer)
	require.Equal(t, expected, result.Value, "tc: %+v, result: %v", context, obj)
}

func Null(t testing.TB, obj types.Object, context interface{}) {
	t.Helper()

	require.Equal(t, globals.Nil, obj, "tc: %+v, result: %v", context, obj)
}

func Error(t testing.TB, obj types.Object, context interface{}) {
	t.Helper()

	require.IsType(t, new(types.Error), obj, "tc: %+v, result: %v", context, obj)
}

func String(t testing.TB, expected string, obj types.Object, context interface{}) {
	t.Helper()

	require.IsType(t, new(types.String), obj, "tc: %+v, result: %v", context, obj)
	result := obj.(*types.String)
	require.Equal(t, expected, result.Value, "tc: %+v, result: %v", context, obj)
}

func IntegerArray(t testing.TB, expected []int64, obj types.Object, context interface{}) {
	t.Helper()

	require.IsType(t, new(types.Array), obj, "tc: %+v, result: %v", context, obj)

	array := obj.(*types.Array)

	for k, v := range array.Elements {
		require.IsType(t, new(types.Integer), v, "tc: %+v, result: %v", context, v)
		result := v.(*types.Integer)
		require.Equal(t, expected[k], result.Value, "tc: %+v, result: %v", context, v)
	}
}
