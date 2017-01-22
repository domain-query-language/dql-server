package query

import "github.com/domain-query-language/dql-server/src/server/vm"

type Query interface {

	Id() vm.Identifier
}
