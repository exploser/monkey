package types

var _ Object = new(Nil)

type Nil struct {
}

func (*Nil) Type() ObjectType {
	return NilT
}

func (i *Nil) String() string {
	return "nil"
}
