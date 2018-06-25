package ast

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/token"
)

var _ Statement = new(BlockStatement)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (*BlockStatement) stmt() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	return fmt.Sprint(bs.Statements)
}
