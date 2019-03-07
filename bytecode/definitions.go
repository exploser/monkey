package bytecode

import (
	"git.exsdev.ru/ExS/monkey/opcode"
)

type operandWidth []int

var OperandWidths [256]operandWidth = [256]operandWidth{
	opcode.Const: {2},
}
