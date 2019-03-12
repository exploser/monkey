package parser_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vasilevp/monkey/ast"
	"github.com/vasilevp/monkey/lexer"
	"github.com/vasilevp/monkey/parser"
)

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
