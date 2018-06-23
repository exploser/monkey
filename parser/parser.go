package parser

import (
	"github.com/pkg/errors"

	"git.exsdev.ru/ExS/gop/ast"
	"git.exsdev.ru/ExS/gop/lexer"
	"git.exsdev.ru/ExS/gop/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []error
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := ast.Program{}
	prog.Statements = make([]ast.Statement, 0, 64)

	for p.curToken.Type != token.EOF {
		if stmt := p.parseStatement(); stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}

		p.nextToken()
	}

	return &prog
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.Ident) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.Assign) {
		return nil
	}

	// TODO: not implemented
	for p.curToken.Type != token.Semicolon {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	// TODO: not implemented
	for p.curToken.Type != token.Semicolon {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) expectPeek(expect token.TokenType) bool {
	if p.peekToken.Type == expect {
		p.nextToken()
		return true
	}

	p.errors = append(p.errors, errors.Errorf("Expected token %q, got %q", expect, p.peekToken))
	return false
}
