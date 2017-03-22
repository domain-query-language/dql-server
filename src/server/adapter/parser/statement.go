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
	ASSIGN	    // =
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
	OBJECT	    // object->key

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
	token.ARROW:	OBJECT,
	token.LBRACKET: INDEX,
	token.ASSIGN: 	ASSIGN,
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

	p := &statementParser{
		t: t,
	};

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INTEGER, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.STRING, p.parseString)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.PLUS, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BOOLEAN, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.OBJECTNAME, p.parseObjectCreation)
	p.registerPrefix(token.RUN, p.parseRunQuery)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOTEQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.ARROW, p.parseInfixExpression)
	p.registerInfix(token.ASSIGN, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseMethodCallExpression)
	p.registerInfix(token.LBRACKET, p.parseArrayAccess)

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

func (p *statementParser) expectCurrent(t token.TokenType) bool {

	if p.curTokenIs(t) {

		p.nextToken()
		return true
	} else {

		p.peekError(t)
		return false
	}
}

func (p *statementParser) peekError(t token.TokenType) {

	if (p.peekToken == nil) {
		p.logError("Expected next token to be '%s', got EOF instead", t)
		return;
	}
	p.logError("Expected next token to be '%s', got '%s' instead", t, p.peekToken.Val)
}

func (p *statementParser) logError(format string, a...interface{}) {
	msg := fmt.Sprintf(format, a...)
	p.error = errors.New(msg)
}


func (p *statementParser) ParseBlockStatement() (*ast.BlockStatement, error) {

	if (!p.curTokenIs(token.LBRACE) ) {
		p.logError("Expected next token to be '%s', got '%s' instead", token.LBRACE, p.curToken.Val)
		return nil, p.error
	}

	p.nextToken()

	block := &ast.BlockStatement{Type:ast.BLOCK_STATEMENT}

	block.Statements = []ast.Node{}

	for !p.curTokenIs(token.RBRACE) && p.error == nil  {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()

		if p.curToken == nil {
			p.logError("Unexpected EOF")
			break;
		}
	}

	return block, p.error
}

func (p *statementParser) parseStatement() ast.Node {

	switch p.curToken.Type {

		case token.RETURN:
			return p.parseReturnStatement()

		case token.IF:
			return p.parseIfStatement()

		case token.FOREACH:
			return p.parseForeachStatement()

		case token.ASSERT:
			return p.parseAssertStatement()

		case token.APPLY:
			return p.parseApplyStatement()

		default:
			return p.parseExpressionStatement()
	}
}


func (p *statementParser) parseExpressionStatement() ast.Node {

	stmt := &ast.ExpressionStatement{Type:ast.EXPRESSION_STATEMENT}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	} else {
		p.logError("Expected %v, got %v", token.SEMICOLON, p.peekToken);
	}

	return stmt
}


func (p *statementParser) parseReturnStatement() *ast.Return {

	stmt := &ast.Return{Type:ast.RETURN_STATEMENT}

	p.nextToken()

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *statementParser) parseForeachStatement() *ast.ForeachStatement {

	stmt := &ast.ForeachStatement{Type:ast.FOREACH_STATEMENT}

	if (!p.expectPeek(token.LPAREN)) {
		return nil
	}

	p.nextToken()

	stmt.Collection = p.parseExpression(LOWEST)

	if (!p.expectPeek(token.AS)) {
		return nil
	}

	p.nextToken()

	if (p.peekTokenIs(token.STRONGARROW)){
		stmt.Key = p.parseIdentifier().(*ast.Identifier)

		p.nextToken()
		p.nextToken()
	}

	stmt.Value =  p.parseIdentifier().(*ast.Identifier)

	if (!p.expectPeek(token.RPAREN)) {
		return nil
	}

	if (!p.expectPeek(token.LBRACE)) {
		return nil
	}

	body, _ := p.ParseBlockStatement()

	stmt.Body = body

	return stmt
}

func (p *statementParser) parseIfStatement() ast.Statement {

	stmt := &ast.IfStatement{Type:ast.IF_STATEMENT}

	p.nextToken()

	stmt.Test = p.parseExpression(LOWEST)

	if (!p.expectPeek(token.LBRACE)) {
		return nil
	}

	consequent, _ := p.ParseBlockStatement()

	stmt.Consequent = consequent

	if p.peekTokenIs(token.ELSE) {

		p.nextToken()
		p.nextToken()

		stmt.Alternate, _ = p.ParseBlockStatement()
	}

	return stmt
}

func (p *statementParser) parseAssertStatement() ast.Statement {

	a := &ast.AssertStatement{ast.ASSERT_STATEMENT, "", nil};

	if (!p.expectPeek(token.IDENT)) {
		return nil
	}

	if (p.curToken.Val != "invariant") {
		p.logError("Expected next identifier to be 'invariant', got '%s' instead", p.curToken.Val)
		return nil
	}

	if (p.peekTokenIs(token.NOT)) {
		a.Operator = "not"
		p.nextToken()
	}

	p.nextToken()

	a.Event = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return a
}

func (p *statementParser) parseApplyStatement() ast.Statement {

	a := &ast.ApplyStatement{ast.APPLY_STATEMENT, nil};

	if (!p.expectPeek(token.IDENT)) {
		return nil
	}

	if (p.curToken.Val != "event") {
		p.logError("Expected next identifier to be 'event', got '%s' instead", p.curToken.Val)
		return nil
	}

	p.nextToken()

	a.Event = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return a
}


func (p *statementParser) parseExpression(precedence int) ast.Expression {

	if (p.curToken == nil) {

		return nil
	}

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

	if (p.peekToken == nil) {
		return LOWEST
	}

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

func (p *statementParser) parsePrefixExpression() ast.Expression {

	operator := p.curToken.Val

	if (p.isIncrementOrDecrement()) {
		p.nextToken()
		operator += p.curToken.Val
	}

	expression := &ast.Prefix{
		Type: ast.PREFIX,
		Operator: operator,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *statementParser) isIncrementOrDecrement() bool {

	return (p.curToken.Type == token.PLUS || p.curToken.Type == token.MINUS) && p.curToken.Type == p.peekToken.Type
}

func (p *statementParser) parseInfixExpression(left ast.Expression) ast.Expression {

	expression := &ast.Infix{
		Type: ast.INFIX,
		Operator: p.curToken.Val,
		Left:     left,
	}

	precedence := p.curPrecedence()

	p.nextToken()

	expression.Right = p.parseExpression(precedence)

	if (expression.Right == nil) {
		p.logError("Expected expression, got nothing")
		return nil
	}

	return expression
}

func (p *statementParser) parseIdentifier() ast.Expression {

	return &ast.Identifier{Type:ast.IDENTIFIER, Value: p.curToken.Val}
}

func (p *statementParser) parseIntegerLiteral() ast.Expression {

	lit := &ast.Integer{Type:ast.INTEGER}

	value, err := strconv.ParseInt(p.curToken.Val, 0, 64)

	if err != nil {
		p.logError("could not parse %q as integer", p.curToken.Val)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *statementParser) parseBoolean() ast.Expression {

	return &ast.Boolean{Type:ast.BOOLEAN, Value: p.curToken.Val == "true"}
}

func (p *statementParser) parseGroupedExpression() ast.Expression {

	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *statementParser) parseFloatLiteral() ast.Expression {

	lit := &ast.Float{Type:ast.FLOAT}

	value, err := strconv.ParseFloat(p.curToken.Val, 64)
	if err != nil {
		p.logError("could not parse %q as float", p.curToken.Val);
		return nil
	}

	lit.Value = value

	return lit
}

func (p *statementParser) parseString() ast.Expression {

	return &ast.String{Type:ast.STRING, Value:p.curToken.Val}
}

func (p *statementParser) parseMethodCallExpression(method ast.Expression) ast.Expression {

	exp := &ast.MethodCall{Type: ast.METHOD_CALL, Method: method}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

func (p *statementParser) parseExpressionList(end token.TokenType) []ast.Expression {

	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}


func (p *statementParser) parseArrayLiteral() ast.Expression {

	array := &ast.Array{Type: ast.ARRAY}

	array.Elements = p.parseExpressionList(token.RBRACKET)

	return array
}

func (p *statementParser) parseArrayAccess(left ast.Expression) ast.Expression {

	exp := &ast.ArrayAccess{Type: ast.ARRAY_ACCESS, Left: left}

	p.nextToken()
	exp.Offset = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}

func (p *statementParser) parseObjectCreation() ast.Expression {

	oc := &ast.ObjectCreation{ast.OBJECT_CREATION, p.curToken.Val, []ast.Expression{}};

	if (p.peekTokenIs(token.LPAREN)) {
		p.nextToken()
		oc.Arguments = p.parseExpressionList(token.RPAREN)
	}

	return oc
}

func (p* statementParser) parseRunQuery() ast.Expression {

	r := &ast.RunQuery{ast.RUN_QUERY, nil}

	if (!p.expectPeek(token.IDENT)) {
		return nil
	}

	if (p.curToken.Val != "query") {
		p.logError("Expected next identifier to be 'query', got '%s' instead", p.curToken.Val)
		return nil
	}

	p.nextToken()

	r.Query = p.parseExpression(LOWEST)

	return r
}
