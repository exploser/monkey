package parser_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
)

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
