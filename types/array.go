package types

import (
	"fmt"
)

var _ Object = new(Array)

const ArrayT ObjectType = "Array"

type Array struct {
	Elements []Object
}

func (*Array) Type() ObjectType {
	return ArrayT
}

func (b *Array) String() string {
	return fmt.Sprint(b.Elements)
}
