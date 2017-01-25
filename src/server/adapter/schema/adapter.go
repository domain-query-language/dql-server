package schema

import (
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"github.com/domain-query-language/dql-server/src/server/adapter/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/query/schema"
	"github.com/domain-query-language/dql-server/src/server/adapter/token"
	"errors"
)

func NewQueryAdapter(statements string) adapter.Adapter {
	tknzr := tokenizer.NewTokenizer(statements);
	return &queryAdapter{tknzr};
}

type queryAdapter struct {
	tokenizer tokenizer.Tokenizer
}

func (a *queryAdapter) Next() (*adapter.Handleable, error) {

	tkn, _ := a.tokenizer.Next();
	if (tkn.Typ == token.LIST) {
		return adapter.NewQuery( &schema.ListDatabases{} );
	}
	return nil, errors.New("Unexpected token '"+tkn.Val+"'")

}