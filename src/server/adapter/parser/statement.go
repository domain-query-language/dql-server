package parser

import (
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

	tokenStream *tokenStream

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewStatement(statements string) *statementParser {

	t := tokenizer.NewTokenizer(statements);

	tknStrm := NewTokenStreamFromFreshTokenizer(t)

	p := &statementParser{
		tokenStream: tknStrm,
	};

	p.bootstrapRules()

	return p;
}

func NewStatementFromTokenizer(t tokenizer.Tokenizer, curToken *token.Token, peekToken *token.Token) *statementParser {

	tknStrm := NewTokenStreamFromExistingTokenizer(t, curToken, peekToken)

	p := &statementParser{
		tokenStream: tknStrm,
	};

	p.bootstrapRules()

	return p;
}

func (p *statementParser) bootstrapRules(){

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
}

func (p *statementParser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *statementParser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *statementParser) expectPeek(t token.TokenType) bool {

	return p.tokenStream.ExpectPeek(t)
}

func (p *statementParser) peekError(t token.TokenType) {

	p.tokenStream.PeekError(t)
}

func (p *statementParser) logError(format string, a...interface{}) {

	p.tokenStream.LogError(format, a...)
}


func (p *statementParser) ParseBlockStatementSurroundedBy(open token.TokenType, close token.TokenType) (*ast.BlockStatement, error) {

	if (!p.tokenStream.CurTokenIs(open)) {
		p.logError("Expected next token to be '%s', got '%s' instead", open, p.tokenStream.CurToken().Val)
		return nil, p.tokenStream.Error()
	}

	p.tokenStream.NextToken()

	block := &ast.BlockStatement{Type:ast.BLOCK_STATEMENT}

	block.Statements = []ast.Node{}

	for !p.tokenStream.CurTokenIs(close) && p.tokenStream.Error() == nil  {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.tokenStream.NextToken()

		if p.tokenStream.CurToken() == nil {
			p.logError("Unexpected EOF")
			break;
		}
	}

	return block, p.tokenStream.Error()
}

func (p *statementParser) ParseBlockStatement() (*ast.BlockStatement, error) {

	return p.ParseBlockStatementSurroundedBy(token.LBRACE, token.RBRACE)
}

func (p *statementParser) parseStatement() ast.Node {

	switch p.tokenStream.CurToken().Type {

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

	if p.tokenStream.PeekTokenIs(token.SEMICOLON) {
		p.tokenStream.NextToken()
	} else {
		p.logError("Expected %v, got %v", token.SEMICOLON, p.tokenStream.PeekToken());
	}

	return stmt
}


func (p *statementParser) parseReturnStatement() *ast.Return {

	stmt := &ast.Return{Type:ast.RETURN_STATEMENT}

	p.tokenStream.NextToken()

	stmt.Expression = p.parseExpression(LOWEST)

	if p.tokenStream.PeekTokenIs(token.SEMICOLON) {
		p.tokenStream.NextToken()
	}

	return stmt
}

func (p *statementParser) parseForeachStatement() *ast.ForeachStatement {

	stmt := &ast.ForeachStatement{Type:ast.FOREACH_STATEMENT}

	if (!p.tokenStream.ExpectPeek(token.LPAREN)) {
		return nil
	}

	p.tokenStream.NextToken()

	stmt.Collection = p.parseExpression(LOWEST)

	if (!p.tokenStream.ExpectPeek(token.AS)) {
		return nil
	}

	p.tokenStream.NextToken()

	if (p.tokenStream.PeekTokenIs(token.STRONGARROW)){
		stmt.Key = p.parseIdentifier().(*ast.Identifier)

		p.tokenStream.NextToken()
		p.tokenStream.NextToken()
	}

	stmt.Value =  p.parseIdentifier().(*ast.Identifier)

	if (!p.tokenStream.ExpectPeek(token.RPAREN)) {
		return nil
	}

	if (!p.tokenStream.ExpectPeek(token.LBRACE)) {
		return nil
	}

	body, _ := p.ParseBlockStatement()

	stmt.Body = body

	return stmt
}

func (p *statementParser) parseIfStatement() ast.Statement {

	stmt := &ast.IfStatement{Type:ast.IF_STATEMENT}

	p.tokenStream.NextToken()

	stmt.Test = p.parseExpression(LOWEST)

	if (!p.tokenStream.ExpectPeek(token.LBRACE)) {
		return nil
	}

	consequent, _ := p.ParseBlockStatement()

	stmt.Consequent = consequent

	if p.tokenStream.PeekTokenIs(token.ELSE) {

		p.tokenStream.NextToken()
		p.tokenStream.NextToken()

		stmt.Alternate, _ = p.ParseBlockStatement()
	}

	return stmt
}

func (p *statementParser) parseAssertStatement() ast.Statement {

	a := &ast.AssertStatement{ast.ASSERT_STATEMENT, "", nil};

	if (!p.tokenStream.ExpectPeek(token.IDENT)) {
		return nil
	}

	if (p.tokenStream.CurToken().Val != "invariant") {
		p.logError("Expected next identifier to be 'invariant', got '%s' instead", p.tokenStream.CurToken().Val)
		return nil
	}

	if (p.tokenStream.PeekTokenIs(token.NOT)) {
		a.Operator = "not"
		p.tokenStream.NextToken()
	}

	p.tokenStream.NextToken()

	a.Event = p.parseExpression(LOWEST)

	if p.tokenStream.PeekTokenIs(token.SEMICOLON) {
		p.tokenStream.NextToken()
	}

	return a
}

func (p *statementParser) parseApplyStatement() ast.Statement {

	a := &ast.ApplyStatement{ast.APPLY_STATEMENT, nil};

	if (!p.tokenStream.ExpectPeek(token.IDENT)) {
		return nil
	}

	if (p.tokenStream.CurToken().Val != "event") {
		p.logError("Expected next identifier to be 'event', got '%s' instead", p.tokenStream.CurToken().Val)
		return nil
	}

	p.tokenStream.NextToken()

	a.Event = p.parseExpression(LOWEST)

	if p.tokenStream.PeekTokenIs(token.SEMICOLON) {
		p.tokenStream.NextToken()
	}

	return a
}


func (p *statementParser) parseExpression(precedence int) ast.Expression {

	if (p.tokenStream.CurToken() == nil) {

		return nil
	}

	prefix := p.prefixParseFns[p.tokenStream.CurToken().Type]
	if prefix == nil {

		return nil
	}
	leftExp := prefix()

	for !p.tokenStream.PeekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {

		infix := p.infixParseFns[p.tokenStream.PeekToken().Type]
		if infix == nil {

			return leftExp
		}

		p.tokenStream.NextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *statementParser) peekPrecedence() int {

	if (p.tokenStream.PeekToken() == nil) {
		return LOWEST
	}

	if p, ok := precedences[p.tokenStream.PeekToken().Type]; ok {
		return p
	}

	return LOWEST
}

func (p *statementParser) curPrecedence() int {

	if p, ok := precedences[p.tokenStream.CurToken().Type]; ok {
		return p
	}

	return LOWEST
}

func (p *statementParser) parsePrefixExpression() ast.Expression {

	operator := p.tokenStream.CurToken().Val

	if (p.isIncrementOrDecrement()) {
		p.tokenStream.NextToken()
		operator += p.tokenStream.CurToken().Val
	}

	expression := &ast.Prefix{
		Type: ast.PREFIX,
		Operator: operator,
	}

	p.tokenStream.NextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *statementParser) isIncrementOrDecrement() bool {

	return (p.tokenStream.CurToken().Type == token.PLUS || p.tokenStream.CurToken().Type == token.MINUS) && p.tokenStream.CurToken().Type == p.tokenStream.PeekToken().Type
}

func (p *statementParser) parseInfixExpression(left ast.Expression) ast.Expression {

	expression := &ast.Infix{
		Type: ast.INFIX,
		Operator: p.tokenStream.CurToken().Val,
		Left:     left,
	}

	precedence := p.curPrecedence()

	p.tokenStream.NextToken()

	expression.Right = p.parseExpression(precedence)

	if (expression.Right == nil) {
		p.logError("Expected expression, got nothing")
		return nil
	}

	return expression
}

func (p *statementParser) parseIdentifier() ast.Expression {

	return &ast.Identifier{Type:ast.IDENTIFIER, Value: p.tokenStream.CurToken().Val}
}

func (p *statementParser) parseIntegerLiteral() ast.Expression {

	lit := &ast.Integer{Type:ast.INTEGER}

	value, err := strconv.ParseInt(p.tokenStream.CurToken().Val, 0, 64)

	if err != nil {
		p.logError("could not parse %q as integer", p.tokenStream.CurToken().Val)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *statementParser) parseBoolean() ast.Expression {

	return &ast.Boolean{Type:ast.BOOLEAN, Value: p.tokenStream.CurToken().Val == "true"}
}

func (p *statementParser) parseGroupedExpression() ast.Expression {

	p.tokenStream.NextToken()

	exp := p.parseExpression(LOWEST)

	if !p.tokenStream.ExpectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *statementParser) parseFloatLiteral() ast.Expression {

	lit := &ast.Float{Type:ast.FLOAT}

	value, err := strconv.ParseFloat(p.tokenStream.CurToken().Val, 64)
	if err != nil {
		p.logError("could not parse %q as float", p.tokenStream.CurToken().Val);
		return nil
	}

	lit.Value = value

	return lit
}

func (p *statementParser) parseString() ast.Expression {

	return &ast.String{Type:ast.STRING, Value:p.tokenStream.CurToken().Val}
}

func (p *statementParser) parseMethodCallExpression(method ast.Expression) ast.Expression {

	exp := &ast.MethodCall{Type: ast.METHOD_CALL, Method: method}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

func (p *statementParser) parseExpressionList(end token.TokenType) []ast.Expression {

	list := []ast.Expression{}

	if p.tokenStream.PeekTokenIs(end) {
		p.tokenStream.NextToken()
		return list
	}

	p.tokenStream.NextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.tokenStream.PeekTokenIs(token.COMMA) {
		p.tokenStream.NextToken()
		p.tokenStream.NextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.tokenStream.ExpectPeek(end) {
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

	p.tokenStream.NextToken()
	exp.Offset = p.parseExpression(LOWEST)

	if !p.tokenStream.ExpectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}

func (p *statementParser) parseObjectCreation() ast.Expression {

	oc := &ast.ObjectCreation{ast.OBJECT_CREATION, p.tokenStream.CurToken().Val, []ast.Expression{}};

	if (p.tokenStream.PeekTokenIs(token.LPAREN)) {
		p.tokenStream.NextToken()
		oc.Arguments = p.parseExpressionList(token.RPAREN)
	}

	return oc
}

func (p* statementParser) parseRunQuery() ast.Expression {

	r := &ast.RunQuery{ast.RUN_QUERY, nil}

	if (!p.tokenStream.ExpectPeek(token.IDENT)) {
		return nil
	}

	if (p.tokenStream.CurToken().Val != "query") {
		p.logError("Expected next identifier to be 'query', got '%s' instead", p.tokenStream.CurToken().Val)
		return nil
	}

	p.tokenStream.NextToken()

	r.Query = p.parseExpression(LOWEST)

	return r
}
