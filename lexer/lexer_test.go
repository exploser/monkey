package lexer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/vasilevp/monkey/lexer"
	"github.com/vasilevp/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	if (5 == 10) {
		return 5 != 10;
	}
	"foobar"
	"foo bar"
	[1, 2];
	a:=7
	`

	tests := []token.Token{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},

		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},

		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LParen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.RParen, ")"},
		{token.LBrace, "{"},

		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},

		{token.RBrace, "}"},
		{token.Semicolon, ";"},

		{token.Let, "let"},
		{token.Ident, "result"},
		{token.Assign, "="},
		{token.Ident, "add"},
		{token.LParen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.RParen, ")"},
		{token.Semicolon, ";"},

		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Int, "5"},
		{token.Semicolon, ";"},

		{token.Int, "5"},
		{token.LessThan, "<"},
		{token.Int, "10"},
		{token.GreaterThan, ">"},
		{token.Int, "5"},
		{token.Semicolon, ";"},

		{token.If, "if"},
		{token.LParen, "("},
		{token.Int, "5"},
		{token.LessThan, "<"},
		{token.Int, "10"},
		{token.RParen, ")"},
		{token.LBrace, "{"},

		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},

		{token.RBrace, "}"},
		{token.Else, "else"},
		{token.LBrace, "{"},

		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.RBrace, "}"},

		{token.If, "if"},
		{token.LParen, "("},
		{token.Int, "5"},
		{token.Equals, "=="},
		{token.Int, "10"},
		{token.RParen, ")"},
		{token.LBrace, "{"},

		{token.Return, "return"},
		{token.Int, "5"},
		{token.NotEqual, "!="},
		{token.Int, "10"},
		{token.Semicolon, ";"},

		{token.RBrace, "}"},

		{token.String, "foobar"},

		{token.String, "foo bar"},

		{token.LBracket, "["},
		{token.Int, "1"},
		{token.Comma, ","},
		{token.Int, "2"},
		{token.RBracket, "]"},
		{token.Semicolon, ";"},

		{token.Ident, "a"},
		{token.DeclareAssign, ":="},
		{token.Int, "7"},

		{token.EOF, "\x00"},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()
		assert.Equal(t, tt, tok)
	}
}
