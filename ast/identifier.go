package ast

import (
	"github.com/vasilevp/monkey/token"
)

var _ Expression = new(Identifier)

type Identifier struct {
	Token token.Token
	Value string
}

func (*Identifier) expr() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (l *Identifier) String() string {
	return l.Value
}
