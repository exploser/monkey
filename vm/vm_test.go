package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vasilevp/monkey/bytecode"
	"github.com/vasilevp/monkey/evaltests"
	"github.com/vasilevp/monkey/lexer"
	"github.com/vasilevp/monkey/parser"
	"github.com/vasilevp/monkey/types"
	"github.com/vasilevp/monkey/vm"
)

func e(t testing.TB, input string) types.Object {
	t.Helper()

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	t.Logf("%+v", program)

	compiler := bytecode.New()
	err := compiler.Compile(program)
	require.Nil(t, err)

	vm := vm.New()
	err = vm.Run(compiler.Bytecode)
	require.Nil(t, err)

	return vm.LastPopped()
}

func TestEverything(t *testing.T) {
	evaltests.Run(t, e, []string{
		// "array",
		"bang",
		"boolean",
		// "declare",
		// "error",
		"evalIntegerExpression",
		// "function",
		// "functionCall",
		// "if",
		"infix",
		// "len",
		// "let",
		// "return",
		// "string",
		// "stringConcat",
	})
}

func BenchmarkEverything(b *testing.B) {
	evaltests.BenchmarkAll(b, e)
}
