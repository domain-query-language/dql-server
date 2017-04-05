package parser

import (
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
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

func (p *functionParser) ParseFunction() (*ast.Function, error) {

	function := &ast.Function{ast.FUNCTION}

	return function, p.error
}
