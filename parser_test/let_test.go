package parser_test

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
