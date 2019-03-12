package ast

import (
	"strings"

	"github.com/vasilevp/monkey/token"
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
	result := strings.Builder{}
	for _, v := range bs.Statements {
		result.WriteString(v.String())
		result.WriteString("; ")
	}

	return result.String()
}
