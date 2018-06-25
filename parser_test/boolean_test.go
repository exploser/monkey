package parser_test

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBooleanLiteralExpression(t *testing.T) {
	input := "true;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 1, len(program.Statements), "program should have 1 statement")

	for _, stmt := range program.Statements {
		require.IsType(t, new(ast.ExpressionStatement), stmt)

		expr := stmt.(*ast.ExpressionStatement)
		require.IsType(t, new(ast.Boolean), expr.Expression)

		integer := expr.Expression.(*ast.Boolean)
		require.Equal(t, integer.Value, true)
		require.Equal(t, integer.TokenLiteral(), "true")
	}
}
