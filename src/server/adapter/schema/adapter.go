package schema

import (
	"errors"
	"fmt"
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"github.com/domain-query-language/dql-server/src/server/adapter/token"
	"github.com/domain-query-language/dql-server/src/server/adapter/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/query/schema"
	"github.com/domain-query-language/dql-server/src/server/vm/handler"
)

type queryAdapter struct {

	t tokenizer.Tokenizer
	error error

	curToken  *token.Token
	peekToken *token.Token
}

func NewQueryAdapter(statements string) adapter.Adapter {

	t := tokenizer.NewTokenizer(statements);

	a:= &queryAdapter{t, nil, nil, nil}

	a.nextToken()
	a.nextToken()

	return a;
}

func (a *queryAdapter) nextToken() {

	a.curToken = a.peekToken
	a.peekToken = a.t.Next();
}

func (a *queryAdapter) curTokenIs(t token.TokenType) bool {

	return a.curToken.Typ == t
}

func (a *queryAdapter) peekTokenIs(t token.TokenType) bool {

	if a.peekToken == nil {
		return false
	}

	return a.peekToken.Typ == t
}

func (a *queryAdapter) expectPeek(t token.TokenType) bool {

 	if a.peekTokenIs(t) {

		a.nextToken()
		return true
	} else {

		a.peekError(t)
		return false
	}
}

func (a *queryAdapter) peekError(t token.TokenType) {

	if (a.peekToken == nil) {
		msg := fmt.Sprintf("Expected next token to be %s, got nil instead", t)
		a.error = errors.New(msg);
		return;
	}
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead", t, a.peekToken.Typ)
	a.error = errors.New(msg);
}

func (a *queryAdapter) Next() (*adapter.Handleable, error) {

	if (a.curToken == nil) {

		return nil, nil;
	}

	qry := a.ParseQuery();

	if (qry == nil) {

		return nil, a.error
	}

	return adapter.NewQuery(qry), nil;
}

func (a *queryAdapter) ParseQuery() handler.Query {

	if (a.curTokenIs(token.LIST)) {

		return a.ParseListQuery();
	}

	a.error = errors.New("Unexpected token '"+a.curToken.Val+"'")
	return nil;
}

func (a *queryAdapter) ParseListQuery() handler.Query {

	qry := &schema.ListDatabases{};

	if (!a.expectPeek(token.DATABASES)) {

		return nil;
	}

	if (!a.expectPeek(token.SEMICOLON)) {

		return nil;
	}

	return qry;
}