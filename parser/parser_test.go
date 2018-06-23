package parser_test

import (
	"testing"

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
	errors := p.Errors()
	if len(errors) > 0 {
		t.Errorf("%d parser errors:", len(errors))
		for _, e := range p.Errors() {
			t.Error(e)
		}

		t.FailNow()
	}

	if program == nil {
		t.Fatal("program is nil")
	}

	if len(program.Statements) != 3 {
		t.Fatal("program should have 3 statements")
	}

	tests := []string{"x", "y", "foobar"}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if !assert.Equal(t, s.TokenLiteral(), "let") {
		return false
	}

	if !assert.IsType(t, new(ast.LetStatement), s) {
		return false
	}

	let := s.(*ast.LetStatement)

	if !assert.Equal(t, let.Name.Value, name) {
		return false
	}

	if !assert.Equal(t, let.Name.TokenLiteral(), name) {
		return false
	}

	return true
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
	errors := p.Errors()

	if !assert.Equal(t, len(errors), 0) {
		for _, e := range p.Errors() {
			t.Error(e)
		}

		t.FailNow()
	}

	if program == nil {
		t.Fatal("program is nil")
	}

	assert.Equal(t, len(program.Statements), 3, "program should have 3 statements")

	for _, stmt := range program.Statements {
		assert.Equal(t, stmt.TokenLiteral(), "return")
		assert.IsType(t, new(ast.ReturnStatement), stmt)
	}
}
