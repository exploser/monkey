package bytecode

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/vasilevp/monkey/opcode"
)

type Instructions []byte

func (i Instructions) String() string {
	p := i
	result := strings.Builder{}

	for len(p) > 0 {
		o, operands, n := p.ReadOpcode()
		result.WriteString(fmt.Sprintf("%04d %v %v\n", len(i)-len(p), o, operands))
		p = p[n:]
	}

	return result.String()
}

func (i Instructions) ReadOpcode() (opcode.Opcode, []int, int) {
	o := i[0]
	widths := OperandWidths[o]
	offset := 1
	operands := make([]int, len(widths))
	for k, width := range widths {
		switch width {
		case 2:
			operands[k] = int(binary.BigEndian.Uint16(i[offset:]))
		case 4:
			operands[k] = int(binary.BigEndian.Uint32(i[offset:]))
		}

		offset += width
	}

	return opcode.Opcode(o), operands, offset
}
