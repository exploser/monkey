package types

var _ Object = new(Builtin)

const BuiltinT ObjectType = "Builtin"

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (*Builtin) Type() ObjectType {
	return BuiltinT
}

func (b *Builtin) String() string {
	return "_builtin(...)"
}
