package ast

import (
	"fmt"

	"git.exsdev.ru/ExS/gop/token"
)

var _ Expression = new(IfExpression)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (*IfExpression) expr() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	res := fmt.Sprintf("if %s %s", ie.Condition, ie.Consequence)
	if ie.Alternative != nil {
		res += fmt.Sprintf(" else %s", ie.Alternative)
	}

	return res
}
