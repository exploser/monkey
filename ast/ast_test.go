package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/vasilevp/monkey/ast"
	"github.com/vasilevp/monkey/token"
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
