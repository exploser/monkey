package parser_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
)

func TestCallExpression(t *testing.T) {
	input := `add(1, 2 * 3, 4 + 5);`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])
	stmt := program.Statements[0].(*ast.ExpressionStatement)

	require.IsType(t, new(ast.CallExpression), stmt.Expression)
	call := stmt.Expression.(*ast.CallExpression)

	testIdentifier(t, call.Function, "add")

	testLiteralExpression(t, call.Arguments[0], 1)
	testInfixExpression(t, call.Arguments[1], 2, "*", 3)
	testInfixExpression(t, call.Arguments[2], 4, "+", 5)
}
