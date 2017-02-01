package parser

import (
	"errors"
	"fmt"
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/tokenizer"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOTEQ:    EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

/** Implementation of the adapter, written using the tokenizer, parser pattern */
type statementParser struct {

	t tokenizer.Tokenizer
	error error

	curToken  *token.Token
	peekToken *token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewStatement(statements string) *statementParser {

	t := tokenizer.NewTokenizer(statements);

	p := &statementParser{t}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INTEGER, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.PLUS, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BOOLEAN, p.parseBoolean)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.nextToken()
	p.nextToken()

	return p;
}

func (p *statementParser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *statementParser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *statementParser) nextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.t.Next();
}

func (p *statementParser) curTokenIs(t token.TokenType) bool {

	return p.curToken.Type == t
}

func (p *statementParser) peekTokenIs(t token.TokenType) bool {

	if p.peekToken == nil {
		return false
	}

	return p.peekToken.Type == t
}

func (p *statementParser) expectPeek(t token.TokenType) bool {

	if p.peekTokenIs(t) {

		p.nextToken()
		return true
	} else {

		p.peekError(t)
		return false
	}
}

func (p *statementParser) peekError(t token.TokenType) {

	if (p.peekToken == nil) {
		msg := fmt.Sprintf("Expected next token to be '%s', got EOF instead", t)
		p.error = errors.New(msg);
		return;
	}
	msg := fmt.Sprintf("Expected next token to be '%s', got '%s' instead", t, p.peekToken.Val)
	p.error = errors.New(msg);
}



func (p *statementParser) ParseBlockStatement() (ast.BlockStatement, error) {

	block := ast.BlockStatement{}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *statementParser) parseStatement() ast.Node {

	return p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
}

func (p *statementParser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {

		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *statementParser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *statementParser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *statementParser) parseIdentifier() ast.Expression {

	return &ast.Identifier{Value: p.curToken.Val}
}

func (p *statementParser) parseIntegerLiteral() ast.Expression {

	lit := ast.IntegerLiteral{}

	value, err := strconv.ParseInt(p.curToken.Val, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Val)
		p.error = errors.New(msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *statementParser) parsePrefixExpression() ast.Expression {

	expression := ast.Prefix{
		Type: "prefix",
		Operator: p.curToken.Val,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *statementParser) parseBoolean() ast.Expression {

	return ast.Boolean{Value: p.curToken.Val == "true"}
}