package types

var _ Object = new(String)

type String struct {
	Value string
}

func (i *String) String() string {
	return i.Value
}

func (*String) object() {}
