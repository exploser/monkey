package ast

import (
	"fmt"

	"github.com/vasilevp/monkey/token"
)

var _ Expression = new(ArrayLiteral)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (*ArrayLiteral) expr() {}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayLiteral) String() string {
	return fmt.Sprint(al.Elements)
}
