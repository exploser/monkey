package types

import (
	"fmt"
)

var _ Object = new(Array)

type Array struct {
	Elements []Object
}

func (b *Array) String() string {
	return fmt.Sprint(b.Elements)
}

func (*Array) object() {}
