package ast

import "github.com/vasilevp/monkey/token"

var _ Expression = new(Nil)

type Nil struct {
	Token token.Token
}

func (*Nil) expr() {}

func (s *Nil) TokenLiteral() string {
	return s.Token.Literal
}

func (s *Nil) String() string {
	return s.Token.Literal
}
