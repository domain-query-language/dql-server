package list_databases

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"
)

type Query struct {

}

func (self Query) TypeId() vm.Identifier {
	return Identifier
}

var QueryHandler = func(q vm.Query, queryable query.Queryable) query.Result {

	return &query.Result_ {
		queryable.(Queryable).List(),
	}
}