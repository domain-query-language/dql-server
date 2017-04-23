package parser

import (
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
	"fmt"
	"errors"
)

/** Implementation of the adapter, written using the tokenizer, parser pattern */
type objectComponentParser struct {

	t tokenizer.Tokenizer
	error error

	curToken  *token.Token
	peekToken *token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewObjectComponent(function string) *objectComponentParser {

	t := tokenizer.NewTokenizer(function);

	p := &objectComponentParser{
		t: t,
	};

	p.nextToken()
	p.nextToken()

	return p;
}

func (p *objectComponentParser) nextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.t.Next();
}

func (p *objectComponentParser) curTokenIs(t token.TokenType) bool {

	return p.curToken.Type == t
}


func (p *objectComponentParser) peekTokenIs(t token.TokenType) bool {

	if p.peekToken == nil {
		return false
	}

	return p.peekToken.Type == t
}

func (p *objectComponentParser) expectPeek(t token.TokenType) bool {

	if p.peekTokenIs(t) {

		p.nextToken()
		return true
	} else {

		p.peekError(t)
		return false
	}
}

func (p *objectComponentParser) expectCurrent(t token.TokenType) bool {

	if p.curTokenIs(t) {

		p.nextToken()
		return true
	} else {

		p.peekError(t)
		return false
	}
}

func (p *objectComponentParser) peekError(t token.TokenType) {

	if (p.peekToken == nil) {
		p.logError("Expected next token to be '%s', got EOF instead", t)
		return;
	}
	p.logError("Expected next token to be '%s', got '%s' instead", t, p.peekToken.Val)
}

func (p *objectComponentParser) logError(format string, a...interface{}) {

	msg := fmt.Sprintf(format, a...)
	p.error = errors.New(msg)
}

func (p *objectComponentParser) ParseObjectComponent() (ast.ObjectComponent, error) {

	switch p.curToken.Type {

	case token.FUNCTION:
		return p.parseFunction()

	case token.CHECK:
		return p.parseCheck()

	case token.HANDLER:
		return p.parseHandler()

	case token.WHEN:
		return p.parseWhen()

	case token.PROPERTIES:
		return p.parseProperties()

	default:
		p.logError("Unexpected token '%s'", p.curToken.Type);
		return nil, p.error
	}
}

func (p *objectComponentParser) parseFunction() (*ast.Function, error) {

	function := &ast.Function{Type: ast.FUNCTION}

	p.expectPeek(token.IDENT)

	function.Name = p.curToken.Val

	p.expectPeek(token.LPAREN)

	p.nextToken()

	function.Parameters = p.parseParameters(token.RPAREN)

	function.Body = p.parseBlockStatement(token.LBRACE, token.RBRACE)

	return function, p.error
}

func (p *objectComponentParser) parseCheck() (*ast.Function, error) {

	check := &ast.Function{Type: ast.FUNCTION}

	check.Name = p.curToken.Val

	p.expectPeek(token.LPAREN)

	check.Parameters = []*ast.Parameter{}

	check.Body = p.parseBlockStatement(token.LPAREN, token.RPAREN)

	return check, p.error
}

func (p *objectComponentParser) parseHandler() (*ast.Function, error) {

	check := &ast.Function{Type: ast.FUNCTION}

	check.Name = p.curToken.Val

	p.expectPeek(token.LBRACE)

	check.Parameters = []*ast.Parameter{}

	check.Body = p.parseBlockStatement(token.LBRACE, token.RBRACE)

	return check, p.error
}

func (p *objectComponentParser) parseWhen() (*ast.When, error) {

	check := &ast.When{Type: ast.WHEN}

	p.expectPeek(token.IDENT)

	p.expectPeek(token.OBJECTNAME)

	check.Event = p.curToken.Val

	p.expectPeek(token.LBRACE)

	check.Body = p.parseBlockStatement(token.LBRACE, token.RBRACE)

	return check, p.error
}

func (p *objectComponentParser) parseParameters(end token.TokenType) []*ast.Parameter {

	list := []*ast.Parameter{}

	if p.curTokenIs(end) {
		p.nextToken()
		return list
	}

	list = append(list, p.parseParameter())

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseParameter())
	}

	if !p.expectPeek(end) {
		return nil
	}

	p.nextToken()

	return list
}

func (p *objectComponentParser) parseParameter() *ast.Parameter{

	if (!p.curTokenIs(token.OBJECTNAME)) {
		p.logError("Expected object name");
		return nil
	}

	param := &ast.Parameter {
		Type: p.curToken.Val,
	}
	p.expectPeek(token.IDENT)

	param.Name = p.curToken.Val

	return param
}

func (p *objectComponentParser) parseBlockStatement(open token.TokenType, close token.TokenType) *ast.BlockStatement {

	statementParser := NewStatementFromTokenizer(p.t, p.curToken, p.peekToken)

	blkStmnt, err := statementParser.ParseBlockStatementSurroundedBy(open, close)

	if (err != nil) {
		p.error = err
		return nil
	}

	//TODO: Cleanup
	p.peekToken = statementParser.peekToken
	p.nextToken()

	return blkStmnt
}

func (p *objectComponentParser) parseProperties() (*ast.Properties, error) {

	props := &ast.Properties{Type:ast.PROPERTIES}

	props.Properties = []*ast.Property{}

	p.expectPeek(token.LBRACE)
	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && p.error == nil  {
		prop := p.parseProperty()
		if prop != nil {
			props.Properties = append(props.Properties, prop)
		}

		if p.curToken == nil {
			p.logError("Unexpected EOF")
			break;
		}
	}


	return props, p.error
}

func (p *objectComponentParser) parseProperty() *ast.Property {

	prop := &ast.Property{Type:ast.PROPERTY}

	prop.ValueType = p.curToken.Val

	p.expectPeek(token.IDENT)

	prop.Name = p.curToken.Val

	if (p.peekTokenIs(token.SEMICOLON)) {
		p.nextToken()
		p.nextToken()
		return prop
	}

	p.expectPeek(token.ASSIGN);
	p.nextToken()

	prop.Exp = p.parseObjectCreation()

	return prop
}

func (p *objectComponentParser) parseObjectCreation() ast.Expression {

	return nil
}



