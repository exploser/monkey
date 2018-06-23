package ast

import (
	"fmt"

	"git.exsdev.ru/ExS/gop/token"
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
	return fmt.Sprintf("%v %v", l.Token.Literal, l.Expression)
}
