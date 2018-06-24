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
		testIdentifier(t, expr.Expression, "foobar")
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
		input    string
		operator string
		expected interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"-b;", "-", "b"},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		require.Len(t, program.Statements, 1)
		require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])

		expr := program.Statements[0].(*ast.ExpressionStatement)

		require.IsType(t, new(ast.PrefixExpression), expr.Expression)
		prefix := expr.Expression.(*ast.PrefixExpression)
		testLiteralExpression(t, prefix.Right, tt.expected)
	}
}

func TestInfixExpression(t *testing.T) {
	infixTests := []struct {
		input    string
		left     interface{}
		operator string
		right    interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"a > b;", "a", ">", "b"},
		{"true == true;", true, "==", true},
		{"false == false;", false, "==", false},
		{"true != false;", true, "!=", false},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := parser.New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		require.Len(t, program.Statements, 1)
		require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		testInfixExpression(t, stmt.Expression, tt.left, tt.operator, tt.right)
	}
}

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

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])
	stmt := program.Statements[0].(*ast.ExpressionStatement)

	require.IsType(t, new(ast.IfExpression), stmt.Expression)
	ifexp := stmt.Expression.(*ast.IfExpression)

	testInfixExpression(t, ifexp.Condition, "x", "<", "y")
	require.Len(t, ifexp.Consequence.Statements, 1)

	require.IsType(t, new(ast.ExpressionStatement), ifexp.Consequence.Statements[0])
	cons := ifexp.Consequence.Statements[0].(*ast.ExpressionStatement)
	testIdentifier(t, cons.Expression, "x")
	require.Nil(t, ifexp.Alternative)
}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])
	stmt := program.Statements[0].(*ast.ExpressionStatement)

	require.IsType(t, new(ast.IfExpression), stmt.Expression)
	ifexp := stmt.Expression.(*ast.IfExpression)

	testInfixExpression(t, ifexp.Condition, "x", "<", "y")
	require.Len(t, ifexp.Consequence.Statements, 1)

	require.IsType(t, new(ast.ExpressionStatement), ifexp.Consequence.Statements[0])
	cons := ifexp.Consequence.Statements[0].(*ast.ExpressionStatement)
	testIdentifier(t, cons.Expression, "x")

	require.NotNil(t, ifexp.Alternative)
	require.Len(t, ifexp.Alternative.Statements, 1)
}

func TestFunctionLiteral(t *testing.T) {
	input := `fn(x, y) { x + y; }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])
	stmt := program.Statements[0].(*ast.ExpressionStatement)

	require.IsType(t, new(ast.FunctionLiteral), stmt.Expression)
	fn := stmt.Expression.(*ast.FunctionLiteral)

	require.Len(t, fn.Parameters, 2)

	testLiteralExpression(t, fn.Parameters[0], "x")
	testLiteralExpression(t, fn.Parameters[1], "y")

	require.Len(t, fn.Body.Statements, 1)

	require.IsType(t, new(ast.ExpressionStatement), fn.Body.Statements[0])
	bstmt := fn.Body.Statements[0].(*ast.ExpressionStatement)

	testInfixExpression(t, bstmt.Expression, "x", "+", "y")
}

func TestFunctionParameters(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn(){};", expectedParams: []string{}},
		{input: "fn(x){};", expectedParams: []string{"x"}},
		{input: "fn(x,y,z){};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		require.Len(t, program.Statements, 1)

		require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])
		stmt := program.Statements[0].(*ast.ExpressionStatement)

		require.IsType(t, new(ast.FunctionLiteral), stmt.Expression)
		fn := stmt.Expression.(*ast.FunctionLiteral)

		require.Len(t, fn.Parameters, len(tt.expectedParams))

		for i, p := range tt.expectedParams {
			testLiteralExpression(t, fn.Parameters[i], p)
		}

		require.Len(t, fn.Body.Statements, 0)
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
