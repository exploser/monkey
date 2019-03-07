package bytecode

import (
	"encoding/binary"

	"git.exsdev.ru/ExS/monkey/opcode"
	"git.exsdev.ru/ExS/monkey/types"
)

type Bytecode struct {
	Constants []types.Object
	Code      Instructions
}

// X bytes per instruction should be enough for everyone
const maxBufferSize = 64

func Make(op opcode.Opcode, operands ...int) Instructions {
	instruction := [maxBufferSize]byte{byte(op)}
	widths := OperandWidths[op]
	off := 1

	for k, v := range operands {
		w := widths[k]

		switch widths[k] {
		case 2:
			binary.BigEndian.PutUint16(instruction[off:], uint16(v))
		case 4:
			binary.BigEndian.PutUint32(instruction[off:], uint32(v))
		}

		off += w
	}

	return instruction[:off]
}
