package types

var _ Object = new(Error)
var _ error = new(Error)

const ErrorT ObjectType = "Error"

type Error struct {
	E error
}

func (*Error) Type() ObjectType {
	return ErrorT
}

func (i *Error) String() string {
	return i.E.Error()
}

func (i *Error) Error() string {
	return i.E.Error()
}
