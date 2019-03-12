package bytecode

import (
	"github.com/vasilevp/monkey/opcode"
)

type operandWidth []int

var OperandWidths [256]operandWidth = [256]operandWidth{
	opcode.Const: {2},
	opcode.Jmp:   {2},
	opcode.Jnt:   {2},
}
