package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
)

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world!";`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 1, len(program.Statements), "program should have 1 statement")

	for _, stmt := range program.Statements {
		require.IsType(t, new(ast.ExpressionStatement), stmt)

		expr := stmt.(*ast.ExpressionStatement)
		require.IsType(t, new(ast.StringLiteral), expr.Expression)

		str := expr.Expression.(*ast.StringLiteral)
		require.Equal(t, str.Value, "hello world!")
		require.Equal(t, str.TokenLiteral(), "hello world!")
	}
}
