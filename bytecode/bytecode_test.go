package bytecode

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/opcode"
	"github.com/stretchr/testify/require"
)

func TestMake(t *testing.T) {
	tests := []struct {
		o        opcode.Opcode
		operands []int
		expected Instructions
	}{
		{opcode.Const, []int{65534}, Instructions{byte(opcode.Const), 255, 254}},
	}

	for _, tc := range tests {
		instruction := Make(tc.o, tc.operands...)
		require.Equal(t, tc.expected, instruction)
	}
}
