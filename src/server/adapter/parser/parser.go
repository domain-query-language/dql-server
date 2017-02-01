package parser

import (
	"errors"
	"fmt"
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/query/schema"
	"github.com/domain-query-language/dql-server/src/server/vm/handler"
)

/** Implementation of the adapter, written using the tokenizer, parser pattern */
type parser struct {

	t tokenizer.Tokenizer
	error error

	curToken  *token.Token
	peekToken *token.Token
}

func New(statements string) adapter.Adapter {

	t := tokenizer.NewTokenizer(statements);

	a:= &parser{t, nil, nil, nil}

	a.nextToken()
	a.nextToken()

	return a;
}

func (a *parser) nextToken() {

	a.curToken = a.peekToken
	a.peekToken = a.t.Next();
}

func (a *parser) curTokenIs(t token.TokenType) bool {

	return a.curToken.Typ == t
}

func (a *parser) peekTokenIs(t token.TokenType) bool {

	if a.peekToken == nil {
		return false
	}

	return a.peekToken.Typ == t
}

func (a *parser) expectPeek(t token.TokenType) bool {

 	if a.peekTokenIs(t) {

		a.nextToken()
		return true
	} else {

		a.peekError(t)
		return false
	}
}

func (a *parser) peekError(t token.TokenType) {

	if (a.peekToken == nil) {
		msg := fmt.Sprintf("Expected next token to be '%s', got EOF instead", t)
		a.error = errors.New(msg);
		return;
	}
	msg := fmt.Sprintf("Expected next token to be '%s', got '%s' instead", t, a.peekToken.Val)
	a.error = errors.New(msg);
}

// Return the next handlable object
func (a *parser) Next() (*adapter.Handleable, error) {

	if (a.curToken == nil) {

		return nil, nil;
	}

	qry := a.parseQuery();

	if (qry == nil) {

		return nil, a.error
	}

	return adapter.NewQuery(qry), nil;
}

func (a *parser) parseQuery() handler.Query {

	// This is where the type of object to be parsed is figured out
	if (a.curTokenIs(token.LIST)) {

		return a.parseListQuery();
	}

	a.error = errors.New("Unexpected token '"+a.curToken.Val+"'")
	return nil;
}

func (a *parser) parseListQuery() handler.Query {

	qry := &schema.ListDatabases{};

	if (!a.expectPeek(token.DATABASES)) {

		return nil;
	}

	if (!a.expectPeek(token.SEMICOLON)) {

		return nil;
	}

	return qry;
}