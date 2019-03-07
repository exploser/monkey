package opcode

// Opcode is the type for all opcodes
type Opcode byte

//go:generate ../stringer -type=Opcode

// Opcode values
const (
	Const Opcode = iota
	Pop
	Add
	Sub
	Mul
	Div
	Not
	Neg
)
