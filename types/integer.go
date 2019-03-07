package types

import "fmt"

var _ Object = new(Integer)

type Integer struct {
	Value int64
}

func (i *Integer) String() string {
	return fmt.Sprint(i.Value)
}

func (*Integer) object() {}
