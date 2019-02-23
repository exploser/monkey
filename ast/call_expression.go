package ast

import (
	"fmt"
	"strings"

	"git.exsdev.ru/ExS/monkey/token"
)

var _ Expression = new(CallExpression)

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (*CallExpression) expr() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (fl *CallExpression) String() string {
	params := make([]string, 0, len(fl.Arguments))
	for _, p := range fl.Arguments {
		params = append(params, p.String())
	}

	return fmt.Sprintf("%s(%s)", fl.Function.String(), strings.Join(params, ", "))
}
