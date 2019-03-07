package types

var _ Object = new(Return)

type Return struct {
	Value Object
}

func (r *Return) String() string {
	return r.Value.String()
}
