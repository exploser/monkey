package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
)

func TestStringify(t *testing.T) {
	tt := []struct {
		input  string
		output string
	}{
		{"a+1", "(a + 1); "},
		{"if(a){b}else{c}", "if (a) { b; } else { c; }; "},
		{"fn(){a;b}", "fn () { a; b; }; "},
		{"fn(){a;b}()", "fn () { a; b; }(); "},
		{"a(b)", "a(b); "},
		{`len("")`, `len(""); `},
	}

	for _, tc := range tt {
		l := lexer.New(tc.input)
		p := parser.New(l)
		prog := p.ParseProgram()
		require.Equal(t, tc.output, prog.String())
	}
}
