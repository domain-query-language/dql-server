package command

import "github.com/domain-query-language/dql-server/src/server/vm"

type Command interface {

	Id() vm.Identifier

	TypeId() vm.Identifier

	AggregateTypeId() vm.Identifier
}
