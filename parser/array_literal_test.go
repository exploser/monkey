package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vasilevp/monkey/ast"
	"github.com/vasilevp/monkey/lexer"
	"github.com/vasilevp/monkey/parser"
)

func TestArrayLiteralExpression(t *testing.T) {
	input := "[1,2*2,3+3];"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 1, len(program.Statements), "program should have 1 statement")

	for _, stmt := range program.Statements {
		require.IsType(t, new(ast.ExpressionStatement), stmt)

		expr := stmt.(*ast.ExpressionStatement)
		require.IsType(t, new(ast.ArrayLiteral), expr.Expression)

		array := expr.Expression.(*ast.ArrayLiteral)
		require.Len(t, array.Elements, 3)

		testIntegerLiteral(t, array.Elements[0], 1)
		testInfixExpression(t, array.Elements[1], 2, "*", 2)
		testInfixExpression(t, array.Elements[2], 3, "+", 3)
	}
}
