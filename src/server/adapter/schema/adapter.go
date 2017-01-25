package schema

import (
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"github.com/domain-query-language/dql-server/src/server/adapter/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/query/schema"
)

func NewQueryAdapter(statements string) adapter.Adapter {
	tknzr := tokenizer.NewTokenizer(statements);
	return &queryAdapter{tknzr};
}

type queryAdapter struct {
	tokenizer tokenizer.Tokenizer
}

func (a *queryAdapter) Next() (*adapter.Handleable, error) {

	return adapter.NewQuery( &schema.ListDatabases{} ), nil

}