package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "Illegal"
	EOF     = "EOF"

	Ident  = "Ident"
	Int    = "Int"
	String = "String"

	Assign   = "Assign"
	Plus     = "Plus"
	Minus    = "Minus"
	Bang     = "Bang"
	Asterisk = "Asterisk"
	Slash    = "Slash"

	Comma     = "Comma"
	Semicolon = "Semicolon"

	LParen = "LParen"
	RParen = "RParen"

	LBrace = "LBrace"
	RBrace = "RBrace"

	LessThan    = "LessThan"
	GreaterThan = "GreaterThan"
	Equals      = "Equals"
	NotEqual    = "NotEqual"

	Function = "Function"
	Let      = "Let"
	True     = "True"
	False    = "False"
	If       = "If"
	Else     = "Else"
	Return   = "Return"
)

func LookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"fn":     Function,
		"let":    Let,
		"if":     If,
		"else":   Else,
		"return": Return,
		"true":   True,
		"false":  False,
	}

	if t, ok := keywords[ident]; ok {
		return t
	}

	return Ident
}
