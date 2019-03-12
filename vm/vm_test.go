package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/monkey/bytecode"
	"git.exsdev.ru/ExS/monkey/evaltests"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"git.exsdev.ru/ExS/monkey/types"
	"git.exsdev.ru/ExS/monkey/vm"
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
