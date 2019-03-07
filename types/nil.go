package types

var _ Object = new(Nil)

type Nil struct {
}

func (i *Nil) String() string {
	return "nil"
}

func (*Nil) object() {}
