package ast

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/token"
)

var _ Statement = new(ReturnStatement)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (*ReturnStatement) stmt() {}

func (l *ReturnStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *ReturnStatement) String() string {
	return fmt.Sprintf("%v %v", l.Token.Literal, l.ReturnValue)
}
