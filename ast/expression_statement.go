package ast

import (
	"github.com/vasilevp/monkey/token"
)

var _ Statement = new(ExpressionStatement)

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (*ExpressionStatement) stmt() {}

func (l *ExpressionStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *ExpressionStatement) String() string {
	return l.Expression.String()
}
