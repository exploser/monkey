package ast

import "git.exsdev.ru/ExS/monkey/token"

var _ Expression = new(StringLiteral)

type StringLiteral struct {
	Token token.Token
	Value string
}

func (*StringLiteral) expr() {}

func (s *StringLiteral) TokenLiteral() string {
	return s.Token.Literal
}

func (s *StringLiteral) String() string {
	return s.Token.Literal
}
