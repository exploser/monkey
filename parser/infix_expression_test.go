package parser_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vasilevp/monkey/ast"
	"github.com/vasilevp/monkey/lexer"
	"github.com/vasilevp/monkey/parser"
)

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
