package types

import "fmt"

var _ Object = new(Integer)

type Integer struct {
	Value int64
}

func (*Integer) Type() ObjectType {
	return IntegerT
}

func (i *Integer) String() string {
	return fmt.Sprint(i.Value)
}
