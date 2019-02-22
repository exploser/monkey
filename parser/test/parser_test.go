package parser_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
)

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{
			"1 + (2 + 3) + 4;",
			"((1 + (2 + 3)) + 4); ",
		},
		{
			"1 + 2 * 3 + 4;",
			"((1 + (2 * 3)) + 4); ",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		prog := p.ParseProgram()
		checkParserErrors(t, p)
		require.Equal(t, tt.expected, prog.String())
	}
}

// Helper functions

func testIdentifier(t *testing.T, expr ast.Expression, value string) {
	require.IsType(t, new(ast.Identifier), expr)
	ident := expr.(*ast.Identifier)

	require.Equal(t, ident.Value, value)
	require.Equal(t, ident.TokenLiteral(), value)
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	require.Equal(t, s.TokenLiteral(), "let")
	require.IsType(t, new(ast.LetStatement), s)
	let := s.(*ast.LetStatement)
	require.Equal(t, name, let.Name.Value)
	require.Equal(t, name, let.Name.TokenLiteral())
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	if !assert.Empty(t, p.Errors()) {
		for _, e := range p.Errors() {
			t.Error(e)
		}

		t.FailNow()
	}
}

func testIntegerLiteral(t *testing.T, expr ast.Expression, value int64) {
	require.IsType(t, new(ast.IntegerLiteral), expr)
	integerLiteral := expr.(*ast.IntegerLiteral)

	require.Equal(t, value, integerLiteral.Value)
	require.Equal(t, fmt.Sprint(value), integerLiteral.TokenLiteral())
}

func testBooleanLiteral(t *testing.T, expr ast.Expression, value bool) {
	require.IsType(t, new(ast.Boolean), expr)
	b := expr.(*ast.Boolean)
	require.Equal(t, value, b.Value)
	require.Equal(t, fmt.Sprintf("%t", value), b.TokenLiteral())
}

func testLiteralExpression(t *testing.T, expr ast.Expression, expected interface{}) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, expr, int64(v))
	case int32:
		testIntegerLiteral(t, expr, int64(v))
	case int64:
		testIntegerLiteral(t, expr, int64(v))
	case string:
		testIdentifier(t, expr, v)
	case bool:
		testBooleanLiteral(t, expr, v)
	default:
		require.Failf(t, "Unknown expectation type", "Unknown expectation type '%T'", expected)
	}
}

func testInfixExpression(t *testing.T, expr ast.Expression, left interface{}, operator string, right interface{}) {
	require.IsType(t, new(ast.InfixExpression), expr)
	ie := expr.(*ast.InfixExpression)
	testLiteralExpression(t, ie.Left, left)
	require.Equal(t, operator, ie.Operator)
	testLiteralExpression(t, ie.Right, right)
}
