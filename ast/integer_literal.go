package ast

import "git.exsdev.ru/ExS/gop/token"

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
