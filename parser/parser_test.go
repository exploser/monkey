package parser_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/gop/ast"
	"git.exsdev.ru/ExS/gop/lexer"
	"git.exsdev.ru/ExS/gop/parser"
	"github.com/stretchr/testify/assert"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar= 838383;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 3, len(program.Statements))

	tests := []string{"x", "y", "foobar"}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt)
	}
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

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 3, len(program.Statements), "program should have 3 statements")

	for _, stmt := range program.Statements {
		assert.Equal(t, stmt.TokenLiteral(), "return")
		assert.IsType(t, new(ast.ReturnStatement), stmt)
	}
}

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
		require.IsType(t, new(ast.Identifier), expr.Expression)

		ident := expr.Expression.(*ast.Identifier)
		require.Equal(t, ident.Value, "foobar")
		require.Equal(t, ident.TokenLiteral(), "foobar")
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 1, len(program.Statements), "program should have 1 statement")

	for _, stmt := range program.Statements {
		require.IsType(t, new(ast.ExpressionStatement), stmt)

		expr := stmt.(*ast.ExpressionStatement)
		require.IsType(t, new(ast.IntegerLiteral), expr.Expression)

		integer := expr.Expression.(*ast.IntegerLiteral)
		require.Equal(t, integer.Value, int64(5))
		require.Equal(t, integer.TokenLiteral(), "5")
	}
}

func TestPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		require.Equal(t, 1, len(program.Statements))
		require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])

		expr := program.Statements[0].(*ast.ExpressionStatement)

		require.IsType(t, new(ast.PrefixExpression), expr.Expression)
		prefix := expr.Expression.(*ast.PrefixExpression)

		require.IsType(t, new(ast.IntegerLiteral), prefix.Right)
		integerLiteral := prefix.Right.(*ast.IntegerLiteral)

		require.Equal(t, tt.integerValue, integerLiteral.Value)
		require.Equal(t, fmt.Sprint(tt.integerValue), integerLiteral.TokenLiteral())
	}
}

func TestInfixExpression(t *testing.T) {
	infixTests := []struct {
		input    string
		left     int64
		operator string
		right    int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

}
