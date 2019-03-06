package types

import "fmt"

var _ Object = new(Boolean)

type Boolean struct {
	Value bool
}

func (*Boolean) Type() ObjectType {
	return BooleanT
}

func (b *Boolean) String() string {
	return fmt.Sprint(b.Value)
}
