package types

var _ Object = new(Return)

const ReturnT ObjectType = "Return"

type Return struct {
	Value Object
}

func (*Return) Type() ObjectType {
	return ReturnT
}

func (r *Return) String() string {
	return r.Value.String()
}
