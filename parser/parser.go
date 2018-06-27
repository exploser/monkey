package parser

import (
	"strconv"

	"github.com/pkg/errors"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/token"
)

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

type precedence int

const (
	lowest precedence = iota
	equals
	lessgreater
	sum
	product
	prefix
	call
)

var precedences = map[token.TokenType]precedence{
	token.Equals:      equals,
	token.NotEqual:    equals,
	token.LessThan:    lessgreater,
	token.GreaterThan: lessgreater,
	token.Plus:        sum,
	token.Minus:       sum,
	token.Asterisk:    product,
	token.Slash:       product,
	token.LParen:      call,
}

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

	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Minus, p.parseInfixExpression)
	p.registerInfix(token.Asterisk, p.parseInfixExpression)
	p.registerInfix(token.Slash, p.parseInfixExpression)

	p.registerInfix(token.Equals, p.parseInfixExpression)
	p.registerInfix(token.NotEqual, p.parseInfixExpression)
	p.registerInfix(token.LessThan, p.parseInfixExpression)
	p.registerInfix(token.GreaterThan, p.parseInfixExpression)

	p.registerPrefix(token.True, p.parseBoolean)
	p.registerPrefix(token.False, p.parseBoolean)

	p.registerPrefix(token.LParen, p.parseGroupedExpression)

	p.registerPrefix(token.If, p.parseIfExpression)

	p.registerPrefix(token.Function, p.parseFunctionLiteral)

	p.registerInfix(token.LParen, p.parseCallExpression)

	p.registerPrefix(token.String, p.parseStringLiteral)

	p.registerPrefix(token.LBracket, p.parseArrayExpression)
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

	p.nextToken()

	stmt.Value = p.parseExpression(lowest)

	if p.peekToken.Type == token.Semicolon {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	stmt.ReturnValue = p.parseExpression(lowest)

	if p.peekToken.Type == token.Semicolon {
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

func (p *Parser) parseExpression(prec precedence) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.errors = append(p.errors, errors.Errorf("No prefix handler for %q found", p.curToken.Type))
		return nil
	}

	leftExp := prefix()
	for p.peekToken.Type != token.Semicolon && prec < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

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

func (p *Parser) parseBoolean() ast.Expression {
	expression := ast.Boolean{
		Token: p.curToken,
		Value: p.curToken.Type == token.True,
	}

	return &expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	prec := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(prec)

	return &expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	expression := p.parseExpression(lowest)

	if !p.expectPeek(token.RParen) {
		return nil
	}

	return expression
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := ast.IfExpression{Token: p.curToken}

	if !p.expectPeek(token.LParen) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(lowest)

	if !p.expectPeek(token.RParen) {
		return nil
	}

	if !p.expectPeek(token.LBrace) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.peekToken.Type == token.Else {
		p.nextToken()

		if !p.expectPeek(token.LBrace) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return &expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := ast.BlockStatement{Token: p.curToken}
	block.Statements = make([]ast.Statement, 0)

	p.nextToken()

	for p.curToken.Type != token.RBrace && p.curToken.Type != token.EOF {
		if stmt := p.parseStatement(); stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	return &block
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	fn := ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LParen) {
		return nil
	}

	fn.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBrace) {
		return nil
	}

	fn.Body = p.parseBlockStatement()
	return &fn
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := make([]*ast.Identifier, 0)

	if p.peekToken.Type == token.RParen {
		p.nextToken()
		return nil
	}

	p.nextToken()

	ident := ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, &ident)

	for p.peekToken.Type == token.Comma {
		p.nextToken()
		p.nextToken()

		ident := ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, &ident)
	}

	if !p.expectPeek(token.RParen) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RParen)
	return &exp
}

func (p *Parser) parseArrayExpression() ast.Expression {
	exp := ast.ArrayLiteral{Token: p.curToken}
	exp.Elements = p.parseExpressionList(token.RBracket)
	return &exp
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	expressions := make([]ast.Expression, 0)

	if p.peekToken.Type == end {
		p.nextToken()
		return nil
	}

	p.nextToken()

	expressions = append(expressions, p.parseExpression(lowest))

	for p.peekToken.Type == token.Comma {
		p.nextToken()
		p.nextToken()

		expressions = append(expressions, p.parseExpression(lowest))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return expressions
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
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

func (p *Parser) registerInfix(tt token.TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}

func (p *Parser) peekPrecedence() precedence {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return lowest
}

func (p *Parser) curPrecedence() precedence {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return lowest
}
