package types

var _ Object = new(String)

const StringT ObjectType = "String"

type String struct {
	Value string
}

func (*String) Type() ObjectType {
	return StringT
}

func (i *String) String() string {
	return i.Value
}
