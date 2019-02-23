package ast

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/token"
)

var _ Expression = new(Declare)

type Declare struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (*Declare) expr() {}

func (l *Declare) TokenLiteral() string {
	return l.Token.Literal
}

func (l *Declare) String() string {
	return fmt.Sprintf("%v := %v", l.Name, l.Value)
}
