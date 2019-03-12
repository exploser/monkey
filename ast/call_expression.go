package ast

import (
	"fmt"
	"strings"

	"github.com/vasilevp/monkey/token"
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

func (ce *CallExpression) String() string {
	params := make([]string, 0, len(ce.Arguments))
	for _, p := range ce.Arguments {
		params = append(params, p.String())
	}

	return fmt.Sprintf("%s(%s)", ce.Function.String(), strings.Join(params, ", "))
}
