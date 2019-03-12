package ast

import "github.com/vasilevp/monkey/token"

var _ Expression = new(IntegerLiteral)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (*IntegerLiteral) expr() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
