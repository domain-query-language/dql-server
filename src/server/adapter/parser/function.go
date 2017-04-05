package parser

import (
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
	"fmt"
	"errors"
)

/** Implementation of the adapter, written using the tokenizer, parser pattern */
type functionParser struct {

	t tokenizer.Tokenizer
	error error

	curToken  *token.Token
	peekToken *token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewFunction(function string) *functionParser {

	t := tokenizer.NewTokenizer(function);

	p := &functionParser{
		t: t,
	};

	p.nextToken()
	p.nextToken()

	return p;
}

func (p *functionParser) nextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.t.Next();
}

func (p *functionParser) curTokenIs(t token.TokenType) bool {

	return p.curToken.Type == t
}


func (p *functionParser) peekTokenIs(t token.TokenType) bool {

	if p.peekToken == nil {
		return false
	}

	return p.peekToken.Type == t
}

func (p *functionParser) expectPeek(t token.TokenType) bool {

	if p.peekTokenIs(t) {

		p.nextToken()
		return true
	} else {

		p.peekError(t)
		return false
	}
}

func (p *functionParser) expectCurrent(t token.TokenType) bool {

	if p.curTokenIs(t) {

		p.nextToken()
		return true
	} else {

		p.peekError(t)
		return false
	}
}

func (p *functionParser) peekError(t token.TokenType) {

	if (p.peekToken == nil) {
		p.logError("Expected next token to be '%s', got EOF instead", t)
		return;
	}
	p.logError("Expected next token to be '%s', got '%s' instead", t, p.peekToken.Val)
}

func (p *functionParser) logError(format string, a...interface{}) {

	msg := fmt.Sprintf(format, a...)
	p.error = errors.New(msg)
}

func (p *functionParser) ParseFunction() (*ast.Function, error) {

	function := &ast.Function{Type: ast.FUNCTION}

	p.expectPeek(token.IDENT)

	function.Name = p.curToken.Val

	p.expectPeek(token.LPAREN)

	p.nextToken()

	function.Parameters = p.parseParameters(token.RPAREN)

	function.Body = p.parseBlockStatement()

	return function, p.error
}

func (p *functionParser) parseParameters(end token.TokenType) []*ast.Parameter {

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

func (p *functionParser) parseParameter() *ast.Parameter{

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

func (p *functionParser) parseBlockStatement() *ast.BlockStatement {

	statementParser := NewStatementFromTokenizer(p.t, p.curToken, p.peekToken)

	blkStmnt, err := statementParser.ParseBlockStatement()

	if (err != nil) {
		p.error = err
		return nil
	}

	//TODO: Cleanup
	p.peekToken = statementParser.peekToken
	p.nextToken()

	return blkStmnt
}
