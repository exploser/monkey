package ast

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/token"
)

var _ Statement = new(LetStatement)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (*LetStatement) stmt() {}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) String() string {
	return fmt.Sprintf("%v %v = %v", l.Token.Literal, l.Name, l.Value)
}
