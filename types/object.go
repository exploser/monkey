package types

//go:generate stringer -type=ObjectType
type ObjectType int

const (
	NilT ObjectType = iota
	IntegerT
	ErrorT
	FunctionT
	ReturnT
	StringT
	ArrayT
	BooleanT
	BuiltinT
)

type Object interface {
	Type() ObjectType
	String() string
}
