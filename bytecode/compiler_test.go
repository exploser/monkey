package bytecode_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "git.exsdev.ru/ExS/monkey/bytecode"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/opcode"
	"git.exsdev.ru/ExS/monkey/parser"
	"git.exsdev.ru/ExS/monkey/test"
	"git.exsdev.ru/ExS/monkey/types"
)

type compilerTest struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tt := []compilerTest{
		{
			"1 + 2",
			c(1, 2),
			i(Make(opcode.Const, 0), Make(opcode.Const, 1), Make(opcode.Add), Make(opcode.Pop)),
		},
		{
			"1 - 2 * 3",
			c(1, 2, 3),
			i(Make(opcode.Const, 0), Make(opcode.Const, 1), Make(opcode.Const, 2), Make(opcode.Mul), Make(opcode.Sub), Make(opcode.Pop)),
		},
		{
			"1 * 211",
			c(1, 211),
			i(Make(opcode.Const, 0), Make(opcode.Const, 1), Make(opcode.Mul), Make(opcode.Pop)),
		},
		{
			"1 / 2",
			c(1, 2),
			i(Make(opcode.Const, 0), Make(opcode.Const, 1), Make(opcode.Div), Make(opcode.Pop)),
		},
	}
	runCompilerTests(t, tt)
}

func runCompilerTests(t *testing.T, tests []compilerTest) {
	t.Helper()

	for _, tc := range tests {
		l := lexer.New(tc.input)
		p := parser.New(l)
		c := New()
		prog := p.ParseProgram()

		err := c.Compile(prog)
		require.Nil(t, err)

		require.Equal(t, tc.expectedInstructions, c.Code)
		testConstants(t, tc.expectedConstants, c.Constants, tc)
	}
}

func testConstants(t *testing.T, expected []interface{}, actual []types.Object, tc interface{}) {
	t.Helper()

	require.Len(t, actual, len(expected))

	for k, v := range expected {
		switch v := v.(type) {
		case int:
			test.Integer(t, int64(v), actual[k], tc)
		}
	}
}

func c(values ...interface{}) []interface{} {
	return values
}

func i(values ...Instructions) Instructions {
	result := Instructions{}
	for _, ii := range values {
		result = append(result, ii...)
	}
	return result
}
