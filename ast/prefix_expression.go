package ast

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/token"
)

var _ Expression = new(PrefixExpression)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (*PrefixExpression) expr() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", pe.Operator, pe.Right)
}
