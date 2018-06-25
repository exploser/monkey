package ast

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/token"
)

var _ Expression = new(InfixExpression)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (*InfixExpression) expr() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left, ie.Operator, ie.Right)
}
