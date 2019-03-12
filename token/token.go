package token

type Type int
type Token struct {
	Type    Type
	Literal string
}

//go:generate stringer -type=Type
const (
	Illegal Type = iota
	EOF
	Ident
	Int
	String
	Nil
	Assign
	Plus
	Minus
	Bang
	Asterisk
	Slash
	Comma
	Semicolon
	LParen
	RParen
	LBrace
	RBrace
	LessThan
	GreaterThan
	Equals
	NotEqual
	Function
	Let
	DeclareAssign
	True
	False
	If
	Else
	Return
	LBracket
	RBracket
)

func LookupIdent(ident string) Type {
	keywords := map[string]Type{
		"fn":     Function,
		"let":    Let,
		"if":     If,
		"else":   Else,
		"return": Return,
		"true":   True,
		"false":  False,
		"nil":    Nil,
	}

	if t, ok := keywords[ident]; ok {
		return t
	}

	return Ident
}
