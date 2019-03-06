package types

var _ Object = new(String)

type String struct {
	Value string
}

func (*String) Type() ObjectType {
	return StringT
}

func (i *String) String() string {
	return i.Value
}
