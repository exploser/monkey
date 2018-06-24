package ast

import (
	"strings"

	"git.exsdev.ru/ExS/gop/token"
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
	var out strings.Builder
	out.WriteRune('(')
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteRune(')')

	return out.String()
}
