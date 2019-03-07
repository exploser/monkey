package types

import "fmt"

var _ Object = new(Boolean)

type Boolean struct {
	Value bool
}

func (b *Boolean) String() string {
	return fmt.Sprint(b.Value)
}
