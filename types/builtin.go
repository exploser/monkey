package types

var _ Object = new(Builtin)

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
