package event

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

var TypeCreated = uuid.FromStringOrNil("b7dbb816-1141-4357-abbf-250e0cb7ec1f")

type Created struct {

	Name string
}

func (self Created) TypeId() vm.Identifier {
	return TypeCreated
}
