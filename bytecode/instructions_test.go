package bytecode_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "git.exsdev.ru/ExS/monkey/bytecode"
	"git.exsdev.ru/ExS/monkey/opcode"
)

func TestInstructionsString(t *testing.T) {
	instructions := i(
		Make(opcode.Const, 0),
		Make(opcode.Const, 1),
		Make(opcode.Const, 65534),
	)

	expected := `0000 Const [0]
0003 Const [1]
0006 Const [65534]
`

	require.Equal(t, expected, instructions.String())
}

func TestReadOpcode(t *testing.T) {
	tests := []struct {
		op        opcode.Opcode
		operands  []int
		bytesRead int
	}{
		{opcode.Const, []int{65534}, 3},
	}

	for _, tc := range tests {
		instruction := Make(tc.op, tc.operands...)
		opcode, operands, n := instruction.ReadOpcode()
		require.Equal(t, tc.op, opcode)
		require.Equal(t, tc.operands, operands)
		require.Equal(t, tc.bytesRead, n)
	}
}
