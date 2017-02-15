package list_databases

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
)

var QueryId, _ = uuid.FromString("c50d6791-3fc5-4be8-91fc-c01f20526872")

type ListDatabases struct {

}

func (self ListDatabases) TypeId() vm.Identifier {
	return QueryId
}

var Handler = func(q vm.Query, queryable query.Queryable) query.Result {

	return &query.Result_ {
		queryable.(list_databases.Queryable).List(),
	}
}
