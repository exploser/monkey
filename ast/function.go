package ast

import (
	"fmt"
	"strings"

	"git.exsdev.ru/ExS/monkey/token"
)

var _ Expression = new(FunctionLiteral)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (*FunctionLiteral) expr() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	params := make([]string, 0, len(fl.Parameters))
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	return fmt.Sprintf("%s (%s) %s", fl.TokenLiteral(), strings.Join(params, ", "), fl.Body)
}
