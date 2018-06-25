package parser_test

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 1, len(program.Statements), "program should have 1 statement")

	for _, stmt := range program.Statements {
		require.IsType(t, new(ast.ExpressionStatement), stmt)

		expr := stmt.(*ast.ExpressionStatement)
		testIdentifier(t, expr.Expression, "foobar")
	}
}
