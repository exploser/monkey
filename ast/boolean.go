package ast

import "github.com/vasilevp/monkey/token"

var _ Expression = new(Boolean)

type Boolean struct {
	Token token.Token
	Value bool
}

func (*Boolean) expr() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}
