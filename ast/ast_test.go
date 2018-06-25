package ast_test

import (
	"testing"

	. "git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/token"
	"github.com/stretchr/testify/assert"
)

func TestAstString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.Let, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.Ident, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.Ident, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	assert.Equal(t, "let myVar = anotherVar; ", program.String())
}
