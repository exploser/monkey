package lexer

import (
	"github.com/vasilevp/monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()

	return l
}

func (l *Lexer) NextToken() (tok token.Token) {
	l.eatWhitespace()

	tok.Literal = string(l.ch)

	switch l.ch {
	case '=':
		tok.Type = token.Assign
		if l.peekChar() == '=' {
			c := l.ch
			l.readChar()
			tok.Type = token.Equals
			tok.Literal = string(c) + string(l.ch)
		}
	case '+':
		tok.Type = token.Plus
	case '-':
		tok.Type = token.Minus
	case '*':
		tok.Type = token.Asterisk
	case '/':
		tok.Type = token.Slash

	case '(':
		tok.Type = token.LParen
	case ')':
		tok.Type = token.RParen

	case '{':
		tok.Type = token.LBrace
	case '}':
		tok.Type = token.RBrace

	case '[':
		tok.Type = token.LBracket
	case ']':
		tok.Type = token.RBracket

	case '<':
		tok.Type = token.LessThan
	case '>':
		tok.Type = token.GreaterThan

	case ',':
		tok.Type = token.Comma
	case ';':
		tok.Type = token.Semicolon

	case '!':
		tok.Type = token.Bang
		if l.peekChar() == '=' {
			c := l.ch
			l.readChar()
			tok.Type = token.NotEqual
			tok.Literal = string(c) + string(l.ch)
		}
	case ':':
		tok.Type = token.Illegal
		if l.peekChar() == '=' {
			c := l.ch
			l.readChar()
			tok.Type = token.DeclareAssign
			tok.Literal = string(c) + string(l.ch)
		}
	case '"':
		tok.Type = token.String
		tok.Literal = l.readString()

	case 0:
		tok.Type = token.EOF
	default:
		switch {
		case isLetter(l.ch):
			ident := l.readIdentifier()
			tok = token.Token{token.LookupIdent(ident), ident}
			return
		case isDigit(l.ch):
			ident := l.readNumber()
			tok = token.Token{token.Int, ident}
			return
		default:
			tok.Type = token.Illegal
		}
	}

	l.readChar()
	return
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) eatWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}

	return l.input[position:l.position]
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
