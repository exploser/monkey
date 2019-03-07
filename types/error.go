package types

var _ Object = new(Error)

type Error struct {
	E error
}

func (e Error) String() string {
	return e.E.Error()
}

func (*Error) object() {}
