package ast

type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
	expr()
}

type Statement interface {
	Node
	stmt()
}
