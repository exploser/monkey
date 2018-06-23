package ast

import (
	"git.exsdev.ru/ExS/gop/token"
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
