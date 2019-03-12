package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vasilevp/monkey/ast"
	"github.com/vasilevp/monkey/lexer"
	"github.com/vasilevp/monkey/parser"
)

func TestDeclareStatements(t *testing.T) {
	input := `
	x := 5;
	y := 10;
	foobar:= 838383;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 3, len(program.Statements))

	tests := []struct {
		name  string
		value int
	}{
		{"x", 5},
		{"y", 10},
		{"foobar", 838383},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		require.IsType(t, new(ast.ExpressionStatement), stmt)

		expr := stmt.(*ast.ExpressionStatement)

		testDeclareAssign(t, expr.Expression, tt.name, tt.value)
	}
}
