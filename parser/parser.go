package parser

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"git.exsdev.ru/ExS/gop/ast"
	"git.exsdev.ru/ExS/gop/lexer"
	"git.exsdev.ru/ExS/gop/token"
)

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

const (
	lowest = iota
	equals
	lessgreater
	sum
	product
	prefix
	call
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []error

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.registerPrefix(token.Ident, p.parseIdentifier)
	p.registerPrefix(token.Int, p.parseIntegerLiteral)
	p.registerPrefix(token.Bang, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := ast.Program{}
	prog.Statements = make([]ast.Statement, 0, 64)

	for p.curToken.Type != token.EOF {
		if stmt := p.parseStatement(); stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
			fmt.Println(":::", stmt)
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
		return p.parseExpressionStatement()
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

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(lowest)

	if p.peekToken.Type == token.Semicolon {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, errors.Wrapf(err, "could not parse integer literal %q", p.curToken.Literal))
		return nil
	}

	lit.Value = value
	return &lit
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.errors = append(p.errors, errors.Errorf("No prefix handler for %q found", p.curToken.Type))
		return nil
	}

	leftExp := prefix()
	return leftExp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression(prefix)

	return &expression
}

func (p *Parser) expectPeek(expect token.TokenType) bool {
	if p.peekToken.Type == expect {
		p.nextToken()
		return true
	}

	p.errors = append(p.errors, errors.Errorf("Expected token %q, got %q", expect, p.peekToken))
	return false
}

func (p *Parser) registerPrefix(tt token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}

func (p *Parser) registerInfix(tt token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}
