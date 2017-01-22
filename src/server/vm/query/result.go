package query

import (
	"github.com/domain-query-language/dql-server/src/server/vm"
)

type Result interface {

	Id() vm.Identifier

	TypeId() vm.Identifier

	MarshalJSON() ([]byte, error)
}
