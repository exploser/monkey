package types

var _ Object = new(Builtin)

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) String() string {
	return "_builtin(...)"
}

func (*Builtin) object() {}
